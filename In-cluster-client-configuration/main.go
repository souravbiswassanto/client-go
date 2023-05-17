package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("There is an Error -> %s\n", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("There is an Error -> %s\n", err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}
	// Getting Deployments
	deployment, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("There is an Error -> %s\n", err.Error())
	}
	for _, deploy := range deployment.Items {
		fmt.Printf("%s\n", deploy.Name)
	}
	for {

	}
}
