package main

import (
	"log"

	docs "github.com/Natalieihs/go-fasthttp-swagger-demo/docs"
	"github.com/Natalieihs/go-fasthttp-swagger-demo/handlers"
	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name T

func main() {
	docs.SwaggerInfo.Title = "Fasthttp Swagger"
	docs.SwaggerInfo.Description = "Fasthttp Swagger with Custom Authentication"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""

	r := router.New()
	r.GET("/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	r.GET("/hello", Hello())
	r.GET("/user", handlers.GetUser)

	log.Println("Server is running at http://localhost:9090")
	if err := fasthttp.ListenAndServe(":9090", r.Handler); err != nil {
		panic(err)
	}
}

// Hello godoc
// @Summary Show a Hello
// @Description get hello
// @Tags hello
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /hello [get]
func Hello() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// 检查认证
		token := string(ctx.Request.Header.Peek("T"))
		if token == "" {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.SetBodyString("Unauthorized: Missing token")
			return
		}
		// 这里可以添加更多的token验证逻辑

		ctx.SetStatusCode(200)
		ctx.SetBody([]byte(`{ "message" : "HELLO" }`))
	}
}
