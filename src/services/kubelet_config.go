package services

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func KubeletConfig() *kubernetes.Clientset {
	var config *rest.Config
	var err error
	if os.Getenv("DEBUG") == "true" {
		config, err = clientcmd.BuildConfigFromFlags("", "/home/joao/.kube/config")
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		panic(err)
	}
	client, _ := kubernetes.NewForConfig(config)
	return client
}
