package main

type composev33 V33

type V33 struct {
	Version  string                `yaml:"version"`
	Services map[string]V33Service `yaml:"services"`
	Networks map[string]V33Network `yaml:"networks,omitempty"`
	Volumes  map[string]V33Volume  `yaml:"volumes,omitempty"`
}

type V33Volume struct {
	Driver     string              `yaml:"driver,omitempty"`
	DriverOpts map[string]string   `yaml:"driver_opts,omitempty"`
	External   V33ExternalResource `yaml:"external,omitempty"`
	Labels     []string            `yaml:"labels,omitempty"`
}

type V33Network struct {
	Driver     string              `yaml:"driver,omitempty"`
	DriverOpts map[string]string   `yaml:"driver_opts,omitempty"`
	EnableIPv6 bool                `yaml:"enable_ipv6,omitempty"`
	IPAM       V33NetworkIPAM      `yaml:"ipam,omitempty"`
	Internal   bool                `yaml:"internal,omitempty"`
	Labels     []string            `yaml:"labels,omitempty"`
	External   V33ExternalResource `yaml:"external,omitempty"`
}

type V33ExternalResource struct {
	Name string `yaml:"name,omitempty"`
}

type V33NetworkIPAM struct {
	Driver string                 `yaml:"driver,omitempty"`
	Config []V33NetworkIPAMConfig `yaml:"config,omitempty"`
}

type V33NetworkIPAMConfig struct {
	Subnet string `yaml:"subnet,omitempty"`
}

type V33Service struct {
	Build           V33ServiceBuild       `yaml:"build,omitempty"`
	CapAdd          []string              `yaml:"cap_add,omitempty"`
	CapDrop         []string              `yaml:"cap_drop,omitempty"`
	Command         string                `yaml:"command,omitempty"`
	CGroupParent    string                `yaml:"cgroup_parent,omitempty"`
	ContainerName   string                `yaml:"container_name,omitempty"`
	Deploy          V33ServiceDeploy      `yaml:"deploy,omitempty"`
	Devices         []string              `yaml:"devices,omitempty"`
	DependsOn       []string              `yaml:"depends_on,omitempty"`
	DNS             []string              `yaml:"dns,omitempty"`
	DNSSearch       []string              `yaml:"dns_search,omitempty"`
	TmpFS           []string              `yaml:"tmpfs,omitempty"`
	Entrypoint      string                `yaml:"entrypoint,omitempty"`
	EnvFile         string                `yaml:"env_file,omitempty"`
	Environment     []string              `yaml:"environment,omitempty"`
	Expose          []string              `yaml:"expose,omitempty"`
	ExternalLinks   []string              `yaml:"external_links,omitempty"`
	ExtraHosts      []string              `yaml:"extra_hosts,omitempty"`
	Healthcheck     V33ServiceHealthcheck `yaml:"healthcheck,omitempty"`
	Image           string                `yaml:"image,omitempty"`
	Isolation       string                `yaml:"isolation,omitempty"`
	Labels          []string              `yaml:"labels,omitempty"`
	Links           []string              `yaml:"links,omitempty"`
	Logging         V33ServiceLogging     `yaml:"logging,omitempty"`
	NetworkMode     string                `yaml:"network_mode,omitempty"`
	Networks        []string              `yaml:"networks,omitempty"`
	Pid             string                `yaml:"pid,omitempty"`
	Ports           []V33ServicePorts     `yaml:"ports,omitempty"`
	Secrets         []V33ServiceSecrets   `yaml:"secrets,omitempty"`
	SecurityOpt     []string              `yaml:"security_opt,omitempty"`
	StopGracePeriod string                `yaml:"stop_grace_period,omitempty"`
	StopSignal      string                `yaml:"stop_signal,omitempty"`
	SysCtls         []string              `yaml:"sysctls,omitempty"`
	ULimits         V1Ulimits             `yaml:"ulimits,omitempty"`
	Volumes         []V33ServiceVolumes   `yaml:"volumes,omitempty"`

	Restart    string `yaml:"restart,omitempty"`
	User       string `yaml:"user,omitempty"`
	WorkingDir string `yaml:"working_dir,omitempty"`
	DomainName string `yaml:"domainname,omitempty"`
	Hostname   string `yaml:"hostname,omitempty"`
	IPC        string `yaml:"ipc,omitempty"`
	MacAddress string `yaml:"mac_address,omitempty"`
	Privileged bool   `yaml:"privileged,omitempty"`
	ReadOnly   bool   `yaml:"read_only,omitempty"`
	ShmSize    string `yaml:"shm_size,omitempty"`
	StdinOpen  bool   `yaml:"stdin_open,omitempty"`
	TTY        bool   `yaml:"tty,omitempty"`
}

