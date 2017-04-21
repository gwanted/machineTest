package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"mydb"
)

type Tags struct {
	ID      bson.ObjectId `bson:"_id" json:"-"`
	UserID  string        `bson:"userID" json:"userID"`
	MovieID string        `bson:"movieID" json:"movieID"`
	Tag     string        `bson:"tag" json:"tag"`

	UpdatedAt time.Time `bson:"updatedAt" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"-"`
}

func TagsCollectionName() string {
	return "local.tags"
}

func (m Tags) Insert() (err error) {
	mydb.Exec(TagsCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}
