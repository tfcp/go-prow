package kube

import (
	"prow/library/log"
	"github.com/gogf/gf/frame/g"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 获取k8s配置
func GetKbClient() (*kubernetes.Clientset, error) {
	configPath := g.Config().GetString("kubectl.config")
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Logger.Fatal(err)
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Logger.Fatal(err)
		return nil, err
	}
	return clientset, nil
}
