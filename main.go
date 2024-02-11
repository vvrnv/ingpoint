package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get ingresses hosts
	ingresses, err := clientset.NetworkingV1().Ingresses("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	uniqueHosts := make(map[string]bool)

	for _, ingress := range ingresses.Items {
		for _, rule := range ingress.Spec.Rules {
			uniqueHosts[rule.Host] = true
		}
	}

	for host := range uniqueHosts {
		fmt.Println(host)
	}

	if len(uniqueHosts) == 0 {
		fmt.Println("No ingress resources found")
	}
}
