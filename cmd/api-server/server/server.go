package server

import (
	"context"
	"fmt"
	"net/http"

	"nueip/cmd/api-server/models"
	"nueip/cmd/api-server/server/routes"
	pkgEcho "nueip/pkg/echo"

	"github.com/labstack/echo/v4"
)

type Server struct {
	dao  *models.DAO
	port string
	srv  *echo.Echo
	ctx  context.Context
}

func (w *Server) WithPort(p string) *Server {
	w.port = p
	return w
}

func (w *Server) WithContext(ctx context.Context) *Server {
	w.ctx = ctx
	return w
}

func New(dao *models.DAO) *Server {
	e := echo.New()
	e.Validator = pkgEcho.NewValidate()
	return &Server{
		srv: e,
		dao: dao,
	}
}

func (w *Server) Start() error {
	port := w.port
	srv := w.srv
	dao := w.dao
	if err := dao.Init(); err != nil {
		return handleServerError(srv, "failed to initial the db: %v", err)
	}

	if err := routes.Init(srv, dao); err != nil {
		return handleServerError(srv, "failed to load the routes (:%s): %v", port, err)
	}

	if err := srv.Start(port); err != nil && err != http.ErrServerClosed {
		return handleServerError(srv, "failed to active the server (:%s): %v", port, err)
	}
	return nil
}

func (w *Server) Stop() error {
	srv := w.srv
	if err := srv.Close(); err != nil {
		return handleServerError(srv, "failed to close api server: %v", err)
	}

	if err := srv.Shutdown(w.ctx); err != nil {
		return handleServerError(srv, "failed to shutdown api server: %v", err)
	}
	return nil
}

func handleServerError(srv *echo.Echo, format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	srv.Logger.Fatal(err.Error())
	return err
}
