package kubectl

import (
	"context"
	"fmt"

	"github.com/decadevvv/miniwf/pkg/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubectlActionConf struct {
	Kubeconfig string `yaml:"kubeconfig" validate:"required,file"`
	Operation  string `yaml:"operation" validate:"required,oneof=get"`
	Resource   string `yaml:"resource" validate:"required,oneof=pod"`
	Namespace  string `yaml:"namespace" validate:"required"`
	Name       string `yaml:"name" validate:"required"`
}

var KubectlAction = core.Action{
	Name: "kubectl",
	Doc:  "use kubectl to manage resources in kubernetes cluster",
	DefaultConf: KubectlActionConf{
		Kubeconfig: "",
		Operation:  "get",
		Resource:   "pod",
		Namespace:  "kube-system",
		Name:       "",
	},
	Run: func(conf interface{}) (interface{}, error) {
		c, ok := conf.(KubectlActionConf)
		if !ok {
			return nil, fmt.Errorf("conf type is not correct")
		}
		clientConfig, err := clientcmd.BuildConfigFromFlags("", c.Kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("failed to build client config: %w", err)
		}
		clientset, err := kubernetes.NewForConfig(clientConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create clientset from client config: %w", err)
		}
		switch c.Resource {
		case "pod":
			podClient := clientset.CoreV1().Pods(c.Namespace)
			switch c.Operation {
			case "get":
				return podClient.Get(context.TODO(), c.Name, metav1.GetOptions{})
			default:
				return nil, fmt.Errorf("undefined operation %s", c.Operation)
			}
		default:
			return nil, fmt.Errorf("undefined resource %s", c.Resource)
		}
	},
}
