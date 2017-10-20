package db

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
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
		Timeout:   1 * time.Second,
		Transport: tr,
	}
	host   = config.ConfigStr("host")
	port   = config.ConfigStr("port")
	env    = config.ConfigStr("environment")
	sec    = config.ConfigBool("secure")
	user   = config.ConfigStr("username")
	pass   = config.ConfigStr("password")
	scheme string
	url    string
)

const ct = "application/json; charset=UTF-8"

func init() {
	if sec {
		scheme = "https"
	} else {
		scheme = "http"
	}
}

func Elastic(d *data.PluralJSON) {
	if url == "" {
		buf := new(bytes.Buffer)
		buf.WriteString(scheme)
		buf.WriteString("://")
		buf.WriteString(host)
		buf.WriteString(":")
		buf.WriteString(port)
		buf.WriteString("/")
		buf.WriteString(env)
		buf.WriteString("/")
		buf.WriteString(strings.Replace(d.Ipaddress, ".", "", 3))
		url = buf.String()
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(d)

	if user != "" && pass != "" {
		r, err := http.NewRequest("POST", url, body)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
		r.Header.Add("Content-Type", ct)
		r.SetBasicAuth(user, pass)
		_, err = c.Do(r)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}

		return
	}

	_, err := c.Post(url, ct, body)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
}
