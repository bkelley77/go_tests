package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// only public fields that are capitalized are encoded to json
type MyK8sPod struct {
	Pod_ip       string `json:"pod_ip"`
	Cluster_name string `json:"cluster_name"`
}

// showing how to pass and modify a map to a function
func test1(pods map[string]MyK8sPod) {
	p := MyK8sPod{Pod_ip: "88.88.88.88", Cluster_name: "cluster88"}
	pods["88.88.88.88"] = p
}

func main() {

	pods := make(map[string]MyK8sPod)

	p := MyK8sPod{Pod_ip: "99.99.99.99", Cluster_name: "cluster99"}
	pods["99.99.99.99"] = p

	// passing a map to a function
	test1(pods)

	// showing how to iterate through a map
	for key, element := range pods {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	fmt.Println("pod map converted to JSON: ")
	var jsonData []byte
	jsonData, err := json.Marshal(pods)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

}
