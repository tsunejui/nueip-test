package routes

import (
	"net/http"
	pkgEcho "nueip/pkg/echo"

	"github.com/labstack/echo/v4"
)

func init() {
	backend.Register(&pkgEcho.RouteGroup{
		Prefix: "",
		Routes: []*pkgEcho.Route{
			{
				Method:  http.MethodGet,
				Path:    "pong",
				Handler: apiHandler.Ping,
			},
		},
	})
}

func (h *handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
