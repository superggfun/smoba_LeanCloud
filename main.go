package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leancloud/go-sdk/leancloud"
	"github.com/leancloud/golang-getting-started/adapters"
	_ "github.com/leancloud/golang-getting-started/functions"
)

func main() {
	leancloud.Engine.Init(leancloud.NewEnvClient())

	port := os.Getenv("LEANCLOUD_APP_PORT")
	if port == "" {
		port = "3000"
	}

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, host=${host}, status=${status}, error=${error}, remote=${remote_ip}\n",
		Output: os.Stdout,
	}))

	adapters.Echo(e)

	e.Logger.Fatal(e.Start(":" + port))
}
