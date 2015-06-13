# plural

Linux agent to send host-based information about the server to ElasticSearch and Kibana

----------

**Overview:**

Ever wanted a dynamic inventory, search, and visualization into your server environments?  

Every wanted to minimize logging into servers, or limiting production access?

Think CMDB like-features without the B.S.   Create graphs/lists on high disk usage/CPU utilization, kernel versions, Docker containers, TCP4/6 listening ports, AWS inventory, installed packages (rpm, dpkg, pip, gem), etc.  The agent is a signally golang compiled binary able to run across platforms without runtime dependencies.

**Example JSON Output:**

    {
       "audit_rules": [
         "-w /var/log/audit/ -p wa -k LOG_audit",
         "-w /etc/audit/auditd.conf -p wa -k CFG_audit",
         "-w /etc/rc.d/init.d/auditd -p wa -k CFG_audit",
         "-w /etc/sysconfig/auditd -p wa -k CFG_audit",
         "-w /etc/audit/audit.rules -p wa -k CFG_audit",
         "-w /etc/localtime -p wa -k time-change,CFG_system"
       ],
       "diskfree": "6.7GB",
       "disktotal": "8.5GB",
       "diskused": "19",
       "dns_nameserver": [
         "8.8.8.8",
         "8.8.4.4"
       ],
       "docker": [
         "dockerui/dockerui:latest, ./dockerui, '9000 9000 tcp 0.0.0.0'",
         "robloach/forge-lamp:latest, supervisord, '22 0 tcp  3306 0 tcp  80 0 tcp '",
         "robloach/forge-lamp:latest, supervisord, '22 0 tcp  3306 0 tcp  80 0 tcp '",
         "robloach/forge-lamp:latest, supervisord, '22 49159 tcp 0.0.0.0 3306 49160 tcp 0.0.0.0 80 49161 tcp 0.0.0.0'",
         "robloach/forge-lamp:latest, supervisord, '22 0 tcp  3306 0 tcp  80 0 tcp '"
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
       "kernelversion": "2.6.32-431.29.2.el6.x86_64",
       "lastrun": "2015-05-21T17:54",
       "load15": "0",
       "load1": "0",
       "load5": "0",
       "memoryfree": "133MB",
       "memorytotal": "604MB",
       "memoryused": "67",
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
       "platformfamily": "rhel",
       "platformverison": "6.5",
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
       "uptime": "471h18m59s",
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
       "virtualizationrole": "guest",
       "virtualizationsystem": "xen"
    }


ElasticSearch terminology:

http://elasticsearch:9200/index/type
 
 Plural terminology: 

http://elasticsearch:9200/environment/hostname

Agent Run Time:

The agent runs every five minutes, it will delete the host out of the environment and post real-time data at the five minute interval.

** If you were to delete all hosts in the environment nightly.   If the agent is running and the server is up, it will populate the inventory currently with only running hosts and their data.  This works very well in dynamic compute environments.

    # Delete all hosts out of the environment
    curl -XDELETE http://elasticsearch:9200/environment/*

----------

**Install Dependencies:**

*Server:*

 - ElasticSearch (Listening on IPv4 0.0.0.0 not 127.0.0.1 or :::) 
 - [elasticsearch-http-basic Plugin](https://github.com/Asquera/elasticsearch-http-basic) (Optional) 
 - Java 7.x / OpenJDK 7
 - Kibana


 **Last step is configure ElasticSearch mappings for all indexes to not be analyzed:**
 
 
       curl -XPUT localhost:9200/_template/template_1 -d '
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
                      "string_fields":{
                         "match":"*",
                         "match_mapping_type":"string",
                         "mapping":{
                            "type":"string",
                            "index":"not_analyzed",
                            "omit_norms":true
                         }
                      }
                   }
                ]
             }
          }
       }'


*Client:*

 - Packages coming soon, for now `go build` project
 - `mkdir -p /opt/plural/{bin,conf}`
 - Move compiled binary to /opt/plural/bin/

----------

**Build Dependencies:**

    go get github.com/spf13/viper
    go get github.com/shirou/gopsutil
    go get github.com/dustin/go-humanize
    go get github.com/fsouza/go-dockerclient
    go get github.com/drael/GOnetstat

----------

**Configuration (YAML, JSON or TOML):**

*/opt/plural/conf/plural.yaml*

    # ElasticSearch Indexer
    elastic_host: 54.145.182.91
    elastic_port: 9200
    
    # ElasticSearch Index Name
    ## This can be anything, it could be aws, softlayer, prod, staging
    environment: dev

    # Interval of agent runs in seconds
    ## Default is every five minutes
    interval: 300
    
    # Username if http-basic plugin is enabled
    username: admin
    
    #Password if http-basic plugin is enabled
    password: admin_pw

*DEFAULT  values if no config is present*

    elastic_host : localhost
    elastic_port : 9200
    environment : dev
    interval: 300
    username: undef
    password: undef

----------

**Tested Againsted:**

 - CentOS/RHEL 6.x
 - Fedora 20
 - Ubuntu 14
 - Mac OS X 13.4.0

----------

**Screenshots**

![Dashboard View](https://s3.amazonaws.com/timski-pictures/dashview.png)

![Search View](https://s3.amazonaws.com/timski-pictures/searchview.png)

![Docker View](https://s3.amazonaws.com/timski-pictures/dockerview.png)
