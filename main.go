// (C) Copyright 2015 Timothy Marcinowski
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//  http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"flag"
	"fmt"
	"plural/cloud"
	"plural/config"
	"plural/data"
	"plural/docker"
	"plural/network"
	"plural/packages"
	"plural/system"
	"time"
)

// Command-line flags
var configFlag = flag.String("config", "", "  Set configuration path, default is /opt/plural/conf")
var daemonFlag = flag.Bool("daemon", false, "  Run in daemon mode")
var outputFlag = flag.String("output", "", "  Output JSON file in a directory specified")

func init() {
	flag.StringVar(configFlag, "c", "", "  Set configuration path, default is /opt/plural/conf")
	flag.BoolVar(daemonFlag, "d", false, "  Run in daemon mode")
	flag.StringVar(outputFlag, "o", "", "  Output JSON file in a directory specified")
}

var usage = `Usage: plural [options] <args>

    -d, --daemon       Run in daemon mode
    -c, --config     Set configuration path, default path is /opt/plural/conf
    -o, --output     Output JSON file in a directory specified


Example:       plural -d -c /opt/plural/conf -o /tmp

Documentation:  https://github.com/marshyski/plural/blob/master/README.md

`

func main() {

	flag.Usage = func() {
		fmt.Println(usage)
	}

	flag.Parse()

	for {

		data.PluralJSON = make(map[string]interface{})
		//fmt.Printf("%s %s INFO %s SHA256 checksum is %x\n", dateStamp, h.Hostname, file.Name(), hash.Sum(nil))
		//network.Conns()
		network.DNS()
		network.DomainName()
		network.IP()
		network.IPRoutes()
		network.IPTables()
		cloud.Aws()
		packages.Deb()
		packages.Pip()
		packages.Rpm()
		packages.Gem()
		docker.Containers()
		system.Audit()
		system.Stats()
		system.Users()
		system.UsersLoggedIn()

		//fmt.Println(data.PluralJSON["Platform"].(string))
		for k, v := range data.PluralJSON {
			fmt.Println(k, v)
		}

		if !*daemonFlag {
			break
		}

		// Sleep / interval time for daemon
		time.Sleep(time.Duration(config.ConfigInt("interval")) * time.Second)

	}
}
