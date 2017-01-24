package main

import (
	"net/http"
	"mydb"
	"conf"
	"app"
)


func main() {
	mydb.InitDB(conf.App.DBAddress)

	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/register", app.Register)

	http.ListenAndServe(":8888", nil)
}