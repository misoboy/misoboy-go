package main

import (
	"fmt"
	"github.com/danryan/env"
	"github.com/unrolled/render"
	bbsController "misoboy_web/controllers/bbs"
	bbsService "misoboy_web/services/bbs"
	bbsRepository "misoboy_web/repository/bbs"
	datasource "misoboy_web/common/datasource"
	"github.com/kataras/iris"
)

type Config struct {
	Environment string `env:"key=ENVIRONMENT default=development"`
	Port        string `env:"key=PORT default=9000"`
	EnableCors  string `env:"key=ENABLE_CORS default=false"`
}

var (
	renderer	*render.Render
	config	*Config
)

func init() {
	var option render.Options
	config = &Config{}
	if err := env.Process(config); err != nil {
		fmt.Println(err)
	}
	if config.Environment == "development" {
		option.IndentJSON = true
	}
	renderer = render.New(option)
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	tmpl := iris.HTML("./views", ".html").
		Layout("layouts/layout.html").
		Reload(true)
	app.StaticWeb("/public", "./public")
	app.RegisterView(tmpl)
	vDataSource := datasource.NewDataSource()
	vBbsRepository := bbsRepository.NewBbsRepository(vDataSource)
	vBbsService := bbsService.NewBbsService(vBbsRepository)
	app.Controller("/bbs", new(bbsController.BbsController), vBbsService)
	//app.Run(iris.Addr(":8081"))

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("common/error.html")
	})

	app.Run(
		iris.Addr("localhost:8081"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)
}
