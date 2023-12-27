package services

import (
	"context"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMapConfig struct {
	Namespace string
	Name      string
}

type ConfigMapData struct {
	Key   string
	Value string
}

func ReturnConfigMapHandler(config ConfigMapConfig) (*v1.ConfigMap, error) {
	k8sHandler := KubeletConfig()
	cmClient := k8sHandler.CoreV1().ConfigMaps(config.Namespace)

	resultGET, errorCMGET := cmClient.Get(context.TODO(), config.Name, metav1.GetOptions{})
	return resultGET, errorCMGET
}

func UpdateConfigMapHandler(config ConfigMapConfig, data *ConfigMapData) (*v1.ConfigMap, error) {
	k8sHandler := KubeletConfig()
	cmClient := k8sHandler.CoreV1().ConfigMaps(config.Namespace)

	resultGET, errorCMGET := cmClient.Get(context.TODO(), config.Name, metav1.GetOptions{})
	if errorCMGET != nil {
		return nil, errorCMGET
	}
	if resultGET.Data != nil {
		resultGET.Data[data.Key] = data.Value
	} else {
		resultGET.Data = map[string]string{
			data.Key: data.Value,
		}
	}
	cm := v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      config.Name,
			Namespace: config.Namespace,
		},
		Data: resultGET.Data,
	}

	result, errorCM := cmClient.Update(context.TODO(), &cm, metav1.UpdateOptions{})
	return result, errorCM
}
