# plural

Linux agent to send host-based facts about the server to ElasticSearch and Kibana

----------

**Overview:**

Ever wanted a dynamic data inventory, search and data visualization into your server environments?  Think CMDB like-features without the B.S.   Create graphs on high disk usage/CPU utilization, kernel versions, AWS inventory, query installed packages, etc.  The agent is a signally golang compiled binary able to run across platforms without runtime dependencies.  

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
 - Java 7.x / OpenJDK 7
 - Kibana


 **Last step is configure ElasticSearch mappings for all indexes to not analyzed:**
 

       curl -XPUT localhost:9200/_template/template_1 -d '
       {
         "template" : "*",
         "settings" : {
         "index.refresh_interval" : "5s"
       },
       "mappings" : {
          "_default_" : {
          "_all" : {"enabled" : true},
          "dynamic_templates" : [ {
         "string_fields" : {
          "match" : "*",
          "match_mapping_type" : "string",
          "mapping" : {
          "type" : "string", "index" : "not_analyzed", "omit_norms" : true
          }
         }
        }]
       }
       }
       }'


*Client:*

 - Packages coming soon, for now `go build` project
 - mkdir -p /opt/plural/{bin,conf}
 - Move compiled binary to /opt/plural/bin/

----------

**Build Dependencies:**

    go get github.com/spf13/viper
    go get github.com/shirou/gopsutil
    go get github.com/dustin/go-humanize

----------

**Configuration (YAML, JSON or TOML):**

*/opt/plural/conf/plural.yaml*

    # ElasticSearch Indexer
    elastic_host: 54.145.182.91
    elastic_port: 9200
    
    # ElasticSearch Index Name
    ## This can be anything, it could be aws, softlayer, prod, staging
    environment: dev

*DEFAULT  values if no config is present*

    elastic_host : localhost
    elastic_port : 9200
    environment : dev

----------

**Example JSON Output:**

    {
       "diskfree": "6.7GB",
       "disktotal": "8.5GB",
       "diskused": "19",
       "domain": "ec2.internal",
       "ec2_ami_id": "ami-bc8131d4",
       "ec2_availability_zone": "us-east-1b",
       "ec2_instance_id": "i-1b8cc9cc",
       "ec2_instance_type": "t1.micro",
       "ec2_profile": "default-paravirtual",
       "ec2_public_ip4": "54.145.182.91",
       "ec2_security_groups": "default",
       "hostname": "ip-10-28-229-205",
       "ipaddress": "10.28.229.205",
       "kernelversion": "2.6.32-431.29.2.el6.x86_64",
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
         "centos-release-6-5.el6.centos.11.2.x86_64",
         ""
       ],
       "platform": "centos",
       "platformfamily": "rhel",
       "platformverison": "6.5",
       "timezone": "UTC",
       "uptime": "351730",
       "virtualizationrole": "guest",
       "virtualizationsystem": "xen"
    }
  
----------

**Tested Againsted:**

 - CentOS/RHEL 6.x
 - Fedora 20
 - Ubuntu 14
 - Mac OS X 13.4.0

----------

**Screenshots**

![Dashboard Widgets](https://s3.amazonaws.com/timski-pictures/dashboard.png)

![Kernel Search](https://s3.amazonaws.com/timski-pictures/kernel-search.png)

