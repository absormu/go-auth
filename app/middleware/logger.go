package middleware

import (
	cm "github.com/absormu/go-auth/pkg/configuration"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

// AddLogger .
func AddLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("logger", log.WithFields(log.Fields{
			"app":    cm.Config.OriginHost,
			"msg_id": xid.New().String(),
		}))
		return next(c)
	}
}

// GetLogger .
func GetLogger(ctx echo.Context) *log.Entry {
	logger := ctx.Get("logger")
	if logger != nil {
		return logger.(*log.Entry)
	}
	return log.WithFields(log.Fields{
		"app": cm.Config.OriginHost,
	})

}
