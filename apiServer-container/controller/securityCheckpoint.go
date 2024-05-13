package controller

import (
	"bytes"
	"io"
	"net/http"
	"server/util"
)

func GetAllCheckpoint(w http.ResponseWriter, r *http.Request) {
	url := "http://" + dbServerHost + ":" + dbServerPort + "/checkpoint"

	req, err := http.NewRequest("GET", url, nil)
	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if util.CheckHttpError(w, err, "Check Client Do") {
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if util.CheckHttpError(w, err, "Check Copying Response") {
		return
	}
}

func GetAllCheckpointState(w http.ResponseWriter, r *http.Request) {
	url := "http://" + dbServerHost + ":" + dbServerPort + "/checkpoint/state"

	req, err := http.NewRequest("GET", url, nil)
	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if util.CheckHttpError(w, err, "Check Client Do") {
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if util.CheckHttpError(w, err, "Check Copying Response") {
		return
	}
}

func CreateCheckpointState(w http.ResponseWriter, r *http.Request) {
	url := "http://" + dbServerHost + ":" + dbServerPort + "/checkpoint/state"
	bodyBytes, err := io.ReadAll(r.Body)
	if util.CheckHttpError(w, err, "Reading body") {
		return
	}
	bodyReader := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest("POST", url, bodyReader)
	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if util.CheckHttpError(w, err, "Check Client Do") {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to save data on remote server"))
	}
}
