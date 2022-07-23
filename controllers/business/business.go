package business

import (
	"echoinit/apps"

	"github.com/labstack/echo/v4"
)

func Hello(ctx echo.Context) error {
	apps.Logs.Info("[business.Hello()]")
	apps.Logs.Debug("Hello")

	return nil
}
