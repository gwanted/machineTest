package app

import (
	"mydb"
	"model"
	"net/http"
	"common"
	"gopkg.in/mgo.v2/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pwd := r.FormValue("pwd")
	if user == "" {
		common.ReturnEFormat(w, 1, "user is null")
		return
	}
	if pwd == "" {
		common.ReturnEFormat(w, 1, "pwd is null")
		return
	}
	var dddd model.User
	db := mydb.GetDbCollection("local", "user")
	err := db.Find(bson.M{"account":user, "pwd":pwd}).One(&dddd)
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}
	common.ReturnFormat(w, 0, dddd, "SUCCESS")
}

func Register(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pwd := r.FormValue("pwd")
	if user == "" {
		common.ReturnEFormat(w, 1, "user is null")
		return
	}
	if pwd == "" {
		common.ReturnEFormat(w, 1, "pwd is null")
		return
	}
	db := mydb.GetDbCollection("local", "user")
	var dddd []model.User
	err := db.Find(bson.M{"account":user, "pwd":pwd}).All(&dddd)
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}
	if len(dddd) != 0 {
		common.ReturnEFormat(w, 1, "user is already exist")
		return
	}
	err = db.Insert(bson.M{"account":user, "pwd":pwd, "status":"N"})
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}
	common.ReturnFormat(w, 0, nil, "SUCCESS")
}
