package main

func (c composev1) tov23() (Output, error) {
	var out composev23

	c.Unmarshal()

	out.Version = "2.3"

	out.Services = make(map[string]V23Service)
	out.Volumes = make(map[string]V32Volume)

	for k, v := range c.Document {
		out.Services[k] = V23Service{
			Build: V23ServiceBuild{
				Context:    v.Build,
				Dockerfile: v.Dockerfile,
			},
			CapAdd:        v.CapAdd,
			CapDrop:       v.CapDrop,
			Command:       v.Command,
			CGroupParent:  v.CGroupParent,
			ContainerName: v.ContainerName,
			Devices:       v.Devices,
			DNS:           v.DNS,
			DNSSearch:     v.DNSSearch,
			Entrypoint:    v.Entrypoint,
			EnvFile:       StringArray(v.EnvFile),
			Environment:   v.Environment,
			Expose:        v.Expose,
			Extends:       v.Extends,
			ExternalLinks: v.ExternalLinks,
			ExtraHosts:    v.ExtraHosts,
			Image:         v.Image,
			Labels:        v.Labels,
			Links:         v.Links,
			Logging: V32ServiceLogging{
				Driver:  v.LogDriver,
				Options: v.LogOpt,
			},
			NetworkMode:  v.Net,
			Pid:          v.Pid,
			Ports:        v.Ports,
			SecurityOpt:  v.SecurityOpt,
			StopSignal:   v.StopSignal,
			ULimits:      v.ULimits,
			Volumes:      parseVolumesToShortFormat(v.Volumes, &out),
			VolumeDriver: v.VolumeDriver,
			VolumesFrom:  v.VolumesFrom,

			CpuShares: v.CpuShares,
			CpuSet:    v.CpuSet,
			Cpus:      v.CpuQuota.AsCpus(),

			User:         v.User,
			WorkingDir:   v.WorkingDir,
			DomainName:   v.DomainName,
			Hostname:     v.Hostname,
			IPC:          v.IPC,
			MacAddress:   v.MacAddress,
			MemLimit:     v.MemLimit,
			MemSwapLimit: v.MemSwapLimit,

			Privileged: v.Privileged,
			Restart:    v.Restart,
			ReadOnly:   v.ReadOnly,
			ShmSize:    v.ShmSize,
			StdinOpen:  v.StdinOpen,
			TTY:        v.TTY,
		}
	}

	output := c.Marshal(&out)

	return output, nil
}

func parseVolumesToShortFormat(in []string, out *composev23) []string {
	for _, vol := range in {
		named := regexpNamedVolume()

		if named.MatchString(vol) {
			result := extractRegexpVars(named, &vol)

			out.Volumes[result["volume_name"]] = V32Volume{}
		}
	}

	// return same array of volumes
	return in
}
