package main

import (
	"net/http"
	"os"
	"server/router"
	"server/util"
)

func main() {
	err := http.ListenAndServe(":"+os.Getenv("API_SERVER_PORT"), router.InitRouter())
	util.CheckRuntimeError(err, "http open err")
}
