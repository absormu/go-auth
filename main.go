package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func initHandlers(e *echo.Echo) {
	// root := e.Group(cm.Config.RootURL)
	// root.POST("/login", handler.LoginHandler)
	// root.POST("/signup", handler.SignupHandler)
	// root.POST("/forgot-passsword", handler.ForgotPasswordHandler)
	// root.GET("/verification-passsword/:id", handler.VerificationPasswordHandler)
	// root.POST("/reset-passsword", handler.ResetPasswordHandler)

	// Start serverlog.Info()
	log.Info("Staring server ...")
}

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":9000"))
}
