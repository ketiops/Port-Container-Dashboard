package controller

import (
	"db-server/model"
	"db-server/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetContainerReview(w http.ResponseWriter, r *http.Request) {
	inspNoHeader := r.Header.Get("inspNo")
	db := model.DBConn()

	rows, err := db.Query("SELECT r.remarkId AS remark_id, r.inspRemark, r.informationId, r.qDate FROM ContainerSpec s JOIN Remarks r ON s.id = r.informationId WHERE s.inspNo = $1;", inspNoHeader)
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var remarks []Remark
	for rows.Next() {
		var r Remark
		err = rows.Scan(&r.RemarkID, &r.InspRemark, &r.InformationID, &r.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		remarks = append(remarks, r)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB rows") {
		return
	}

	if len(remarks) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No Reviews"}`))
		return
	}

	jsonResponse, err := json.Marshal(remarks)
	if util.CheckHttpError(w, err, "Check JSON") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func AppendContainerReview(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()
	var remark Remark
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&remark)

	if util.CheckHttpError(w, err, "Decoding JSON error") {
		return
	}

	query := `
		INSERT INTO Remarks (inspRemark, informationId, qDate) 
		VALUES ($1, $2, $3) RETURNING remarkId`

	err = db.QueryRow(query, remark.InspRemark, remark.InformationID, remark.QDate).Scan(&remark.RemarkID)
	if util.CheckHttpError(w, err, "Save DB error") {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
