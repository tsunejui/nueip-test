package echo

import (
	"fmt"
	"reflect"

	modEcho "github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler modEcho.HandlerFunc
}

type RouteGroup struct {
	Prefix     string
	Routes     []*Route
	Mideleware []modEcho.MiddlewareFunc
}

type RouteGroups struct {
	groups []*RouteGroup
}

func NewGroups() *RouteGroups {
	return &RouteGroups{}
}

func (rg *RouteGroups) Register(group *RouteGroup) *RouteGroups {
	rg.groups = append(rg.groups, group)
	return rg
}

func (rg *RouteGroups) Enrich(e *modEcho.Echo) error {
	for _, g := range rg.groups {
		fmt.Printf("\n\nPrefix: %s", g.Prefix)
		echoGroup := e.Group(g.Prefix)
		echoGroup.Use(g.Mideleware...)
		for _, r := range g.Routes {
			fmt.Printf("\nload path: [%s] %s", r.Method, r.Path)
			function := reflect.ValueOf(echoGroup).MethodByName(r.Method)
			function.Call([]reflect.Value{
				reflect.ValueOf(r.Path),
				reflect.ValueOf(r.Handler),
			})
		}
	}
	return nil
}
