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
	"encoding/json"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/marshyski/plural/cloud"
	"github.com/marshyski/plural/config"
	"github.com/marshyski/plural/containers"
	"github.com/marshyski/plural/data"
	"github.com/marshyski/plural/db"
	"github.com/marshyski/plural/network"
	"github.com/marshyski/plural/packages"
	"github.com/marshyski/plural/system"
)

var (
	configFlag = flag.String("config", "", "  Set configuration path, defaults are ['./','/etc/plural/conf','/opt/plural/conf','./conf']")
	daemonFlag = flag.Bool("daemon", false, "  Run in daemon mode")
)

func init() {
	flag.StringVar(configFlag, "c", "", "  Set configuration path, defaults are ['./','/etc/plural/conf','/opt/plural/conf','./conf']")
	flag.BoolVar(daemonFlag, "d", false, "  Run in daemon mode")
}

var usage = `Usage: plural [options] <args>

    -d, --daemon     Run in daemon mode
    -c, --config     Set configuration path, defaults are ['./','/etc/plural/conf','/opt/plural/conf','./conf']

Example:       plural -d -c /opt/plural/conf

Documentation:  https://github.com/marshyski/plural/blob/master/README.md

`

func main() {

	flag.Usage = func() {
		fmt.Println(usage)
	}

	flag.Parse()

	for {
		var (
			d  data.PluralJSON
			wg sync.WaitGroup
		)
		wg.Add(17)
		go func() {
			defer wg.Done()
			network.Conns(&d)
		}()
		go func() {
			defer wg.Done()
			network.DNS(&d)
		}()
		go func() {
			defer wg.Done()
			network.DomainName(&d)
		}()
		go func() {
			defer wg.Done()
			network.IP(&d)
		}()
		go func() {
			defer wg.Done()
			network.IPRoutes(&d)
		}()
		go func() {
			defer wg.Done()
			network.IPTables(&d)
		}()
		go func() {
			defer wg.Done()
			cloud.Aws(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Deb(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Pip(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Rpm(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Gem(&d)
		}()
		go func() {
			defer wg.Done()
			containers.Docker(&d)
		}()
		go func() {
			defer wg.Done()
			system.Audit(&d)
		}()
		go func() {
			defer wg.Done()
			system.Stats(&d)
		}()
		go func() {
			defer wg.Done()
			system.Users(&d)
		}()
		go func() {
			defer wg.Done()
			system.UsersLoggedIn(&d)
		}()
		go func() {
			defer wg.Done()
			d.Lastrun = time.Now().Format(time.RFC3339)
		}()
		wg.Wait()

		if !*daemonFlag {
			j, err := json.Marshal(d)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			fmt.Println(string(j))
			break
		}

		db.Elastic(&d)

		time.Sleep(time.Duration(config.ConfigInt("interval")) * time.Second)
	}
}
