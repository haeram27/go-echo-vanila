package routers

import (
	"github.com/labstack/echo/v4"

	"echoinit/apps"
	"echoinit/controllers/business"
)

func Route(e *echo.Echo) {
	var dataRoutes = []struct {
		mode    string
		path    string
		handler echo.HandlerFunc
	}{
		{"GET", "/hello", business.Hello},
		{"POST", "/hello", business.Hello},

		{"GET", "/corona", business.CoronaStatus},
		{"POST", "/corona", business.CoronaStatus},
	}

	for _, data := range dataRoutes {
		apps.Logs.Debugf("%s %s", data.mode, data.path)

		switch data.mode {
		case "GET":
			e.GET(data.path, data.handler)
		case "PUT":
			e.PUT(data.path, data.handler)
		case "POST":
			e.POST(data.path, data.handler)
		case "DELETE":
			e.DELETE(data.path, data.handler)
		default:
			apps.Logs.Errorf("unknown mode '%s'", data.mode)
		}
	}

}
