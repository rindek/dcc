package main

type composev23 V23

type V23 struct {
	Version  string                `yaml:"version"`
	Services map[string]V23Service `yaml:"services"`
	Networks map[string]V23Network `yaml:"networks,omitempty"`
	Volumes  map[string]V32Volume  `yaml:"volumes,omitempty"`
}

type V23Network struct {
	Driver     string              `yaml:"driver,omitempty"`
	DriverOpts map[string]string   `yaml:"driver_opts,omitempty"`
	EnableIPv6 bool                `yaml:"enable_ipv6,omitempty"`
	IPAM       V23NetworkIPAM      `yaml:"ipam,omitempty"`
	Internal   bool                `yaml:"internal,omitempty"`
	Labels     []string            `yaml:"labels,omitempty"`
	External   V23ExternalResource `yaml:"external,omitempty"`
}

type V23ExternalResource struct {
	Name string `yaml:"name,omitempty"`
}

type V23NetworkIPAM struct {
	Driver  string                 `yaml:"driver,omitempty"`
	Config  []V23NetworkIPAMConfig `yaml:"config,omitempty"`
	Options map[string]string      `yaml:"options,omitempty"`
}

type V23NetworkIPAMConfig struct {
	Subnet       string            `yaml:"subnet,omitempty"`
	IPRange      string            `yaml:"ip_range,omitempty"`
	Gateway      string            `yaml:"gateway,omitempty"`
	AuxAddresses map[string]string `yaml:"aux_addresses,omitempty"`
}

type V23Service struct {
	Build           V23ServiceBuild       `yaml:"build,omitempty"`
	CapAdd          []string              `yaml:"cap_add,omitempty"`
	CapDrop         []string              `yaml:"cap_drop,omitempty"`
	Command         string                `yaml:"command,omitempty"`
	CGroupParent    string                `yaml:"cgroup_parent,omitempty"`
	ContainerName   string                `yaml:"container_name,omitempty"`
	Devices         []string              `yaml:"devices,omitempty"`
	DependsOn       []string              `yaml:"depends_on,omitempty"`
	DNS             []string              `yaml:"dns,omitempty"`
	DNSOpt          []string              `yaml:"dns_opt,omitempty"`
	DNSSearch       []string              `yaml:"dns_search,omitempty"`
	TmpFS           []string              `yaml:"tmpfs,omitempty"`
	Entrypoint      string                `yaml:"entrypoint,omitempty"`
	EnvFile         []string              `yaml:"env_file,omitempty"`
	Environment     []string              `yaml:"environment,omitempty"`
	Expose          []string              `yaml:"expose,omitempty"`
	Extends         V1Extends             `yaml:"extends,omitempty"`
	ExternalLinks   []string              `yaml:"external_links,omitempty"`
	ExtraHosts      []string              `yaml:"extra_hosts,omitempty"`
	GroupAdd        []string              `yaml:"group_add,omitempty"`
	Healthcheck     V23ServiceHealthcheck `yaml:"healthcheck,omitempty"`
	Image           string                `yaml:"image,omitempty"`
	Init            string                `yaml:"init,omitempty"`
	Isolation       string                `yaml:"isolation,omitempty"`
	Labels          []string              `yaml:"labels,omitempty"`
	Links           []string              `yaml:"links,omitempty"`
	Logging         V32ServiceLogging     `yaml:"logging,omitempty"`
	NetworkMode     string                `yaml:"network_mode,omitempty"`
	Networks        []string              `yaml:"networks,omitempty"`
	Pid             string                `yaml:"pid,omitempty"`
	PidsLimit       int                   `yaml:"pids_limit,omitempty"`
	Ports           []string              `yaml:"ports,omitempty"`
	Scale           int                   `yaml:"scale,omitempty"`
	SecurityOpt     []string              `yaml:"security_opt,omitempty"`
	StopGracePeriod string                `yaml:"stop_grace_period,omitempty"`
	StopSignal      string                `yaml:"stop_signal,omitempty"`
	StorageOpt      string                `yaml:"storage_opt,omitempty"`
	SysCtls         []string              `yaml:"sysctls,omitempty"`
	ULimits         V1Ulimits             `yaml:"ulimits,omitempty"`
	UsernsMode      string                `yaml:"userns_mode,omitempty"`
	Volumes         []string              `yaml:"volumes,omitempty"`
	VolumeDriver    string                `yaml:"volume_driver,omitempty"`
	VolumesFrom     []string              `yaml:"volumes_from,omitempty"`

	Restart string `yaml:"restart,omitempty"`

	CpuCount   int      `yaml:"cpu_count,omitempty"`
	CpuPercent int      `yaml:"cpu_percent,omitempty"`
	Cpus       string   `yaml:"cpus,omitempty"`
	CpuShares  int      `yaml:"cpu_shares,omitempty"`
	CpuQuota   CpuQuota `yaml:"cpu_quota,omitempty"`
	CpuSet     string   `yaml:"cpuset,omitempty"`

	User           string `yaml:"user,omitempty"`
	WorkingDir     string `yaml:"working_dir,omitempty"`
	DomainName     string `yaml:"domainname,omitempty"`
	Hostname       string `yaml:"hostname,omitempty"`
	IPC            string `yaml:"ipc,omitempty"`
	MacAddress     string `yaml:"mac_address,omitempty"`
	MemLimit       string `yaml:"mem_limit,omitempty"`
	MemSwapLimit   string `yaml:"memswap_limit,omitempty"`
	MemReservation string `yaml:"mem_reservation,omitempty"`

	Privileged     bool `yaml:"privileged,omitempty"`
	OOMScoreAdj    int  `yaml:"oom_score_adj,omitempty"`
	OOMKillDisable bool `yaml:"oom_kill_disable,omitempty"`

	ReadOnly  bool   `yaml:"read_only,omitempty"`
	ShmSize   string `yaml:"shm_size,omitempty"`
	StdinOpen bool   `yaml:"stdin_open,omitempty"`
	TTY       bool   `yaml:"tty,omitempty"`
}

type V23ServiceHealthcheck struct {
	Test        string `yaml:"test,omitempty"`
	Interval    string `yaml:"interval,omitempty"`
	Timeout     string `yaml:"timeout,omitempty"`
	Retries     int    `yaml:"retries,omitempty"`
	StartPeriod string `yaml:"start_period,omitempty"`
	Disable     bool   `yaml:"disable,omitempty"`
}

type V23ServiceBuild struct {
	Context    string            `yaml:"context,omitempty"`
	Dockerfile string            `yaml:"dockerfile,omitempty"`
	Args       map[string]string `yaml:"args,omitempty"`
	ExtraHosts []string          `yaml:"extra_hosts,omitempty"`
	Labels     []string          `yaml:"labels,omitempty"`
	Network    string            `yaml:"network,omitempty"`
	ShmSize    string            `yaml:"shm_size,omitempty"`
	Target     string            `yaml:"target,omitempty"`
}
