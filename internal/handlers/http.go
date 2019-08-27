package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gs-open-provider/op-auth-admin-gateway/internal/logger"
)

func handlePostBasicRequest(url string, body interface{}) (*http.Response, error) {
	var inInterface map[string]string
	in, _ := json.Marshal(body)
	json.Unmarshal(in, &inInterface)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(in))
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return resp, err
}
