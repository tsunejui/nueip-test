package routes

import (
	"fmt"
	"nueip/cmd/api-server/models"
	pkgEcho "nueip/pkg/echo"
	"nueip/pkg/lib/web"

	"github.com/labstack/echo/v4"
)

var (
	backend    *pkgEcho.RouteGroups = pkgEcho.NewGroups()
	apiHandler *handler             = &handler{}
)

type handler struct {
	web.Handler
	DAO *models.DAO
}

func Init(server *echo.Echo, dao *models.DAO) error {
	apiHandler.DAO = dao
	if err := backend.Enrich(server); err != nil {
		return fmt.Errorf("failed to enrich APIs on server: %v", err)
	}
	return nil
}
