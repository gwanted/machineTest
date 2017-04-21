package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"mydb"
)

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Account   string        `bson:"account" json:"account"`
	Pwd       string        `bson:"pwd" json:"pwd"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"-"`
	CreatedAt time.Time     `bson:"createdAt" json:"-"`
}

func UserCollectionName() string {
	return "local.user"
}

func (m User) Insert() (err error) {
	mydb.Exec(UserCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}

func FindUser(condition bson.M) (result *User, err error) {
	mydb.Exec(UserCollectionName(), func(c *mgo.Collection) {
		err = c.Find(condition).One(&result)
	})
	return
}

func FindUsers(condition bson.M) (result []*User, err error) {
	mydb.Exec(UserCollectionName(), func(c *mgo.Collection) {
		err = c.Find(condition).All(&result)
	})
	return
}
