package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func main() {
	app := iris.Default()
	app.Get("/hello", func(ctx iris.Context) {
		ctx.WriteString("hello world!")
	})
	v1 := app.Party("/v1")
	v1.Use(func(ctx iris.Context) {
		logrus.Info("自定义中间件")
		ctx.Next()
	})
	v1.Get("/users/{id:uint64 min(2)}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0)
		ctx.WriteString(strconv.Itoa(int(id)))
	})
	v1.Get("/orders/{action:string prefix(a_)}", func(ctx iris.Context) {
		a := ctx.Params().Get("action")
		ctx.WriteString(a)
	})
	//有panic时会调用这个
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.WriteString("看起来出错了")
	})
	app.OnErrorCode(http.StatusNotFound, func(ctx iris.Context) {
		ctx.WriteString("访问路径不存在")
	})
	err := app.Run(iris.Addr(":8082"))
	fmt.Println(err)
}
