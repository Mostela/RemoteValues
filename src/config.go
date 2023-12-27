package src

import "os"

type GlobalConfigKeyRotation struct {
	Namespace     string
	ConfigMapName string
	Debug         bool
}

func ConfigGlobal() GlobalConfigKeyRotation {
	config := GlobalConfigKeyRotation{
		Namespace:     os.Getenv("K8S_NAMESPACE"),
		ConfigMapName: os.Getenv("K8S_CONFIGMAP_NAME"),
		Debug:         os.Getenv("DEBUG") == "true",
	}
	return config
}
