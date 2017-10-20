# plural


#### Table of Contents

1. [Overview](#overview)
	* [Example JSON Output](#example-json-output)
2. [Install Dependencies](#install-dependencies)
    * [Server](#server)
    * [Client](#client)
3. [Command-Line Arguments](#command-line-arguments)
4. [Configuration](#configuration)
5. [Platforms Tested On](#platforms-tested-on)
6. [Screenshots](#screenshots)


## Overview

A lightweight system information collector for storing data in ElasticSearch or Stdout.  Great for keeping track of elastic environments and auditing configurations.

Resources gathered if applicable:

- RHEL Audit Rules
- CPU Count
- Disk Stats
- Docker Containers
- Domain Name
- EC2 Instance Metadata
- Ruby Gems
- Hostname
- IP Address
- IPTables Rules
- IP Routes
- Kernel Version
- Load Averages
- Memory Stats
- RPM / Deb Packages
- Python Pip Packages
- OS Platform
- OS Family
- OS Version
- TCP 4 Listening
- TCP 6 Listening
- Timezone
- Uptime
- Users
- Users Logged In
- Virtualization Role
- Virtualization System


### Example JSON Output

    {
       "audit_rules": [
         "-w /var/log/audit/ -p wa -k LOG_audit",
         "-w /etc/audit/auditd.conf -p wa -k CFG_audit",
         "-w /etc/rc.d/init.d/auditd -p wa -k CFG_audit",
         "-w /etc/sysconfig/auditd -p wa -k CFG_audit",
         "-w /etc/audit/audit.rules -p wa -k CFG_audit",
         "-w /etc/localtime -p wa -k time-change,CFG_system"
       ],
       "cpu_count": 4,
       "diskfree_gb": 6,
       "disktotal_gb": 8,
       "diskused_gb": 19,
       "dns_nameserver": [
         "8.8.8.8",
         "8.8.4.4"
       ],
       "docker": [
         "image=dockerui/dockerui:latest, command=./dockerui, port='9000 9000 tcp 0.0.0.0'",
         "image=robloach/forge-lamp:latest, command=supervisord, port='22 0 tcp  3306 0 tcp  80 0 tcp '",
         "image=robloach/forge-lamp:latest, command=supervisord, port='22 0 tcp  3306 0 tcp  80 0 tcp '",
         "image=robloach/forge-lamp:latest, command=supervisord, port='22 49159 tcp 0.0.0.0 3306 49160 tcp 0.0.0.0 80 49161 tcp 0.0.0.0'",
         "image=robloach/forge-lamp:latest, command=supervisord, port='22 0 tcp  3306 0 tcp  80 0 tcp '"
       ],
       "domain": "ec2.internal",
       "ec2_ami_id": "ami-bc8131d4",
       "ec2_availability_zone": "us-east-1b",
       "ec2_instance_id": "i-1b8cc9cc",
       "ec2_instance_type": "t1.micro",
       "ec2_profile": "default-paravirtual",
       "ec2_public_ip4": "54.145.182.91",
       "ec2_security_groups": "default",
       "environment": "dev",
       "gem": [
         "arr-pm-0.0.9",
         "backports-3.6.4",
         "cabin-0.7.1",
         "childprocess-0.5.6",
         "clamp-0.6.4",
         "ffi-1.9.8",
         "fpm-1.3.3",
         "json-1.8.2"
       ],
       "hostname": "ip-10-28-229-205",
       "ipaddress": "10.28.229.205",
       "iptables": [
         "ACCEPT     tcp  --  anywhere             anywhere             state RELATED,ESTABLISHED",
         "DROP       all  -f  anywhere             anywhere            ",
         "ACCEPT     tcp  --  localhost            anywhere             tcp dpt:webcache",
         "ACCEPT     tcp  --  localhost            anywhere             tcp dpt:webcache",
         "DROP       tcp  --  anywhere             anywhere             tcp dpt:webcache",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp dpt:http state NEW,ESTABLISHED",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp dpt:http limit: avg 25/min burst 100",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp spt:http state ESTABLISHED",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp spt:webcache state ESTABLISHED"
       ],
       "ip_route": [
         "default via 192.168.1.1 dev eth0 ",
         "172.17.0.0/16 dev docker0  proto kernel  scope link  src 172.17.42.1 ",
         "192.168.1.0/24 dev eth0  proto kernel  scope link  src 192.168.1.10 "
       ],
       "kernel_version": "2.6.32-431.29.2.el6.x86_64",
       "lastrun": "2015-05-21T23:29:49-04:00",
       "load15": 0,
       "load1": 0,
       "load5": 0,
       "memoryfree_gb": 2,
       "memorytotal_gb": 16,
       "memoryused_gb": 14,
       "os": "linux",
       "packages": [
         "acl-2.2.49-6.el6.x86_64",
         "acpid-1.0.10-2.1.el6.x86_64",
         "alsa-lib-1.0.22-3.el6.x86_64",
         "atk-1.30.0-1.el6.x86_64",
         "attr-2.4.44-7.el6.x86_64",
         "audit-2.2-4.el6_5.x86_64",
         "audit-libs-2.2-4.el6_5.x86_64",
         "authconfig-6.1.12-13.el6.x86_64",
         "avahi-libs-0.6.25-15.el6.x86_64",
         "b43-openfwwf-5.2-4.el6.noarch",
         "basesystem-10.0-4.el6.noarch",
         "bash-4.1.2-15.el6_5.2.x86_64",
         "binutils-2.20.51.0.2-5.36.el6.x86_64",
         "bzip2-1.0.5-7.el6_0.x86_64",
         "bzip2-libs-1.0.5-7.el6_0.x86_64",
         "ca-certificates-2014.1.98-65.0.el6_5.noarch",
         "cairo-1.8.8-3.1.el6.x86_64",
         "centos-release-6-5.el6.centos.11.2.x86_64"
       ],
       "pip": [
         "distribute-0.6.10",
         "Flask-0.10.1",
         "Flask-Limiter-0.7.4",
         "Flask-SSLify-0.1.4",
         "gunicorn-19.1.1",
         "iniparse-0.3.1",
         "itsdangerous-0.24",
         "Jinja2-2.7.3",
         "limits-1.0.4",
         "MarkupSafe-0.23",
         "pycurl-7.19.0",
         "pygpgme-0.1",
         "pyOpenSSL-0.14",
         "python-magic-0.4.6",
         "PyYAML-3.11",
         "six-1.9.0",
         "SQLAlchemy-0.9.8",
         "urlgrabber-3.9.1",
         "Werkzeug-0.10.4",
         "yum-metadata-parser-1.1.2",
         "yum-presto-0.4.4"
       ],
       "platform": "centos",
       "platform_family": "rhel",
       "platform_verison": "6.5",
       "tcp4_listen": [
         "0.0.0.0:9200 /usr/lib/jvm/java-1.7.0-openjdk-1.7.0.75.x86_64/jre/bin/java",
         "0.0.0.0:8080 /usr/bin/python",
         "0.0.0.0:80 /opt/kibana-4.0.2-linux-x64/node/bin/node",
         "0.0.0.0:9300 /usr/lib/jvm/java-1.7.0-openjdk-1.7.0.75.x86_64/jre/bin/java",
         "0.0.0.0:22 /usr/sbin/sshd",
         "127.0.0.1:25 /usr/libexec/postfix/master"
       ],
       "tcp6_listen": [
         "0000:0000:0000:0000:0000:0000:0001:0000:25 /usr/libexec/postfix/master",
         "0000:0000:0000:0000:0000:0000:0000:0000:8090 /opt/influxdb/versions/0.8.8/influxdb",
         "0000:0000:0000:0000:0000:0000:0000:0000:8099 /opt/influxdb/versions/0.8.8/influxdb",
         "0000:0000:0000:0000:0000:0000:0000:0000:8083 /opt/influxdb/versions/0.8.8/influxdb",
         "0000:0000:0000:0000:0000:0000:0000:0000:8086 /opt/influxdb/versions/0.8.8/influxdb",
         "0000:0000:0000:0000:0000:0000:0000:0000:22 /usr/sbin/sshd",
         "0000:0000:0000:0000:0000:0000:0000:0000:3000 /usr/sbin/grafana-server"
       ],
       "timezone": "UTC",
       "uptime_days": 9,
       "users": [
         "root:x:0:0:root:/root:/bin/bash",
         "adm:x:3:4:adm:/var/adm:/sbin/nologin",
         "shutdown:x:6:0:shutdown:/sbin:/sbin/shutdown",
         "nginx:x:998:997:Nginx web server:/var/lib/nginx:/sbin/nologin",
         "varnish:x:997:996:Varnish Cache:/var/lib/varnish:/sbin/nologin"
       ],
       "users_loggedin": [
         "root-pts/0",
         "timski-pts/1"
       ],
       "virtualization_role": "guest",
       "virtualization_system": "xen"
    }


ElasticSearch terminology:

http://elasticsearch:9200/index/type

 Plural terminology:

http://elasticsearch:9200/environment/ipaddress

Agent Run Time:

The agent runs every five minutes, and post real-time data to ElasticSearch.

** If you were to delete all hosts in the environment nightly.   If the agent is running and the server is up, it will populate the inventory currently with only running hosts and their data.  This works very well in elastic compute environments.


## Install Dependencies

### Server

 - ElasticSearch 5.x
 - [elasticsearch-http-user-auth Plugin](https://github.com/elasticfence/elasticsearch-http-user-auth) (Optional)
 - Java 8.x / OpenJDK 8
 - Kibana 5.x


 **Last step is configure ElasticSearch mappings for all indexes to not be analyzed:**


       curl -XPUT -H "Content-Type: application/json" http://localhost:9200/_template/template_1 -d '
       {
          "template":"*",
          "settings":{
             "index.refresh_interval":"5s"
          },
          "mappings":{
             "_default_":{
                "_all":{
                   "enabled":true
                },
                "dynamic_templates":[
                   {
                      "fields":{
                         "match":"*",
                         "match_mapping_type":"string",
                         "mapping":{
                            "type":"keyword",
                            "index":true,
                            "norms":false
                         }
                      }
                   }
                ]
             }
          }
       }'


### Client

 - Go 1.9
 - Make
 

## Command-Line Arguments

No flags / arguments will do a one-time run and produce a JSON file in the current path of the binary

    -d, --daemon     Run in daemon mode
    -c, --config     Set configuration path, defaults are ['./','/etc/plural/conf','/opt/plural/conf','./conf']


## Configuration

Configurations can be written in YAML, JSON or TOML.

*/opt/plural/conf/plural.yaml*
*DEFAULT  values if no config is present*

    # ElasticSearch Indexer
    host: localhost
    port: 9200

    # ElasticSearch Index Name
    ## This can be anything, it could be aws, softlayer, prod, staging
    environment: dev

    # Interval of agent runs in seconds
    ## Default is every five minutes
    interval: 300

    # Username if http-basic plugin is enabled
    username:

    #Password if http-basic plugin is enabled
    password:

    # Secure true enables HTTPS instead of HTTP)
    secure: false


## Platforms Tested On

 - CentOS/RHEL 7.x
 - Fedora 20
 - Ubuntu 16
 - Mac OS X 16.7.0


## Screenshots

![Dashboard View](https://s3.amazonaws.com/timski-pictures/dashview.png)

![Search View](https://s3.amazonaws.com/timski-pictures/searchview.png)

![Docker View](https://s3.amazonaws.com/timski-pictures/dockerview.png)
