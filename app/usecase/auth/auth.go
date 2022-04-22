package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/absormu/go-auth/app/constants"
	"github.com/absormu/go-auth/app/dto"
	"github.com/absormu/go-auth/app/entity"
	md "github.com/absormu/go-auth/app/middleware"
	repoauth "github.com/absormu/go-auth/app/repository/auth"
	cm "github.com/absormu/go-auth/pkg/configuration"
	lg "github.com/absormu/go-auth/pkg/response"
	resp "github.com/absormu/go-auth/pkg/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	sdk "gitlab.com/d3386/library"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context, req entity.Auth) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", req).Info("usecase: Login")

	// Throws unauthorized error
	if req.Email == "" || req.Password == "" {
		logger.Error("Catch error missing mandatory parameter")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: nil, English: "Missing mandatory parameter"}, nil, nil)
		return
	}

	// cek email & get password
	params := make(map[string]string)
	params["email"] = req.Email
	params["active"] = "1"
	params["is_deleted"] = "0"

	var user entity.User

	if user, e = repoauth.GetAuthEmail(c, params); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetAuthEmail")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query"}, nil, nil)
		return
	}

	// login bycrypt
	passDB := user.Password
	password := req.Password
	match := CheckPasswordHash(password, passDB)

	if !match {
		logger.Error("Catch error user not found")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_USER_NOT_FOUND,
			lg.Language{Bahasa: "Email atau kata sandi salah", English: "Email or password is not correct"}, nil, nil)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = xid.New().String()
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = req.Email
	claims["user_contact_id"] = user.UserContactID
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Duration(cm.Config.TokenLifeTime) * time.Second).Unix()

	// Generate encoded token and send it as response.
	t, e := token.SignedString([]byte(cm.Config.ClientSecret))
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error generate encoded token")
		e = resp.CustomError(c, http.StatusUnauthorized, sdk.ERR_UNAUTHORIZED,
			lg.Language{Bahasa: nil, English: "Unauthorized"}, nil, nil)
		return
	}

	res := entity.OAuthMessage{
		AccessToken: t,
		TokenType:   "bearer",
		ExpiresIn:   cm.Config.TokenLifeTime,
	}

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, res)
	return
}

func Signup(c echo.Context, req entity.SellerData) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", req).Info("usecase: Signup")

	if req.Name == "" || req.Email == "" || req.Password == "" {
		logger.Error("Missing mandatory parameter")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_MISSING,
			lg.Language{Bahasa: nil, English: "Missing mandatory parameter"}, nil, nil)
		return
	}

	// bcrypt hashing/ save hasil hashing password
	hashPassword, err := HashPassword(req.Password)
	if err != nil {
		println(fmt.Println("Error hashing password"))
		return
	}

	var user entity.User
	// params
	params := make(map[string]string)
	params["email"] = req.Email

	if user, e = repoauth.GetAuth(c, params); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query GetAuth")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query"}, nil, nil)
		return
	}

	var empty entity.User
	if user != empty {
		logger.Error("Catch error user already exist")
		e = resp.CustomError(c, http.StatusForbidden, sdk.ERR_USER_EXIST,
			lg.Language{Bahasa: "Email sudah ada", English: "Email already exist"}, nil, nil)
		return
	}

	// params
	paramsSeller := make(map[string]interface{})
	// params seller
	paramsSeller["name"] = req.Name
	paramsSeller["city_id"] = req.City.ID
	if req.Telephone != nil {
		paramsSeller["telephone"] = req.Telephone
	}
	if req.Address != nil {
		paramsSeller["address"] = req.Address
	}
	// type 2 seller
	paramsSeller["type"] = 2
	paramsSeller["active"] = 1
	paramsSeller["created_by"] = "SYSTEM"

	// params user
	paramsUser := make(map[string]interface{})
	paramsUser["name"] = req.Name
	paramsUser["email"] = req.Email
	paramsUser["password"] = hashPassword
	// role id admin = 2
	paramsUser["role_id"] = 2
	paramsUser["active"] = 1
	paramsUser["created_by"] = "SYSTEM"

	if e = repoauth.Signup(c, paramsSeller, paramsUser); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error failure query Signup")
		e = resp.CustomError(c, http.StatusInternalServerError, sdk.ERR_DATABASE,
			lg.Language{Bahasa: nil, English: "Failure query signup"}, nil, nil)
		return
	}

	to := fmt.Sprintf("%v", params["email"])

	reqEmail := dto.SetEmail(to, constants.MSG3, constants.MSG4)
	go repoauth.SignupNotification(c, reqEmail)

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, nil)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
