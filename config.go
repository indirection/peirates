//Build API configuration (svc account token, namespace, API server) -- automated prereq for other steps

// Locations of the svc account and token for i/o reads
// account token -  /run/secrets/kubernetes.io/serviceaccount
// name space
package main

import (
//	"os/exec"
	"fmt"
	"io/ioutil"
)

type config_Info struct {
	token []byte
	ca_crt []byte
	namespaces []byte
	my_config string
	err_Read error
	err_Write error
}

func Builder() {

	//creating config_Info type and storing in a variable
	var config_InfoVars = config_Info{}

	// Reading token file and storing in variable token
	config_InfoVars.token, config_InfoVars.err_Read = ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/token")

	//Error message If statement based on failure to read the file
	if config_InfoVars.err_Read != nil {
		fmt.Println("Token location error: ", config_InfoVars.err_Read)
	}

	// Reading namespaces file and storing in variable namespaces
	config_InfoVars.namespaces, config_InfoVars.err_Read = ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/namespace")
	if config_InfoVars.err_Read != nil {
		fmt.Println("Namespaces location error", config_InfoVars.err_Read)
	}

	//Reading Ca.Crt File and storing in variable ca_crt
	config_InfoVars.ca_crt, config_InfoVars.err_Read = ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/ca.crt")
	if config_InfoVars.err_Read != nil {
		fmt.Println("Ca.Crt location error: ", config_InfoVars.err_Read)
	}

	//Creating variable to write to file output
	config_InfoVars.my_config = fmt.Sprintf("KUBE_TOKEN=%s\nKUBE_NAMESPACES=%s\nKUBE_CA_CRT=%s", string(config_InfoVars.token), string(config_InfoVars.namespaces), string(config_InfoVars.ca_crt))

	//Output config to specified location
	ioutil.WriteFile(".peirates.conf", []byte(config_InfoVars.my_config), 0700)
}
// Main Fucntion of the program

func main() {
	Builder()
}
