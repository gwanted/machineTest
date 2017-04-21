package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"mydb"
)

type Movies struct {
	ID      bson.ObjectId `bson:"_id" json:"-"`
	MovieID string        `bson:"movieID" json:"movieID"`
	Title   string        `bson:"title" json:"title"`
	Genres  string        `bson:"genres" json:"genres"`

	UpdatedAt time.Time `bson:"updatedAt" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"-"`
}

func MoviesCollectionName() string {
	return "local.movies"
}

func (m Movies) Insert() (err error) {
	mydb.Exec(MoviesCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}

func FindMovies(condition bson.M) (result []*Movies, err error) {
	mydb.Exec(MoviesCollectionName(), func(c *mgo.Collection) {
		err = c.Find(condition).All(&result)
	})
	return
}
