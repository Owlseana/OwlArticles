package service

import (
	"net/http"
	"owlarticles/conf"
	"owlarticles/dao"
	"owlarticles/model"

	"github.com/gin-gonic/gin"
)

var (
	article *model.Article
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return
}

func (s *Service) CreateArticle(c *gin.Context) {
	s.dao.CreateArticle(c, article)
}

func (s *Service) GetArticleList(c *gin.Context) {
	s.dao.GetArticleList(c, article)
}

func (s *Service) DeleteArticle(c *gin.Context, param *model.DeleteRequest) {
	// if err := c.ShouldBindJSON(param); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	log.Errorf("DeleteArticle ShouldBindJSON error %v ", err)
	// 	return
	// }

	if param.ID != "" {
		s.dao.DeleteArticleByID(c, param.ID)
		return
	}
	if param.Title != "" {
		s.dao.DeleteArticleByTitle(c, param.Title)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "ID or Title required"})
}
