package main

import (
	"log"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	typev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"fmt"
	"strings"
	"errors"
	"k8s.io/apimachinery/pkg/labels"
)

func main {
	services, err := clientset.Core().Services(name).List(api.ListOptions{})
	if err != nil {
		log.Errorf("Get service from kubernetes cluster error:%v", err)
		return
	}

	for _, service := range services.Items {
		if name == "default" && service.GetName() == "kubernetes" {
			continue
		}
		log.Infoln("namespace", name, "serviceName:", service.GetName(), "serviceKind:", service.Kind, "serviceLabels:", service.GetLabels(), service.Spec.Ports, "serviceSelector:", service.Spec.Selector)

		// labels.Parser
		set := labels.Set(service.Spec.Selector)

		if pods, err := clientset.Core().Pods(name).List(api.ListOptions{LabelSelector: set.AsSelector()}); err != nil {
			log.Errorf("List Pods of service[%s] error:%v", service.GetName(), err)
		} else {
			for _, v := range pods.Items {
				log.Infoln(v.GetName(), v.Spec.NodeName, v.Spec.Containers)
			}
		}
	}
}