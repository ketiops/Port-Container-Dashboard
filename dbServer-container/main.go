package main

import (
	"db-server/model"
	"db-server/router"
	"db-server/util"
	"net/http"
	"os"
)

func main() {
	model.InitDB()
	err := http.ListenAndServe(":"+os.Getenv("DB_SERVER_PORT"), router.InitRouter())
	util.CheckRuntimeError(err, "http open err")
}
