package data

type PluralJSON struct {
	AuditRules           []string `json:"audit_rules,omitempty"`
	CPUCount             int32    `json:"cpu_count,omitempty"`
	Diskfree             string   `json:"diskfree_gb,omitempty"`
	Disktotal            string   `json:"disktotal_gb,omitempty"`
	Diskused             string   `json:"diskused_gb,omitempty"`
	DNSNameserver        []string `json:"dns_nameserver,omitempty"`
	Docker               []string `json:"docker,omitempty"`
	Domain               string   `json:"domain,omitempty"`
	Ec2AmiID             string   `json:"ec2_ami_id,omitempty"`
	Ec2AvailabilityZone  string   `json:"ec2_availability_zone,omitempty"`
	Ec2InstanceID        string   `json:"ec2_instance_id,omitempty"`
	Ec2InstanceType      string   `json:"ec2_instance_type,omitempty"`
	Ec2Profile           string   `json:"ec2_profile,omitempty"`
	Ec2PublicIP4         string   `json:"ec2_public_ip4,omitempty"`
	Ec2SecurityGroups    []string `json:"ec2_security_groups,omitempty"`
	Environment          string   `json:"environment,omitempty"`
	Gem                  []string `json:"gem,omitempty"`
	Hostname             string   `json:"hostname,omitempty"`
	IPRoute              []string `json:"ip_route,omitempty"`
	Ipaddress            string   `json:"ipaddress,omitempty"`
	Iptables             []string `json:"iptables,omitempty"`
	Kernelversion        string   `json:"kernelversion,omitempty"`
	Lastrun              string   `json:"lastrun,omitempty"`
	Load1                float64  `json:"load1,omitempty"`
	Load15               float64  `json:"load15,omitempty"`
	Load5                float64  `json:"load5,omitempty"`
	Memoryfree           string   `json:"memoryfree_gb,omitempty"`
	Memorytotal          string   `json:"memorytotal_gb,omitempty"`
	Memoryused           string   `json:"memoryused_gb,omitempty"`
	Os                   string   `json:"os,omitempty"`
	Packages             []string `json:"packages,omitempty"`
	Pip                  []string `json:"pip,omitempty"`
	Platform             string   `json:"platform,omitempty"`
	Platformfamily       string   `json:"platformfamily,omitempty"`
	Platformverison      string   `json:"platformverison,omitempty"`
	TCP4Listen           []string `json:"tcp4_listen,omitempty"`
	TCP6Listen           []string `json:"tcp6_listen,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Uptime               int64    `json:"uptime,omitempty"`
	Users                []string `json:"users,omitempty"`
	UsersLoggedin        []string `json:"users_loggedin,omitempty"`
	Virtualizationrole   string   `json:"virtualizationrole,omitempty"`
	Virtualizationsystem string   `json:"virtualizationsystem,omitempty"`
}
