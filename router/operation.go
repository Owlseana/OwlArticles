package router

import (
	"owlarticles/model"

	"github.com/gin-gonic/gin"
)

func deleteArticle(c *gin.Context) {
	var (
		err   error
		param *model.DeleteRequest
	)
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	srv.DeleteArticle(c, param)
}
