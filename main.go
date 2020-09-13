package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
)

var needInjectDeployments []v1.Deployment = make([]v1.Deployment, 7, 8)

func main() {
	println("go start")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/getPodList", printPodList)

	println("go listen start")

	r.Run(":8092")


	//
	//kubeconfigf := flag.String("kubeconfig", "/Users/mhausenblas/.kube/config", "absolute path to the kubeconfig file")
	//flag.Parse()
	//
	//kubeconfig := filepath.Join(
	//	os.Getenv("HOME"), ".kube", "config",
	//)
	//clientcmd.BuildConfigFromFlags()
	//
	//config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	//judgeError(err)
	//
	//clientset, err := kubernetes.NewForConfig(config)



	//clientset.CoreV1().ConfigMaps().List()
	//dList, dErr := clientset.AppsV1().Deployments("namespace").List(
	//	nil,
	//	metav1.ListOptions{},
	//)
	//judgeError(dErr)
	//
	//println(dList)



	//for _, deployment := range dList.Items {
	//	if filterDeployment(deployment) {
	//		needInjectDeployments = append(needInjectDeployments, deployment)
	//	}
	//}
	//
	//response := v1beta1.AdmissionResponse{Allowed: true}
	//println(response)

	//v1beta1.AdmissionResponse{Allowed: true, Patch: "aaa", PatchType: func() {}}
	//v1beta1.PatchTypeJSONPatch
	//v1beta
	//spec := dList.Items[0].Spec
	//specStr := spec.String()
	//spec.Template.
	//
	//for _, deploy := range dList.Items {
	//
	//}

}

//func updateK8s(old :v1.oldDeployment, new :v1) {
//
//}

func printPodList(c *gin.Context) {
	kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	judgeError(err)

	clientset, err := kubernetes.NewForConfig(config)
	judgeError(err)

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	judgeError(err)

	for _, pod := range pods.Items {
		fmt.Printf("%s %s\n", pod.GetName(), pod.GetCreationTimestamp())
	}
}


func filterDeployment(deployment v1.Deployment) bool {
	return strings.Contains(strings.ToLower(deployment.Name), "tikv")
}

func judgeError(err error) {
	if err != nil {
		panic(err)
	}
}