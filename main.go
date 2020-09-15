package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	//"gopkg.in/yaml.v2"
)

var needInjectDeployments []v1.Deployment = make([]v1.Deployment, 7, 8)

func main() {
	println("go start")

	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		obj := map[string]string{"message": "pong"}
		bytes, err := json.Marshal(obj)
		judgeError(err)
		w.Write(bytes)
		w.WriteHeader(200)
	})
	http.HandleFunc("/latency/{number}", latencyProcess)
	http.HandleFunc("/pods/inject-latency-container", injectLatencyContainer)

	println("go listen start")

	http.ListenAndServe(":8092", nil)

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

func getAllNetCard() []string {
	cmd := exec.Command("bash", "-c", "ifconfig -s | awk 'NR>2{print $1}'")
	out, err := cmd.Output()
	judgeError(err)

	return strings.Split(string(out), "\n")
}

func latencyProcess(w http.ResponseWriter, req *http.Request) {
	//todo check float or no

	latencyStr := req.FormValue("number")
	netcards := getAllNetCard()
	latencyNum, err := strconv.Atoi(latencyStr)
	judgeError(err)
	println(latencyStr)

	for _, netcard := range netcards {
		setNetcardLatency(netcard, latencyNum)
	}

	obj := map[string]string{"latency": latencyStr}
	bytes, err := json.Marshal(obj)
	judgeError(err)
	w.Write(bytes)
	w.WriteHeader(200)
}

func setNetcardLatency(netcard string, latencyNum int) {
	deleteErr := exec.Command("tc",
		"qdisc", "delete", "dev",
		netcard, "root", "netem",
		"delay").Start()
	judgeError(deleteErr)

	addErr := exec.Command("tc",
		"qdisc", "add", "dev",
		netcard, "root", "netem",
		"delay", strconv.Itoa(latencyNum/2)+"ms").Start()

	println("iota: " + strconv.Itoa(latencyNum/2))
	judgeError(addErr)
}

func initClient() *kubernetes.Clientset {
	kubeconfig := flag.String("kubeconfig",
		"config", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	judgeError(err)

	clientset, err := kubernetes.NewForConfig(config)
	judgeError(err)
	return clientset
}

func injectLatencyContainer(w http.ResponseWriter, req *http.Request) {
	clientset := initClient()

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	judgeError(err)

	for _, pod := range pods.Items {
		fmt.Printf("%s %s\n", pod.GetName(), pod.GetCreationTimestamp())
	}

	//todo: write with tikv deploy name
	deployments, err := clientset.AppsV1beta1().Deployments(metav1.NamespaceDefault).
		List(context.TODO(), metav1.ListOptions{})
	judgeError(err)

	//todo: put in static dir
	sidecar_container_yaml, err := ioutil.
		ReadFile("manifests/latency-setting-sidecar-single.yaml")
	judgeError(err)

	patchList, err := clientset.AppsV1beta1().
		Deployments("default").Patch(context.TODO(), "",
		types.ApplyPatchType, sidecar_container_yaml, metav1.PatchOptions{})
	judgeError(err)
	println(patchList)

	respObj := map[string]string{
		"pods":                   pods.String(),
		"deployments":            deployments.String(),
	}
	bytes, err := json.Marshal(respObj)
	judgeError(err)
	w.Write(bytes)
	w.WriteHeader(200)
}


func filterDeployment(deployment v1.Deployment) bool {
	return strings.Contains(strings.ToLower(deployment.Name), "tikv")
}

func judgeError(err error) {
	if err != nil {
		panic(err)
	}
}