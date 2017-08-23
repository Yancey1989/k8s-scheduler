package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"encoding/json"

	"github.com/topicai/candy"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func get_admin_config() (*restclient.Config, error) {
	caData, err := ioutil.ReadFile("./test/ca.pem")
	candy.Must(err)
	certData, err := ioutil.ReadFile("./test/admin.pem")
	candy.Must(err)
	keyData, err := ioutil.ReadFile("./test/admin-key.pem")
	candy.Must(err)
	config := clientcmdapi.NewConfig()
	config.Clusters["dlnel"] = &clientcmdapi.Cluster{
		Server: "https://kubernetes:443",
		CertificateAuthorityData: caData,
	}
	config.AuthInfos["dlnel"] = &clientcmdapi.AuthInfo{
		ClientCertificateData: certData,
		ClientKeyData:         keyData,
	}
	config.Contexts["dlnel"] = &clientcmdapi.Context{
		Cluster:  "dlnel",
		AuthInfo: "dlnel",
	}
	config.CurrentContext = "dlnel"
	clientBuilder := clientcmd.NewNonInteractiveClientConfig(*config, "dlnel", &clientcmd.ConfigOverrides{}, nil)
	return clientBuilder.ClientConfig()
}

func main() {
	schedulerName := flag.String("scheduler-name", "my-scheduler", "Your custom scheduler name")
	flag.Parse()

	// creates the in-cluster config
	config, err := get_admin_config()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		for _, item := range nodes.Items {
			fmt.Println(item.ObjectMeta.Name)
		}
		fmt.Printf("There ara %d nodes in the cluster \n", len(nodes.Items))
		// list all pods
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for _, item := range pods.Items {
			if item.Spec.SchedulerName == *schedulerName &&
				item.Spec.NodeName == "" {
				nodeName := nodes.Items[rand.Intn(len(nodes.Items))].ObjectMeta.Name
				fmt.Printf("Start assign %s\n", item.ObjectMeta.Name)
				binding := fmt.Sprintf(`{
					"apiVersion": "v1",
					"kind": "Binding",
					"metadata": {
						"name": "%s"
					},
					"target": {
						"apiVersion": "v1",
						"kind": "Node",
						"name": "%s"
					}
				}`, item.ObjectMeta.Name, nodeName)
				bindingObj := &v1.Binding{}
				json.Unmarshal([]byte(binding), bindingObj)
				err := clientset.CoreV1().Pods(item.ObjectMeta.Namespace).Bind(bindingObj)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("Assign %s to %s\n", item.ObjectMeta.Name, nodeName)
			}

		}

		time.Sleep(10 * time.Second)
	}
}
