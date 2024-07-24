package dao

import (
	"net/http"
	"owlarticles/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *Dao) CreateArticle(c *gin.Context, article *model.Article) (err error) {
	if err = c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.ID = primitive.NewObjectID()
	article.CreatedAt = time.Now().Format(time.RFC3339)
	article.UpdatedAt = time.Now().Format(time.RFC3339)

	_, err = d.articleCol.InsertOne(c, article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Errorf("CreateArticle collection.InsertOne error %v ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article added successfully!", "article": article})
	return
}

func (d *Dao) GetArticleList(c *gin.Context, article *model.Article) (err error) {
	var articles []model.Article

	filter := bson.M{}
	// 执行查询
	cur, err := d.articleCol.Find(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Errorf("GetArticleList collection.Find error %v ", err)
		return
	}
	defer cur.Close(c)

	for cur.Next(c) {
		var article model.Article
		if err = cur.Decode(&article); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Errorf("GetArticleList cursor.Decode error %v ", err)
			return
		}
		articles = append(articles, article)
	}

	if err = cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Errorf("GetArticleList cursor error %v ", err)
		return
	}

	// 返回文章列表
	c.JSON(http.StatusOK, gin.H{"articles": articles})

	return
}

func (d *Dao) DeleteArticleByID(c *gin.Context, id string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		log.Errorf("DeleteArticleByID error %v ", err)
		return err
	}

	filter := bson.M{"id": objectID}
	// 执行删除操作
	res, err := d.articleCol.DeleteOne(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Errorf("DeleteArticle collection.DeleteOne error %v ", err)
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully!"})
	return
}

func (d *Dao) DeleteArticleByTitle(c *gin.Context, title string) (err error) {
	filter := bson.M{"title": title}
	// 执行删除操作
	res, err := d.articleCol.DeleteOne(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Errorf("DeleteArticle collection.DeleteOne error %v ", err)
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully!"})
	return
}
