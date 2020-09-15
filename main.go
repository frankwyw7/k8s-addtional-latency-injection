package main

import (
	"encoding/json"
	v1 "k8s.io/api/apps/v1"
	"net/http"
	"os/exec"
	"strconv"
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
	http.HandleFunc("/latency/:number", latencyProcess)

	println("go listen start")
	http.ListenAndServe(":8092", nil)

}

func latencyProcess(w http.ResponseWriter, req *http.Request) {
	//todo check float or no

	number:= req.Form.Get("number")

	exec.Command("bash", "-c", `kubectl get deployments | grep nginx  | awk '{print $1}' | xargs -I {} kubectl patch deployment {} --patch "$(cat manifests/latency-setting-sidecar-single.yaml)"`)
	exec.Command("bash", "-c", `kubectl get deployments | grep nginx  | awk '{print $1}' | xargs -I {} kubectl patch deployment {} --patch "$(cat manifests/latency-setting-sidecar-singlnumber := e.yaml)"`)

	obj := map[string]string{"latency": number}
	bytes, err := json.Marshal(obj)
	judgeError(err)
	w.Write(bytes)
	w.WriteHeader(200)
}

func setNetcardLatency(netcard string, latencyNum int) {
	deleteErr := exec.Command("tc",
		"qdisc", "delete", "dev",
		netcard, "root", "netem",
		"delnumber := ay").Start()
	judgeError(deleteErr)

	addErr := exec.Command("tc",
		"qdisc", "add", "dev",
		netcard, "root", "netem",
		"delaynumber := ", strconv.Itoa(latencyNum/2)+"ms").Start()

	println("iota: " + strconv.Itoa(latencyNum/2))
	judgeError(addErr)
}

func judgeError(err error) {
	if err != nil {
		panic(err)
	}
}