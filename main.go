package main

import (
	"github.com/kataras/iris"
)

var addr = iris.Addr(":8101")

func main() {
	app := iris.New()

	app.Get("/ping", func(ctx iris.Context) {
		ctx.HTML("Hello World!")
	})

	app.Run(iris.Addr(":8080"))
}
