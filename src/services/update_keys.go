package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BodyKeyUpdate struct {
	DataValue string `json:"dataValue"`
	Key       string `json:"key"`
}

type KeyUpdate struct {
	ConfigMapData BodyKeyUpdate
	PodInfo       PodInfo
}

func UpdateKeysRequest(keyValue KeyUpdate) (bool, error) {
	bodyUpdate, errorBodyUpdate := json.Marshal(keyValue.ConfigMapData)
	if errorBodyUpdate != nil {
		return false, errorBodyUpdate
	}
	response, errorResponse := http.Post(
		fmt.Sprintf("http://%s:%d/remoteconfig", keyValue.PodInfo.IP, keyValue.PodInfo.Port),
		"application/json; charset=utf-8",
		bytes.NewBuffer(bodyUpdate),
	)
	if errorResponse != nil {
		return false, errorResponse
	}

	_ = response.Body.Close()
	return true, nil
}
