package main

import (
	"fmt"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v2"
)

func v1tov32(bytes *[]byte) ([]byte, error) {
	var in composev1
	var out composev32

	err := yaml.Unmarshal(*bytes, &in)
	if err != nil {
		return nil, err
	}

	fmt.Println("# NOTE: In v3.2 the following keys are not supported and will be ignored:")
	fmt.Println("# \t- extends")
	fmt.Println("# \t- volumes_from")
	fmt.Println("")

	out.Version = "3.2"
	out.Services = make(map[string]V32Service)
	out.Volumes = make(map[string]V32Volume)

	for k, v := range in {

		out.Services[k] = V32Service{
			Build: V32ServiceBuild{
				Context:    v.Build,
				Dockerfile: v.Dockerfile,
			},
			CapAdd:        v.CapAdd,
			CapDrop:       v.CapDrop,
			Command:       v.Command,
			CGroupParent:  v.CGroupParent,
			ContainerName: v.ContainerName,
			Deploy: V32ServiceDeploy{
				Resources: V32ServiceDeployResources{
					Limits: V32ServiceDeployResourcesTable{
						Memory: v.MemLimit,
						Cpus:   v.CpuQuota.AsCpus(),
					},
				},
				RestartPolicy: V32ServiceDeployRestartPolicy{
					Condition: parseRestartPolicy(v.Restart),
				},
			},
			Devices:       v.Devices,
			DNS:           v.DNS,
			DNSSearch:     v.DNSSearch,
			Entrypoint:    v.Entrypoint,
			EnvFile:       v.EnvFile,
			Environment:   v.Environment,
			Expose:        v.Expose,
			ExternalLinks: v.ExternalLinks,
			ExtraHosts:    v.ExtraHosts,
			Image:         v.Image,
			Labels:        v.Labels,
			Links:         v.Links,
			Logging: V32ServiceLogging{
				Driver:  v.LogDriver,
				Options: v.LogOpt,
			},
			NetworkMode: v.Net,
			Pid:         v.Pid,
			Ports:       parsePortsToLongFormat(v.Ports),
			SecurityOpt: v.SecurityOpt,
			StopSignal:  v.StopSignal,
			ULimits:     v.ULimits,
			Volumes:     parseVolumesToLongFormat(v.Volumes, out.Volumes),

			Restart:    v.Restart,
			User:       v.User,
			WorkingDir: v.WorkingDir,
			DomainName: v.DomainName,
			Hostname:   v.Hostname,
			IPC:        v.IPC,
			MacAddress: v.MacAddress,
			Privileged: v.Privileged,
			ReadOnly:   v.ReadOnly,
			ShmSize:    v.ShmSize,
			StdinOpen:  v.StdinOpen,
			TTY:        v.TTY,
		}
	}

	bytesout, err := yaml.Marshal(&out)
	if err != nil {
		return nil, err
	}

	return bytesout, nil
}

func extractRegexpVars(reg *regexp.Regexp, str *string) map[string]string {
	match := reg.FindStringSubmatch(*str)
	result := make(map[string]string)
	for i, name := range reg.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	return result
}

func getProto(proto string) string {
	if proto == "" {
		return "tcp"
	} else {
		return proto
	}
}

type PortRange struct {
	Start int
	End   int
}