type V33ServiceVolumes struct {
	Type     string                  `yaml:"type,omitempty"`
	Source   string                  `yaml:"source,omitempty"`
	Target   string                  `yaml:"target,omitempty"`
	ReadOnly bool                    `yaml:"read_only,omitempty"`
	Bind     V33ServiceVolumesBind   `yaml:"bind,omitempty"`
	Volume   V33ServiceVolumesVolume `yaml:"volume,omitempty"`
}

type V33ServiceVolumesBind struct {
	Propagation string `yaml:"propagation,omitempty"`
}

type V33ServiceVolumesVolume struct {
	Nocopy bool `yaml:"nocopy,omitempty"`
}

type V33ServiceSecrets struct {
	Source string `yaml:"source,omitempty"`
	Target string `yaml:"target,omitempty"`
	UID    string `yaml:"uid,omitempty"`
	GID    string `yaml:"gid,omitempty"`
	Mode   int    `yaml:"mode,omitempty"`
}

type V33ServicePorts struct {
	Target    int    `yaml:"target,omitempty"`
	Published string `yaml:"published,omitempty"`
	Protocol  string `yaml:"protocol,omitempty"`
	Mode      string `yaml:"mode,omitempty"`
}

type V33ServiceLogging struct {
	Driver  string            `yaml:"driver,omitempty"`
	Options map[string]string `yaml:"options,omitempty"`
}

type V33ServiceHealthcheck struct {
	Test     string `yaml:"test,omitempty"`
	Interval string `yaml:"interval,omitempty"`
	Timeout  string `yaml:"timeout,omitempty"`
	Retries  int    `yaml:"retries,omitempty"`
	Disable  bool   `yaml:"disable,omitempty"`
}

type V33ServiceBuild struct {
	Context    string            `yaml:"context,omitempty"`
	Dockerfile string            `yaml:"dockerfile,omitempty"`
	Args       map[string]string `yaml:"args,omitempty"`
	CacheFrom  []string          `yaml:"cache_from,omitempty"`
	Labels     []string          `yaml:"labels,omitempty"`
}

type V33ServiceDeploy struct {
	Mode          string                        `yaml:"mode,omitempty"`
	Replicas      int                           `yaml:"replicas,omitempty"`
	Placement     map[string][]string           `yaml:"placement,omitempty"`
	UpdateConfig  V33ServiceDeployUpdateConfig  `yaml:"update_config,omitempty"`
	Resources     V33ServiceDeployResources     `yaml:"resources,omitempty"`
	RestartPolicy V33ServiceDeployRestartPolicy `yaml:"restart_policy,omitempty"`
	Labels        []string                      `yaml:"labels,omitempty"`
}

type V33ServiceDeployRestartPolicy struct {
	Condition   string `yaml:"condition,omitempty"`
	Delay       string `yaml:"delay,omitempty"`
	MaxAttempts int    `yaml:"max_attempts,omitempty"`
	Window      string `yaml:"window,omitempty"`
}

type V33ServiceDeployResources struct {
	Limits       V33ServiceDeployResourcesTable `yaml:"limits,omitempty"`
	Reservations V33ServiceDeployResourcesTable `yaml:"reservations,omitempty"`
}

type V33ServiceDeployResourcesTable struct {
	Cpus   string `yaml:"cpus,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}

type V33ServiceDeployUpdateConfig struct {
	Parallelism     int    `yaml:"parallelism,omitempty"`
	Delay           string `yaml:"delay,omitempty"`
	FailureAction   string `yaml:"failure_action,omitempty"`
	Monitor         string `yaml:"monitor,omitempty"`
	MaxFailureRatio int    `yaml:"max_failure_ratio,omitempty"`
}
