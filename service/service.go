package service

import (
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
