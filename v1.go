package main

import (
	"strconv"
)

type composev1 map[string]V1
type CpuQuota int

type V1 struct {
	Build         string            `yaml:"build,omitempty"`
	Dockerfile    string            `yaml:"dockerfile,omitempty"`
	CapAdd        []string          `yaml:"cap_add,omitempty"`
	CapDrop       []string          `yaml:"cap_drop,omitempty"`
	Command       string            `yaml:"command,omitempty"`
	CGroupParent  string            `yaml:"cgroup_parent,omitempty"`
	ContainerName string            `yaml:"container_name,omitempty"`
	Devices       []string          `yaml:"devices,omitempty"`
	DNS           []string          `yaml:"dns,omitempty"`
	DNSSearch     []string          `yaml:"dns_search,omitempty"`
	Entrypoint    string            `yaml:"entrypoint,omitempty"`
	EnvFile       string            `yaml:"env_file,omitempty"`
	Environment   []string          `yaml:"environment,omitempty"`
	Expose        []string          `yaml:"expose,omitempty"`
	Extends       V1Extends         `yaml:"extends,omitempty"`
	ExternalLinks []string          `yaml:"external_links,omitempty"`
	ExtraHosts    []string          `yaml:"extra_hosts,omitempty"`
	Image         string            `yaml:"image,omitempty"`
	Labels        []string          `yaml:"labels,omitempty"`
	Links         []string          `yaml:"links,omitempty"`
	LogDriver     string            `yaml:"log_driver,omitempty"`
	LogOpt        map[string]string `yaml:"log_opts,omitempty"`
	Net           string            `yaml:"net,omitempty"`
	Pid           string            `yaml:"pid,omitempty"`
	Ports         []string          `yaml:"ports,omitempty"`
	SecurityOpt   []string          `yaml:"security_opt,omitempty"`
	StopSignal    string            `yaml:"stop_signal,omitempty"`
	ULimits       V1Ulimits         `yaml:"ulimits,omitempty"`
	Volumes       []string          `yaml:"volumes,omitempty"`
	VolumeDriver  string            `yaml:"volume_driver,omitempty"`
	VolumesFrom   []string          `yaml:"volumes_from,omitempty"`

	CpuShares    int      `yaml:"cpu_shares,omitempty"`
	CpuQuota     CpuQuota `yaml:"cpu_quota,omitempty"`
	CpuSet       string   `yaml:"cpuset,omitempty"`
	User         string   `yaml:"user,omitempty"`
	WorkingDir   string   `yaml:"working_dir,omitempty"`
	DomainName   string   `yaml:"domainname,omitempty"`
	Hostname     string   `yaml:"hostname,omitempty"`
	IPC          string   `yaml:"ipc,omitempty"`
	MacAddress   string   `yaml:"mac_address,omitempty"`
	MemLimit     string   `yaml:"mem_limit,omitempty"`
	MemSwapLimit string   `yaml:"memswap_limit,omitempty"`
	Privileged   bool     `yaml:"privileged,omitempty"`
	Restart      string   `yaml:"restart,omitempty"`
	ReadOnly     bool     `yaml:"read_only,omitempty"`
	ShmSize      string   `yaml:"shm_size,omitempty"`
	StdinOpen    bool     `yaml:"stdin_open,omitempty"`
	TTY          bool     `yaml:"tty,omitempty"`
}

type V1Extends struct {
	File    string `yaml:"file,omitempty"`
	Service string `yaml:"service,omitempty"`
}

type V1Ulimits struct {
	NProc  int             `yaml:"nproc,omitempty"`
	Nofile V1UlimitsNofile `yaml:"nofile,omitempty"`
}

type V1UlimitsNofile struct {
	Soft int `yaml:"soft,omitempty"`
	Hard int `yaml:"hard,omitempty"`
}

func (q *CpuQuota) AsCpus() string {
	cpusf := float64(*q) / 100000
	cpuss := strconv.FormatFloat(cpusf, 'f', -1, 32)

	if cpuss == "0" {
		cpuss = "1.0"
	}

	return cpuss
}
