package jwt

import (
	"net/http"
	"strings"

	"github.com/absormu/go-auth/app/entity"
	md "github.com/absormu/go-auth/app/middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, nil)
	if err == nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func ExtractToken(c echo.Context) (extractToken entity.ExtractToken, e error) {
	logger := md.GetLogger(c)
	authorization := c.Request().Header.Get(echo.HeaderAuthorization)
	if authorization == "" {
		logger.WithField("error", "Catch error Authorization missing")
		return extractToken, echo.NewHTTPError(http.StatusUnauthorized)
	}
	tokens := strings.SplitN(authorization, " ", 2)

	if len(tokens) < 2 || strings.ToLower(tokens[0]) != "bearer" {
		logger.WithField("error", "Catch error Invalid bearer type")
		return extractToken, echo.NewHTTPError(http.StatusUnauthorized)
	}

	tokenMap, err := ExtractClaims(tokens[1])
	if err != nil {
		logger.WithField("error", "Catch error Error parse token")
		return extractToken, echo.NewHTTPError(http.StatusUnauthorized)
	}

	extractToken = entity.ExtractToken{
		Name:          tokenMap["name"].(string),
		UserID:        int64(tokenMap["user_id"].(float64)),
		Email:         tokenMap["email"].(string),
		UserContactID: int64(tokenMap["user_contact_id"].(float64)),
		RoleID:        int64(tokenMap["role_id"].(float64)),
	}

	return
}
