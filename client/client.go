package client

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func Client(path string) (*kubernetes.ClientSet, error) {
	conf, err := genClientConfig(path)
	if err != nil {
		return nil, fmt.Errorf("unable to render client : %v", err)
	}
	return kuberneted.NewForConfig(conf)
}

func genClientConfig(path string) (*rest.Config, error) {
	if path != "" {
		return clientcmd.BuildConfigFromFlags("", path)
	}
	return rest.InClusterConfig()
}
