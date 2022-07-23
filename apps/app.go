package apps

import (
	"context"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	appctx    context.Context
	appCancel context.CancelFunc
	flogger   *lumberjack.Logger
	handler   *echo.Echo
	Logs      echo.Logger
)

func ApplicationContext() context.Context {
	return appctx
}

func ApplicationData() *AppData {
	return appctx.Value(ctxKeyAppData).(*AppData)
}

func Echo() *echo.Echo {
	return handler
}

func Close() {
	appCancel()

	if handler != nil {
		if e := handler.Shutdown(context.Background()); e != nil {
			Logs.Fatal(e)
		}

		if e := handler.Close(); e != nil {
			Logs.Fatal(e)
		}
	}

	if flogger != nil {
		flogger.Close()
	}
}

func init() {
	// make echo instances
	handler = echo.New()

	// setup logger
	flogger = &lumberjack.Logger{
		Filename:   applicationName + ".log",
		MaxBackups: 3,
		MaxAge:     7, //days
	}
	Logs = handler.Logger
	Logs.SetOutput(io.MultiWriter(os.Stdout, flogger))
	Logs.SetHeader("[${time_rfc3339}][${short_file}:${line}][${level}]")
	Logs.SetLevel(log.DEBUG)

	// setup application context
	appctx, appCancel = initContext()
}
