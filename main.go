package main

import (
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	ctrl "sigs.k8s.io/controller-runtime"
)

var needInjectDeployments []v1.Deployment = make([]v1.Deployment, 7, 8)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

	config, err := clientcmd.BuildConfigFromFlags("", "")
	judgeError(err)

	client, err := kubernetes.NewForConfig(config)
	dList, derr := client.AppsV1().Deployments("namespace").List(
		nil,
		metav1.ListOptions{},
	)
	judgeError(derr)

	for _, deployment := range dList.Items {
		if(filterDeployment(deployment)) {
			_ = append(needInjectDeployments, deployment)
		}
	}

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

func filterDeployment(deployment v1.Deployment) bool {
	return strings.Contains(strings.ToLower(deployment.Name), "tikv")
}

func judgeError(err error) {
	if err != nil {
		panic(err)
	}
}