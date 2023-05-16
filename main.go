package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/user/.kube/config", "Expecting flag, default location, Comment")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// handle error
	}
	clientset, err := kubernetes.NewForConfig(config)
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	fmt.Println(pods)
}
