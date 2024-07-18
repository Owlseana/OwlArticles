package dao

import (
	"owlarticles/conf"

	"github.com/quan-xie/tuba/database/mongo"
	xmongo "go.mongodb.org/mongo-driver/mongo"
)

type Dao struct {
	mongoClient *xmongo.Client
	articleCol  *xmongo.Collection
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		mongoClient: mongo.NewMongo(c.Mongo),
	}

	d.articleCol = d.mongoClient.Database("articles").Collection("article")
	return
}
