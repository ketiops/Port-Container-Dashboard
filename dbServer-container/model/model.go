package model

import (
	"database/sql"
	"db-server/util"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	util.CheckRuntimeError(err, "db open err")

	createPreInformationTableQuery := `
	CREATE TABLE IF NOT EXISTS PreInformation (
		id SERIAL PRIMARY KEY,
		inspEqNo VARCHAR(20) NOT NULL,
		cntrNo VARCHAR(11) NOT NULL,
		truckNo VARCHAR(8),
		typeNo VARCHAR(20),
		qDate DATE
	);`
	_, err = db.Exec(createPreInformationTableQuery)
	util.CheckRuntimeError(err, "Information db create err")

	createPreInformationHistoryTableQuery := `
	CREATE TABLE IF NOT EXISTS PreInformationHistory (
		id SERIAL PRIMARY KEY,
		inspEqNo VARCHAR(20) NOT NULL,
		cntrNo VARCHAR(11) NOT NULL,
		truckNo VARCHAR(8),
		typeNo VARCHAR(20),
		qDate DATE
	);`
	_, err = db.Exec(createPreInformationHistoryTableQuery)
	util.CheckRuntimeError(err, "Information db create err")

	createContainerSpecTableQuery := `
	CREATE TABLE IF NOT EXISTS ContainerSpec (
		id SERIAL PRIMARY KEY,
		inspEqNo VARCHAR(20) NOT NULL,
		inspNo VARCHAR(19) NOT NULL,
		inspStartTime VARCHAR(14),
		inspEndTime VARCHAR(14),
		pckMatch VARCHAR(1),
		inspRsltCD VARCHAR(2),
		detectionCnt VARCHAR(3),
		faultCD VARCHAR(1000),
		inspRsltImgDir VARCHAR(1000),
		qDate DATE
	);
	`
	_, err = db.Exec(createContainerSpecTableQuery)
	util.CheckRuntimeError(err, "Spec db create err")

	createRemarksTableQuery := `
	CREATE TABLE IF NOT EXISTS Remarks (
		remarkId SERIAL PRIMARY KEY,
		inspRemark VARCHAR(1000),
		informationId INT REFERENCES ContainerSpec(id),
		qDate DATE
	);
	`
	_, err = db.Exec(createRemarksTableQuery)
	util.CheckRuntimeError(err, "Remarks db create err")

	createEqInformationTableQuery := `
	CREATE TABLE IF NOT EXISTS EqInformation (
		id SERIAL PRIMARY KEY,
		inspEqNo VARCHAR(20) NOT NULL,
		inspAuto VARCHAR(1) NOT NULL,
		inspName VARCHAR(200),
		inspLoc VARCHAR(200),
		inspContact VARCHAR(20),
		qDate DATE
	);
	`
	_, err = db.Exec(createEqInformationTableQuery)
	util.CheckRuntimeError(err, "EqOperateInfo db create err")

	createEqStateTableQuery := `
	CREATE TABLE IF NOT EXISTS EqState (
		id SERIAL PRIMARY KEY,
		inspEqNo VARCHAR(20) NOT NULL,
		inspEqStatus VARCHAR(1) NOT NULL,
		qDate DATE
	);
	`
	_, err = db.Exec(createEqStateTableQuery)
	util.CheckRuntimeError(err, "EqStateInfo db create err")
}

func DBConn() *sql.DB {
	if db == nil {
		return nil
	}
	return db
}
