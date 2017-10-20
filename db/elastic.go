package db

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/marshyski/plural/config"
	"github.com/marshyski/plural/data"
)

var (
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c = &http.Client{
		Timeout:   500 * time.Millisecond,
		Transport: tr,
	}
	host   = config.ConfigStr("elastic_host")
	port   = config.ConfigStr("elastic_port")
	env    = config.ConfigStr("environment")
	sec    = config.ConfigBool("secure")
	scheme string
)

func init() {
	if sec {
		scheme = "https"
	} else {
		scheme = "http"
	}
}

func Elastic(d *data.PluralJSON) {
	url := fmt.Sprintf("%s://%s:%s/%s/%s", scheme, host, port, env, strings.Replace(d.Ipaddress, ".", "", 3))
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(d)
	_, err := c.Post(url, "application/json; charset=UTF-8", body)
	if err != nil {
		log.Printf("Error: %s", err)
	}
}
