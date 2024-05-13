package controller

import (
	"bytes"
	"db-server/util"
	"io"
	"net/http"
	"os"
)

func GetTosContainer(w http.ResponseWriter, r *http.Request) {
	tosIp := os.Getenv("TOS_HOST")
	tosPort := os.Getenv("TOS_PORT")
	tosPath := os.Getenv("TOS_PATH")
	url := "http://" + tosIp + ":" + tosPort + "/" + tosPath

	bodyBytes, err := io.ReadAll(r.Body)
	if util.CheckHttpError(w, err, "Reading body") {
		return
	}
	bodyReader := bytes.NewBuffer(bodyBytes)

	req, err := http.NewRequest("POST", url, bodyReader)
	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
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
