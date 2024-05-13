package controller

import (
	"os"
)

var (
	dbServerHost string
	dbServerPort string
)

func init() {
	dbServerHost = os.Getenv("DB_SERVER_HOST")
	dbServerPort = os.Getenv("DB_SERVER_PORT")
}
