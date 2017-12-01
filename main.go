package main

import (
	"fmt"
	"net/http"
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

func App() http.Handler {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	return nil
}


func main() {
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	vDataSource := datasource.NewDataSource()
	vBbsRepository := bbsRepository.NewBbsRepository(vDataSource)
	vBbsService := bbsService.NewBbsService(vBbsRepository)
	app.Controller("/bbs", new(bbsController.BbsController), vBbsService)
	app.Run(iris.Addr(":8080"))
}
