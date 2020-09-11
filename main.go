package main

import (
	"github.com/gin-gonic/gin"
	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
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

	config := rest.Config{Host: "45.32.39.122", APIPath: "api/v1/"}

	client, err := kubernetes.NewForConfig(&config)
	judgeError(err)

	dList, derr := client.AppsV1().Deployments("namespace").List(
		nil,
		metav1.ListOptions{},
	)
	judgeError(derr)

	for _, deployment := range dList.Items {
		if filterDeployment(deployment) {
			needInjectDeployments = append(needInjectDeployments, deployment)
		}
	}

	response := v1beta1.AdmissionResponse{Allowed: true}
	println(response)

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

func filterDeployment(deployment v1.Deployment) bool {
	return strings.Contains(strings.ToLower(deployment.Name), "tikv")
}

func judgeError(err error) {
	if err != nil {
		panic(err)
	}
}