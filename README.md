# plural

Linux agent to send host-based facts about the server to ElasticSearch and Kibana

----------

**Overview:**

Ever wanted a dynamic data inventory, search and data visualization into your server environments?  Think CMDB like-features without the B.S.   Create graphs on high disk usage/CPU utilization, kernel versions, platforms, AWS inventory, etc.  The agent is a signally golang compiled binary able to run across platforms without runtime dependencies.  

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

*Client:*

 - Packages coming soon, for now `go build` project
 - mkdir -p /opt/plural/{bin,conf}
 - Move compiled binary to /opt/plural/bin/

----------

**Build Dependencies:**

    go get github.com/spf13/viper
    go get github.com/shirou/gopsutil

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
        "diskfree": "6833799168",
        "disktotal": "8454070272",
        "diskused": "19",
        "domain": "ec2.internal",
        "ec2_ami_id": "ami-bc8131d4",
        "ec2_availability_zone": "us-east-1b",
        "ec2_instance_id": "i-1b8cc9cc",
        "ec2_instance_type": "t1.micro",
        "ec2_profile": "default-paravirtual",
        "ec2_security_groups": "default",
        "hostname": "ip-10-28-229-205",
        "ipaddress": "10.28.229.205",
        "kernelversion": "2.6.32-431.29.2.el6.x86_64",
        "load15": "3",
        "load1": "3",
        "load5": "3",
        "memoryfree": "129888000",
        "memorytotal": "604480000",
        "memoryused": "63",
        "os": "linux",
        "platform": "centos",
        "platformfamily": "rhel",
        "platformverison": "6.5",
        "timezone": "UTC",
        "uptime": "6168",
        "virtualizationrole": "guest",
        "virtualizationsystem": "xen"
    }

  
----------

**Tested Againsted:**

 - CentOS/RHEL 6.x
 - Mac OS X 13.4.0

----------

**Screenshots**

![Dashboard Widgets](https://s3.amazonaws.com/timski-pictures/dashboard.png)

![Kernel Search](https://s3.amazonaws.com/timski-pictures/kernel-search.png)
