package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mydb"
	"time"
)

type Ratings struct {
	ID      bson.ObjectId `bson:"_id" json:"-"`
	UserID  string        `bson:"userID" json:"userID"`
	MovieID string        `bson:"movieID" json:"movieID"`
	Rating  string        `bson:"rating" json:"rating"`

	UpdatedAt time.Time `bson:"updatedAt" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"-"`
}

func RatingsCollectionName() string {
	return "local.ratings"
}

func (m Ratings) Insert() (err error) {
	mydb.Exec(RatingsCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}

func FindRatings(condition bson.M) (result []*Ratings, err error) {
	mydb.Exec(RatingsCollectionName(), func(c *mgo.Collection) {
		err = c.Find(condition).All(&result)
	})
	return
}
