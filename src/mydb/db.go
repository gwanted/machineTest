package mydb

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

func InitDB(urls string) error {
	sessions, err := mgo.Dial(urls) //连接服务器
	if err != nil {
		panic(err)
	}
	Session = sessions
	fmt.Println("mongodb init finish")
	return nil
}

func GetDbCollection(collection string) (db string, collectionName string) {
	db = ""
	collectionName = ""
	strs := strings.Split(collection, ".")
	if strs[0] != "" {
		db = strs[0]
	}
	if strs[1] != "" {
		collectionName = strs[1]
	}
	return
}

func Exec(collection string, f func(*mgo.Collection)) {
	db, collectionName := GetDbCollection(collection)
	c := Session.DB(db).C(collectionName)
	f(c)
}
