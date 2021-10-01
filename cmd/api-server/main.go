package main

import (
	"fmt"
	"nueip/cmd/api-server/app"
	pkgApp "nueip/pkg/lib/app"
	"os"
)

func main() {
	if err := pkgApp.Handle("API Server", app.New()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit()
}
