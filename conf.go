package go_common

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"gopkg.in/yaml.v2"
)

var confUrl string

func init() {
	confUrl = os.Getenv("CONF_URL")
}

func GetAppConfig(appName string, v interface{}) {

	environment := os.Getenv("GO_ENV")
	if environment == "" {
		environment = "dev"
	}

	resp, err := http.Get(fmt.Sprintf("%s/%s.yml", confUrl, appName))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll data error : %v", err)
	}

	err = yaml.Unmarshal([]byte(body), v)
	if err != nil {
		log.Fatalf("Unmarshal data error : %v", err)
	}

	log.Printf("%v", v)

	resp, err = http.Get(fmt.Sprintf("%s/%s-%s.yml", confUrl, appName, environment))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	err = yaml.Unmarshal([]byte(body), v)

	if err != nil {
		log.Fatalf("Unmarshal data error : %v", err)
	}

	log.Printf("%v", v)

}
