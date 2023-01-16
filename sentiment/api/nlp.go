package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Endpoint(endpoint string) string {
	c := http.Client{Timeout: time.Duration(5) * time.Second}
	response, err := c.Get(endpoint)
	if err != nil {
		log.Print(err.Error())
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err.Error())
	}
	return string(responseData)
}
