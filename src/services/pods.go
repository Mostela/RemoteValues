package services

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentConfig struct {
	Namespace string
	Name      string
}

type PodInfo struct {
	IP   string
	Name string
	Port int32
}

func GetContainerPodsHandler(config *DeploymentConfig) ([]PodInfo, error) {
	k8sHandler := KubeletConfig()
	var containerList []PodInfo = make([]PodInfo, 0)
	podsList, errListPods := k8sHandler.CoreV1().Pods(config.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf(`app=%s`, config.Name),
	})
	if errListPods != nil {
		return nil, errListPods
	}
	for _, pod := range podsList.Items {
		var statusContainers = true
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Terminated != nil || !containerStatus.Ready {
				statusContainers = false
			}
		}
		if statusContainers {
			containerList = append(containerList, PodInfo{
				IP:   pod.Status.PodIP,
				Port: pod.Spec.Containers[0].Ports[0].ContainerPort,
				Name: pod.Name,
			})
		}
	}
	return containerList, errListPods
}

func SendRequestContainerHandler(config *DeploymentConfig, dataConfigMap ConfigMapData) (bool, error) {
	containerList, errorList := GetContainerPodsHandler(config)
	if errorList != nil {
		return false, errorList
	}
	for _, pod := range containerList {
		status, err := UpdateKeysRequest(KeyUpdate{
			ConfigMapData: BodyKeyUpdate{
				DataValue: dataConfigMap.Value,
				Key:       dataConfigMap.Key,
			},
			PodInfo: pod,
		})
		if err != nil {
			fmt.Printf("%s", err.Error())
			return false, err
		}
		fmt.Printf("%s", status)
	}
	return true, nil
}
