package main

import (
	"log"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// k8s 控制器运行时 client-go
// https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client/config
func main() {
	// GetConfig() 创建一个用于连接集群 APIServer 的 *rest.Configs。
	// 如果设置了 --kubeconfig 标志，则使用指定的 kubeconfig 文件；否则将假定该程序在集群中运行，并使用集群内部体统的 kubeconfig
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(cfg)
}
