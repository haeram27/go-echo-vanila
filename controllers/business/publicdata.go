package business

import (
	"echoinit/controllers/clients/restapis"

	"context"

	"github.com/labstack/echo/v4"
)

func CoronaStatus(ctx echo.Context) error {
	restapis.CoronaStatus(context.Background())
	return nil
}
