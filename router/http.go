package router

import (
	"owlarticles/conf"
	"owlarticles/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
)

var (
	httpClient *fasthttp.Client
	srv        *service.Service
)

func startHttp(c *conf.Config) {
	httpClient = &fasthttp.Client{
		MaxConnsPerHost: 16384, // MaxConnsPerHost  default is 512, increase to 16384
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
	}
	srv = service.New(c)
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create", srv.CreateArticle)
	r.POST("/delete", srv.DeleteArticle)
	r.GET("/list", srv.GetArticleList)
	r.Run("192.168.128.129:8000")
}
