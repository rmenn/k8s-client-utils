package version

import (
	"k8s.io/client-go/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func GetApiVersion(client *kubernetes.ClientSet) (string, error) {
	version, err := client.ServerVersion()
	if err != nil {
		return "", err
	}
	return version.GitVersion, nil
}
