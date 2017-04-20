package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mydb"
	"time"
)

type Links struct {
	ID      bson.ObjectId `bson:"_id" json:"-"`
	MovieID string        `bson:"movieID" json:"movieID"`
	ImdbID  string        `bson:"imdbID" json:"imdbID"`
	TmdbID  string        `bson:"tmdbID" json:"tmdbID"`

	UpdatedAt time.Time `bson:"updatedAt" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"-"`
}

func LinksCollectionName() string {
	return "local.links"
}

func (m Links) Insert() (err error) {
	mydb.Exec(LinksCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}
