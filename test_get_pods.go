package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

// This program lists the pods in a cluster equivalent to
//
// kubectl get pods
//
func main() {
	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace")

	// Bootstrap k8s configuration from local 	Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

	// print pods
	for i, pod := range pods.Items {
		fmt.Println("\nPod: ", i)
		fmt.Println("pod ip            : ", pod.Status.PodIP)
		fmt.Println("name              : ", pod.GetName())
		fmt.Println("namespace         : ", pod.GetNamespace())
		fmt.Println("cluster name      : ", pod.GetClusterName())
		fmt.Println("labels            : ", pod.GetLabels())

	}

	services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
	for i, service := range services.Items {
		fmt.Println("\nService: ", i)
		fmt.Println("name             : ", service.GetName())
		fmt.Println("namespace        : ", service.GetNamespace())
		fmt.Println("cluster ip       : ", service.Spec.ClusterIP)
		fmt.Println("selector         : ", service.Spec.Selector)

	}
}
