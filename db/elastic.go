package db

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Elastic() {

// HTTP client timeout
var timeout = time.Duration(300 * time.Millisecond)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}

	// Check to see if ElasticSearch server is up
	elasticResponse, err := client.Get(elastic_url)
	if elasticResponse != nil {
		jsonStr, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(dateStamp, h.Hostname, "INFO elasticsearch endpoint:", elastic_url)
		if overwrite == "enable" {
			reqDelete, err := http.NewRequest("DELETE", elastic_url, nil)
			if username != "undef" {
				reqDelete.SetBasicAuth(username, password)
			}
			respDelete, err := http.DefaultClient.Do(reqDelete)
			fmt.Println(dateStamp, h.Hostname, "DELETE elasticsearch type status:", respDelete.Status)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		reqPost, err := http.NewRequest("POST", elastic_url, bytes.NewBuffer(jsonStr))
		if password != "undef" {
			reqPost.SetBasicAuth(username, password)
		}
		reqPost.Header.Set("Content-Type", "application/json")

		clientReq := &http.Client{}
		respPost, err := clientReq.Do(reqPost)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer respPost.Body.Close()
		fmt.Println(dateStamp, h.Hostname, "POST json elasticsearch type status:", respPost.Status)
		postBody, _ := ioutil.ReadAll(respPost.Body)
		fmt.Println(dateStamp, h.Hostname, "POST response body:", string(postBody))
	} else {
		fmt.Println(dateStamp, h.Hostname, "FAIL unable to connect to elasticeearch server:", "http://"+elastic_host+":"+elastic_port)
	}
}
