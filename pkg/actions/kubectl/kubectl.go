package kubectl

// import (
// 	"context"
// 	"fmt"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// )

// type KubectlActionConf struct {
// 	Kubeconfig string `yaml:"kubeconfig" validate:"required,file"`
// 	Operation  string `yaml:"operation" validate:"required,oneof=get"`
// 	Resource   string `yaml:"resource" validate:"required,oneof=pod"`
// 	Namespace  string `yaml:"namespace" validate:"required"`
// 	Name       string `yaml:"name" validate:"required"`
// }

// type KubectlAction struct {
// }

// func NewKubectlAction() *KubectlAction {
// 	return &KubectlAction{}
// }

// func (a *KubectlAction) Name() string {
// 	return "kubectl"
// }

// func (a *KubectlAction) DefaultConf() interface{} {
// 	return KubectlActionConf{
// 		Kubeconfig: "",
// 		Operation:  "get",
// 		Resource:   "pod",
// 		Namespace:  "kube-system",
// 		Name:       "",
// 	}
// }

// func (a *KubectlAction) Run(conf interface{}) error {
// 	c, ok := conf.(KubectlActionConf)
// 	if !ok {
// 		return fmt.Errorf("conf type is not correct")
// 	}
// 	clientConfig, err := clientcmd.BuildConfigFromFlags("", c.Kubeconfig)
// 	if err != nil {
// 		return fmt.Errorf("failed to build client config: %w", err)
// 	}
// 	clientset, err := kubernetes.NewForConfig(clientConfig)
// 	if err != nil {
// 		return fmt.Errorf("failed to create clientset from client config: %w", err)
// 	}
// 	switch c.Resource {
// 	case "pod":
// 		podClient := clientset.CoreV1().Pods(c.Namespace)
// 		switch c.Operation {
// 		case "get":
// 			pod, err := podClient.Get(context.TODO(), c.Name, metav1.GetOptions{})
// 			if err != nil {

// 			}
// 		}
// 	}
// 	return nil
// }
