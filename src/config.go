package src

import "os"

type GlobalConfigKeyRotation struct {
	Namespace         string
	ConfigMapName     string
	Debug             bool
	EndpointSetValues string
}

func ConfigGlobal() GlobalConfigKeyRotation {
	endpointConfig := os.Getenv("ENDPOINT_SET_VALUES")
	if endpointConfig == "" {
		endpointConfig = "remoteconfig"
	}

	config := GlobalConfigKeyRotation{
		Namespace:         os.Getenv("K8S_NAMESPACE"),
		ConfigMapName:     os.Getenv("K8S_CONFIGMAP_NAME"),
		Debug:             os.Getenv("DEBUG") == "true",
		EndpointSetValues: os.Getenv("ENDPOINT_SET_VALUES"),
	}
	return config
}
