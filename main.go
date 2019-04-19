package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

type pathList []struct {
	Path      string     `json:"path"`
	Overrides []override `json:"overrides"`
}

type override struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {

	jsonFile, err := os.Open("paths.json")
	if err != nil {
		log.Fatal(err)
	}

	var paths pathList
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &paths)

	localConfig := &api.Config{Address: os.Getenv("LOCAL_ADDR")}
	localClient, err := api.NewClient(localConfig)
	if err != nil {
		log.Fatal(err)
	}
	localClient.SetToken(os.Getenv("LOCAL_TOKEN"))
	localLogical := localClient.Logical()

	externalConfig := &api.Config{Address: os.Getenv("EXTERNAL_ADDR")}
	externalClient, err := api.NewClient(externalConfig)
	if err != nil {
		log.Fatal(err)
	}
	externalClient.SetToken(os.Getenv("EXTERNAL_TOKEN"))
	externalLogical := externalClient.Logical()

	for _, p := range paths {
		fmt.Printf("Copying %s\n", p.Path)
		externalSecret, err := externalLogical.Read(p.Path)
		if err != nil {
			log.Fatal(err)
		}

		for _, o := range p.Overrides {
			fmt.Printf("--Overriding %s\n", o.Key)
			externalSecret.Data[o.Key] = o.Value
		}

		localLogical.Write(p.Path, externalSecret.Data)

	}

}
