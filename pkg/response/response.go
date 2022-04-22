package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int         `json:"status"`
	Message Language    `json:"message,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Language struct {
	Bahasa  interface{} `json:"Indonesian,omitempty" validate:"max=64,min=1"`
	English interface{} `json:"English,omitempty" validate:"max=64,min=1"`
}

func CustomError(ctx echo.Context, httpstatus int, status int, msg Language, meta interface{}, data interface{}) error {
	return ctx.JSON(httpstatus, Response{
		Status:  status,
		Message: msg,
		Meta:    meta,
		Data:    data,
	})
}
