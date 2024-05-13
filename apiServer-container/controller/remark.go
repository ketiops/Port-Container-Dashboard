package controller

import (
	"io"
	"net/http"
	"server/util"
)

func AppendContainerReview(w http.ResponseWriter, r *http.Request) {
	url := "http://" + dbServerHost + ":" + dbServerPort + "/container/spec/review"

	req, err := http.NewRequest("POST", url, r.Body)
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

func GetContainerReview(w http.ResponseWriter, r *http.Request) {
	inspNoHeader := r.Header.Get("inspNo")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/container/spec/review"

	req, err := http.NewRequest("GET", url, nil)
	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}

	req.Header.Set("inspNo", inspNoHeader)

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
