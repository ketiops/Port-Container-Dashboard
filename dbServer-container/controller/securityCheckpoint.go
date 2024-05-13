package controller

import (
	"db-server/model"
	"db-server/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllCheckpoint(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT id, inspEqNo, inspAuto, inspName, inspLoc, inspContact, qDate FROM EqInformation")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var eqOperateInfos []EqInformation
	for rows.Next() {
		var eoi EqInformation
		err = rows.Scan(&eoi.ID, &eoi.InspEqNo, &eoi.InspAuto, &eoi.InspName, &eoi.InspLoc, &eoi.InspContact, &eoi.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		eqOperateInfos = append(eqOperateInfos, eoi)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB rows") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eqOperateInfos)
}

func GetAllCheckpointState(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT id, inspEqNo, inspEqStatus, qDate FROM EqState")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var eqStateInfos []EqState
	for rows.Next() {
		var esi EqState
		err = rows.Scan(&esi.ID, &esi.InspEqNo, &esi.InspEqStatus, &esi.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		eqStateInfos = append(eqStateInfos, esi)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB rows") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eqStateInfos)
}

func CreateCheckpointState(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	var eqStateInfos EqState
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&eqStateInfos)
	if util.CheckHttpError(w, err, "Decoding JSON error") {
		return
	}

	query := `
		INSERT INTO EqState 
		(inspEqNo, inspEqStatus, qDate) 
		VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(query, eqStateInfos.InspEqNo, eqStateInfos.InspEqStatus, eqStateInfos.QDate).Scan(&eqStateInfos.ID)
	if util.CheckHttpError(w, err, "Save DB error") {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
