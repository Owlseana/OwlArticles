package dao

import (
	"net/http"
	"owlarticles/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *Dao) CreateArticle(c *gin.Context, article *model.Article) {
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.ID = primitive.NewObjectID()
	article.CreatedAt = time.Now().Format(time.RFC3339)
	article.UpdatedAt = time.Now().Format(time.RFC3339)

	_, err := d.articleCol.InsertOne(c, article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article added successfully!", "article": article})
}
