package main

import (
	"fmt"
	"encoding/json"
	//"time"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	//"bytes"
	//"os"
	//"io"
	"os/exec"
	"io/ioutil"
	//"archive/tar"
	//"path/filepath"
	"./commandModels"
)


func getConfig () (output string){
	kube_conf, err := ioutil.ReadFile("/root/.kube/config")
	if err != nil {
		return string(err.Error())
	}else {
		output = string(kube_conf)
		return output
	}
}


func getKubeToken () (output string){
	// cmd := exec.Command("kubeadm", "token", "create", "--print-join-command")
	/*cmd := exec.Command("ls", "-alt")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Run()
	//cmd.Start()

	stdin.Write([]byte("12 34 +p\n"))

	out := make([]byte, 1024)

	if err != nil {

		output = string(err.Error())
	}else {
		n, _ := stdout.Read(out)

		oot := out[:n]

		output = string(oot)
	}*/

	out, err := exec.Command("kubeadm", "token", "create", "--print-join-command").Output()
	if err != nil {
		log.Fatal(err)
	}
	output = string(out)

	return output
}

func startSSH () {
	out, err := exec.Command("gotty", "-w", "/bin/sh").Output()

	if err != nil {
		log.Fatal(err)
	}
	output := string(out)

	fmt.Printf(output)
}

func certsTar () {
	out, err := exec.Command("tar", "-cvf", "taryMcTarball.tar", "/etc/kubernetes/pki").Output()

	if err != nil {
		log.Fatal(err)
	}
	output := string(out)

	fmt.Printf(output)
}

func certsGzip () {
	out, err := exec.Command("gzip", "taryMcTarball.tar").Output()

	if err != nil {
		log.Fatal(err)
	}
	output := string(out)

	fmt.Printf(output)
}

func deleteCertsGzip () {
	out, err := exec.Command("rm", "taryMcTarball.tar.gz").Output()

	if err != nil {
		log.Fatal(err)
	}
	output := string(out)

	fmt.Printf(output)
}

func startMaster0 () {
	out, err := exec.Command("/usr/bin/kube-start.sh").Output()

	if err != nil {
		log.Fatal(err)
	}
	output := string(out)

	fmt.Printf(output)
}

func joinCommand (cmnd string) {
	out, err := exec.Command(cmnd).Output()

	if err != nil {
		log.Fatal(err)
	}
	output := string(out)

	fmt.Printf(output)
}

func getKubeJoin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cmdJson := commandModels.JoinCommand{}

	json.NewDecoder(r.Body).Decode(&cmdJson)

	cmd := cmdJson.JoinCmd

	go joinCommand(cmd)

	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Kube join initiated")
}


func getConf(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
    	fmt.Fprint(w, getConfig())
}

func getToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
    	fmt.Fprint(w, getKubeToken())
}

func getCertTar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	certsTar()
	certsGzip()
	w.Header().Set("Content-Type", "octet-stream")
    	http.ServeFile(w, r, "taryMcTarball.tar.gz")
	deleteCertsGzip()
}

func getSSH(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	go startSSH()
	fmt.Fprint(w, "SSH started at Device on Port 8080 in your web browser.")
}

func getMasterKube(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	go startMaster0()
	w.Header().Set("Content-Type", "text/plain")
    	fmt.Fprint(w, "Master Kube0 bootstrap started")
}

func main() {
	router := httprouter.New()
    	router.GET("/get_config", getConf)
	router.GET("/get_join_token", getToken)
	router.GET("/get_cert_tar.tar.gz", getCertTar)
	router.GET("/getSSH", getSSH)
	router.GET("/startMasterKube", getMasterKube)
	router.POST("/joinKube", getKubeJoin)

	log.Fatal(http.ListenAndServe(":5000", router))

}
