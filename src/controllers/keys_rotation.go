package controllers

import (
	config "RemoteValues/src"
	k8s "RemoteValues/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewKeyRotationBody struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Service string `json:"service"`
}

func ReturnConfigs() k8s.ConfigMapConfig {
	configCM := k8s.ConfigMapConfig{
		Name:      config.ConfigGlobal().ConfigMapName,
		Namespace: config.ConfigGlobal().Namespace,
	}
	return configCM
}

func SetNewKeyRotation(w *gin.Context) {
	var dataBody NewKeyRotationBody
	configCM := ReturnConfigs()
	if err := w.Bind(&dataBody); err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	configMapData := k8s.ConfigMapData{
		Key:   dataBody.Key,
		Value: dataBody.Value,
	}
	result, err := k8s.UpdateConfigMapHandler(configCM, &configMapData)
	if err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error(), "resource": result.Name})
		return
	}

	_, errorSendRequest := k8s.SendRequestContainerHandler(&k8s.DeploymentConfig{
		Namespace: configCM.Namespace,
		Name:      dataBody.Service,
	}, configMapData)

	if errorSendRequest != nil {
		w.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errorSendRequest.Error(), "resource": result.Name})
		return
	}

	w.JSON(http.StatusAccepted, gin.H{"status": "updated", "resource": result.Name})
	return
}

func ReturnKeysRotation(w *gin.Context) {
	configCM := ReturnConfigs()
	listConfigs, errorListConfig := k8s.ReturnConfigMapHandler(configCM)
	if errorListConfig != nil {
		w.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": errorListConfig.Error()})
		return
	}
	w.JSON(http.StatusOK, &listConfigs.Data)
	return
}
