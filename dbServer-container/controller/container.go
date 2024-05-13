package controller

import (
	"database/sql"
	"db-server/model"
	"db-server/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPreContainerInfo(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	var inputInfo PreInformation
	err := json.NewDecoder(r.Body).Decode(&inputInfo)
	if util.CheckHttpError(w, err, "Failed to decode the request body") {
		return
	}

	if inputInfo.TruckNo == "" {
		inputInfo.TruckNo = ""
	}
	if inputInfo.TypeNo == "" {
		inputInfo.TypeNo = ""
	}

	var exist int
	err = db.QueryRow("SELECT id FROM PreInformation WHERE inspEqNo = $1 AND cntrNo = $2 AND truckNo = $3 AND typeNo = $4 AND qDate = $5",
		inputInfo.InspEqNo, inputInfo.CntrNo, inputInfo.TruckNo, inputInfo.TypeNo, inputInfo.QDate).Scan(&exist)

	if err == sql.ErrNoRows {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No matching data"}`))
		return
	} else if err != nil {
		util.CheckHttpError(w, err, "Error querying the database")
		return
	}

	_, err = db.Exec("INSERT INTO PreInformationHistory (inspEqNo, cntrNo, truckNo, typeNo, qDate) VALUES ($1, $2, $3, $4, $5)",
		inputInfo.InspEqNo, inputInfo.CntrNo, inputInfo.TruckNo, inputInfo.TypeNo, inputInfo.QDate)
	if util.CheckHttpError(w, err, "Failed to insert data") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Data exists"}`))
}

func GetAllPreContainersInfo(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT id, inspEqNo, cntrNo, truckNo, typeNo, qDate FROM PreInformation")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}

	var informations []PreInformation
	for rows.Next() {
		var info PreInformation
		err = rows.Scan(&info.ID, &info.InspEqNo, &info.CntrNo, &info.TruckNo, &info.TypeNo, &info.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		informations = append(informations, info)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB Rows") {
		return
	}

	if len(informations) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not exist data"}`))
		return
	}

	jsonResponse, err := json.Marshal(informations)
	if util.CheckHttpError(w, err, "Check JSON") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func CreateContainerSpec(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	var spec ContainerSpec
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&spec)
	if util.CheckHttpError(w, err, "Decoding JSON error") {
		return
	}

	query := `
		INSERT INTO ContainerSpec 
		(inspEqNo, inspNo, inspStartTime, inspEndTime, pckMatch, inspRsltCD, detectionCnt, faultCD, inspRsltImgDir, qDate) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err = db.QueryRow(query, spec.InspEqNo, spec.InspNo, spec.InspStartTime, spec.InspEndTime, spec.PckMatch, spec.InspRsltCD, spec.DetectionCnt, spec.FaultCD, spec.InspRsltImgDir, spec.QDate).Scan(&spec.ID)
	if util.CheckHttpError(w, err, "Save DB error") {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAllContainerSpec(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT * FROM ContainerSpec")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var specs []ContainerSpec
	for rows.Next() {
		var spec ContainerSpec
		err = rows.Scan(&spec.ID, &spec.InspEqNo, &spec.InspNo, &spec.InspStartTime, &spec.InspEndTime, &spec.PckMatch, &spec.InspRsltCD, &spec.DetectionCnt, &spec.FaultCD, &spec.InspRsltImgDir, &spec.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		specs = append(specs, spec)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(specs)
}