func (r *PortRange) validate(msg string) {
	if r.End < r.Start {
		panic(fmt.Sprintf("%s port range is out of bounds, start is %d, end is %d", msg, r.Start, r.End))
	}
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func parseRestartPolicy(in string) string {
	switch in {
	case "always", "unless-stopped": // "always" means "any" in deploy, which is default, do not return anything
		// "unless-stopped" is not handled by "service", defaulting to nil
		return ""
	case "no":
		return "none"
	}

	return in
}

func parseVolumesToLongFormat(in []string, volumes map[string]V32Volume) []V32ServiceVolumes {
	var out []V32ServiceVolumes

	for _, vol := range in {
		named := regexpNamedVolume()
		path := regexpPathVolume()

		if named.MatchString(vol) {
			result := extractRegexpVars(named, &vol)

			out = append(out, V32ServiceVolumes{
				Type:   "volume",
				Source: result["volume_name"],
				Target: result["container_path"],
			})

			volumes[result["volume_name"]] = V32Volume{
				Driver: "local",
				External: V32ExternalResource{
					Name: result["volume_name"],
				},
			}

		} else if path.MatchString(vol) {
			result := extractRegexpVars(path, &vol)

			out = append(out, V32ServiceVolumes{
				Type:   "bind",
				Source: result["host_path"],
				Target: result["container_path"],
			})
		}
	}

	return out
}

func parsePortsToLongFormat(in []string) []V32ServicePorts {
	var out []V32ServicePorts

	for _, port := range in {
		full := regexpIpSourceRangeTargetRange()
		ranges := regexpSourceRangeTargetRange()
		ipsourcetarget := regexpIpSourceTarget()
		sourcetarget := regexpSourceTarget()
		rrange := regexpRange()
		justport := regexpJustPort()

		if full.MatchString(port) {
			result := extractRegexpVars(full, &port)
			proto := getProto(result["proto"])

			publishRange := PortRange{
				Start: Atoi(result["ss"]),
				End:   Atoi(result["se"]),
			}

			targetRange := PortRange{
				Start: Atoi(result["ts"]),
				End:   Atoi(result["te"]),
			}

			publishRange.validate("Published")
			targetRange.validate("Target")

			if publishRange.End-publishRange.Start != targetRange.End-targetRange.Start {
				panic(fmt.Sprintf("Number of published/target ports in range is different, matching port is %s", port))
			}

			counter := 0
			for i := publishRange.Start; i <= publishRange.End; i++ {
				out = append(out, V32ServicePorts{
					Target:    targetRange.Start + counter,
					Published: result["ip"] + ":" + strconv.Itoa(i),
					Protocol:  proto,
					Mode:      "host",
				})

				counter = counter + 1
			}
		} else if ranges.MatchString(port) {
			result := extractRegexpVars(ranges, &port)
			proto := getProto(result["proto"])

			publishRange := PortRange{
				Start: Atoi(result["ss"]),
				End:   Atoi(result["se"]),
			}

			targetRange := PortRange{
				Start: Atoi(result["ts"]),
				End:   Atoi(result["te"]),
			}

			publishRange.validate("Published")
			targetRange.validate("Target")

			if publishRange.End-publishRange.Start != targetRange.End-targetRange.Start {
				panic(fmt.Sprintf("Number of published/target ports in range is different, matching port is %s", port))
			}

			counter := 0
			for i := publishRange.Start; i <= publishRange.End; i++ {
				out = append(out, V32ServicePorts{
					Target:    targetRange.Start + counter,
					Published: strconv.Itoa(i),
					Protocol:  proto,
					Mode:      "host",
				})

				counter = counter + 1
			}
		} else if ipsourcetarget.MatchString(port) {
			result := extractRegexpVars(ipsourcetarget, &port)
			proto := getProto(result["proto"])

			target, _ := strconv.Atoi(result["target"])

			out = append(out, V32ServicePorts{
				Target:    target,
				Published: result["ip"] + ":" + result["publish"],
				Protocol:  proto,
				Mode:      "host",
			})
		} else if sourcetarget.MatchString(port) {
			result := extractRegexpVars(sourcetarget, &port)
			proto := getProto(result["proto"])

			target, _ := strconv.Atoi(result["target"])

			out = append(out, V32ServicePorts{
				Target:    target,
				Published: result["publish"],
				Protocol:  proto,
				Mode:      "host",
			})
		} else if rrange.MatchString(port) {
			result := extractRegexpVars(rrange, &port)
			proto := getProto(result["proto"])

			targetRange := PortRange{
				Start: Atoi(result["rangestart"]),
				End:   Atoi(result["rangeend"]),
			}
			targetRange.validate("Target")

			for i := targetRange.Start; i <= targetRange.End; i++ {
				out = append(out, V32ServicePorts{
					Target:   i,
					Protocol: proto,
					Mode:     "host",
				})
			}
		} else if justport.MatchString(port) {
			result := extractRegexpVars(justport, &port)
			proto := getProto(result["proto"])

			p, _ := strconv.Atoi(result["port"])

			out = append(out, V32ServicePorts{
				Target:   p,
				Mode:     "host",
				Protocol: proto,
			})
		}
	}

	return out
}

func createRegexp(match string) *regexp.Regexp {
	reg, err := regexp.Compile(match)
	if err != nil {
		panic(err)
	}

	return reg
}

func regexpIpSourceRangeTargetRange() *regexp.Regexp {
	return createRegexp(`^(?P<ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}):(?P<ss>\d+)-(?P<se>\d+):(?P<ts>\d+)-(?P<te>\d+)(\/(?P<proto>[a-z]+))?$`)
}

func regexpSourceRangeTargetRange() *regexp.Regexp {
	return createRegexp(`^(?P<ss>\d+)-(?P<se>\d+):(?P<ts>\d+)-(?P<te>\d+)(\/(?P<proto>[a-z]+))?$`)
}

func regexpIpSourceTarget() *regexp.Regexp {
	return createRegexp(`^(?P<ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}):(?P<publish>\d+):(?P<target>\d+)(\/(?P<proto>[a-z]+))?$`)
}

func regexpSourceTarget() *regexp.Regexp {
	return createRegexp(`^(?P<publish>\d+):(?P<target>\d+)(\/(?P<proto>[a-z]+))?$`)
}

func regexpRange() *regexp.Regexp {
	return createRegexp(`^(?P<rangestart>\d+)-(?P<rangeend>\d+)(\/(?P<proto>[a-z]+))?$`)
}

func regexpJustPort() *regexp.Regexp {
	return createRegexp(`^(?P<port>\d+)(\/(?P<proto>[a-z]+))?$`)
}

func regexpNamedVolume() *regexp.Regexp {
	return createRegexp(`^(?P<volume_name>[a-z0-9-_]+):(?P<container_path>.+)$`)
}

func regexpPathVolume() *regexp.Regexp {
	return createRegexp(`^(?P<host_path>[.\/](.+)):(?P<container_path>[.\/](.+))$`)
}
