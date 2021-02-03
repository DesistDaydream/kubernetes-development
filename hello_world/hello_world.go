package main

import (
	"context"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InOrOut 判断当前环境是在集群内部，还是集群外部
func InOrOut() string {
	// 如果容器内具有环境变量 KUBERNETES_SERVICE_HOST 且不为空，则当前代码是在容器内运行，否则是在集群外部运行
	if h := os.Getenv("KUBERNETES_SERVICE_HOST"); h != "" {
		return "inCluster"
	}
	return "outCluster"
}

// Deployment 获取指定 namespace 下所有的 deployment 对象
func get(clientset *kubernetes.Clientset, namespace string) {
	// 获取指定 名称空间 下所有的 deployment 对象
	deployments, _ := clientset.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
	for i, deploy := range deployments.Items {
		fmt.Printf("%d -> %s\n", i+1, deploy.Name)
	}
}

func main() {
	var config *rest.Config
	// 根据代码所在环境，决定如何创建一个连接集群所需的配置。
	switch InOrOut() {
	case "inCluster":
		// 根据容器内的 /var/run/secrets/kubernetes.io/serviceaccount/ 目录下的 token 与 ca.crt 文件创建一个用于连接集群的配置。
		config, _ = rest.InClusterConfig()
	case "outCluster":
		// 根据指定的 kubeconfig 文件创建一个用于连接集群的配置，/root/.kube/config 为 kubectl 命令所用的 config 文件
		config, _ = clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
		// 注意，clientcmd.BuildConfigFromFlags() 内部实际上也是有调用 rest.InClusterConfig() 的逻辑，只要满足条件即可。条件如下：
		// 若第二个参数为空的话，则会直接调用 rest.InClusterConfig()
	}

	// 根据 BuildConfigFromFlags 创建的配置，返回一个可以连接集群的指针
	clientset, _ := kubernetes.NewForConfig(config)

	// 获取指定 namespace 下所有的 deployment 对象
	get(clientset, "kube-system")
}
