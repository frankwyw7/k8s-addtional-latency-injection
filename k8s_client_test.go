package main

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
	"k8s.io/client-go/kubernetes"
)

func TestK8s(t *testing.T) {

	config, err := clientcmd.BuildConfigFromFlags("", "")
	judgeError(err)

	client, err := kubernetes.NewForConfig(config)

	deploymentList, err := client.AppsV1beta1().
		Deployments("default").
		List(nil, v1.ListOptions{AllowWatchBookmarks: true})

	println(deploymentList)

	//r, _ := http.NewRequest("GET", "/", nil)
	//w := httptest.NewRecorder()
	//beego.BeeApp.Handlers.ServeHTTP(w, r)
	//
	//beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())


	//Convey("Subject: Test Station Endpoint\n", t, func() {
	//	Convey("Status Code Should Be 200", func() {
	//		So(w.Code, ShouldEqual, 200)
	//	})
	//	Convey("The Result Should Not Be Empty", func() {
	//		So(w.Body.Len(), ShouldBeGreaterThan, 0)
	//	})
	//})
}