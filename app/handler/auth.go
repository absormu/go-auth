package handler

import (
	"net/http"

	usecaseauth "github.com/absormu/go-auth/app/usecase/auth"
	"github.com/labstack/echo/v4"

	"github.com/absormu/go-auth/app/entity"
	md "github.com/absormu/go-auth/app/middleware"
	lg "github.com/absormu/go-auth/pkg/response"
	resp "github.com/absormu/go-auth/pkg/response"
	sdk "gitlab.com/d3386/library"
)

func LoginHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: LoginHandler")

	req := entity.Auth{}
	if e = c.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: e.Error()}, nil, nil)
		return
	}

	e = usecaseauth.Login(c, req)

	return
}

func SignupHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: SignupHandler")

	req := entity.SellerData{}
	if e = c.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: e.Error()}, nil, nil)
		return
	}

	e = usecaseauth.Signup(c, req)

	return
}
