package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mydb"
	"time"
)

type GenomeScores struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	MovieID   string        `bson:"movieID" json:"movieID"`
	TagID     string        `bson:"tagID" json:"tagID"`
	Relevance string        `bson:"relevance" json:"relevance"`

	UpdatedAt time.Time `bson:"updatedAt" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"-"`
}

func GenomeScoresCollectionName() string {
	return "local.genomescores"
}

func (m GenomeScores) Insert() (err error) {
	mydb.Exec(GenomeScoresCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}
