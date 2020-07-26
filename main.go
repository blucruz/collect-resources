package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// cmd.Execute()
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	api := clientset.CoreV1()
	pods, err := api.Pods("").List(context.TODO(), metav1.ListOptions{})

	for _, pod := range pods.Items {
		fmt.Println(pod.Spec.Containers[0].Name, "--", pod.Spec.Containers[0].Resources.Requests.Cpu())
	}
}
