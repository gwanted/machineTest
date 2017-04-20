package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mydb"
	"time"
)

type GenomeTags struct {
	ID    bson.ObjectId `bson:"_id" json:"-"`
	TagID string        `bson:"tagID" json:"tagID"`
	Tag   string        `bson:"tag" json:"tag"`

	UpdatedAt time.Time `bson:"updatedAt" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"-"`
}

func GenomeTagsCollectionName() string {
	return "local.genometags"
}

func (m GenomeTags) Insert() (err error) {
	mydb.Exec(GenomeTagsCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}
