package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("v1tov32", func() {
	Describe("parseRestartPolicy", func() {

		tests := map[string]string{
			"always":         "",
			"unless-stopped": "",
			"no":             "none",
			"other":          "other",
		}

		It("returns proper values", func() {
			for t, e := range tests {
				Expect(parseRestartPolicy(t)).To(Equal(e))
			}
		})
	})

	Describe("getProto", func() {
		tests := map[string]string{
			"":    "tcp",
			"udp": "udp",
		}

		It("returns proper values", func() {
			for t, e := range tests {
				Expect(getProto(t)).To(Equal(e))
			}
		})
	})

	Describe("PortRange validate", func() {
		Context("invalid port range", func() {
			pr := PortRange{
				Start: 100,
				End:   50,
			}

			f := func() {
				pr.validate("test")
			}

			It("panics", func() {
				Expect(f).To(Panic())
			})
		})

		Context("valid port range", func() {
			pr := PortRange{
				Start: 100,
				End:   150,
			}

			f := func() { pr.validate("test") }

			It("does not panic", func() {
				Expect(f).NotTo(Panic())
			})
		})
	})

	Describe("extractRegexpVars", func() {
		re := createRegexp(`^(?P<volume_name>[a-z0-9-_]+):(?P<container_path>.+)$`)

		str := "test:/var/www"

		It("expects to match the test string", func() {
			Expect(re.MatchString(str)).To(BeTrue())
		})

		exp := map[string]string{
			"volume_name":    "test",
			"container_path": "/var/www",
		}

		It("expects to properly extract the vars", func() {
			Expect(extractRegexpVars(re, &str)).To(Equal(exp))
		})
	})

	Describe("parseVolumesToLongFormat", func() {
		Context("named volumes", func() {
			in := []string{"test:/var"}
			v := map[string]V32Volume{}
			out := parseVolumesToLongFormat(in, v)

			It("has proper volume defined", func() {
				e := []V32ServiceVolumes{
					V32ServiceVolumes{
						Type:   "volume",
						Source: "test",
						Target: "/var",
					},
				}

				Expect(out).To(Equal(e))
			})

			It("has a global volume defined", func() {
				Expect(v).NotTo(BeEmpty())
			})

			It("has proper global volume defined", func() {
				e := V32Volume{
					Driver: "local",
					External: V32ExternalResource{
						Name: "test",
					},
				}

				Expect(v["test"]).To(Equal(e))
			})
		})
		Context("path volumes", func() {
			in := []string{"/var/www:/app"}
			v := map[string]V32Volume{}
			out := parseVolumesToLongFormat(in, v)

			It("has proper volume defined", func() {
				e := []V32ServiceVolumes{
					V32ServiceVolumes{
						Type:   "bind",
						Source: "/var/www",
						Target: "/app",
					},
				}

				Expect(out).To(Equal(e))
			})

			It("has no global volume defined", func() {
				Expect(v).To(BeEmpty())
			})
		})
	})

	Describe("parsePortsToLongFormat", func() {
		Context("ip, source and target is range", func() {
			in := []string{"127.0.0.1:5000-5001:6000-6001"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "127.0.0.1:5000",
						Protocol:  "tcp",
						Mode:      "host",
					},
					V32ServicePorts{
						Target:    6001,
						Published: "127.0.0.1:5001",
						Protocol:  "tcp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("source and target is range", func() {
			in := []string{"5000-5001:6000-6001"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "5000",
						Protocol:  "tcp",
						Mode:      "host",
					},
					V32ServicePorts{
						Target:    6001,
						Published: "5001",
						Protocol:  "tcp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("single port with ip", func() {
			in := []string{"127.0.0.1:5000:6000"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "127.0.0.1:5000",
						Protocol:  "tcp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("single port without ip", func() {
			in := []string{"5000:6000"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "5000",
						Protocol:  "tcp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("port ranges without publish", func() {
			in := []string{"5000-5001"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    5000,
						Published: "",
						Protocol:  "tcp",
						Mode:      "host",
					},
					V32ServicePorts{
						Target:    5001,
						Published: "",
						Protocol:  "tcp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("just port", func() {
			in := []string{"5000"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    5000,
						Published: "",
						Protocol:  "tcp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})

		Context("ip, source and target is range udp", func() {
			in := []string{"127.0.0.1:5000-5001:6000-6001/udp"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "127.0.0.1:5000",
						Protocol:  "udp",
						Mode:      "host",
					},
					V32ServicePorts{
						Target:    6001,
						Published: "127.0.0.1:5001",
						Protocol:  "udp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("source and target is range udp", func() {
			in := []string{"5000-5001:6000-6001/udp"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "5000",
						Protocol:  "udp",
						Mode:      "host",
					},
					V32ServicePorts{
						Target:    6001,
						Published: "5001",
						Protocol:  "udp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("single port with ip udp", func() {
			in := []string{"127.0.0.1:5000:6000/udp"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "127.0.0.1:5000",
						Protocol:  "udp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("single port without ip udp", func() {
			in := []string{"5000:6000/udp"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    6000,
						Published: "5000",
						Protocol:  "udp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("port ranges without publish udp", func() {
			in := []string{"5000-5001/udp"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    5000,
						Published: "",
						Protocol:  "udp",
						Mode:      "host",
					},
					V32ServicePorts{
						Target:    5001,
						Published: "",
						Protocol:  "udp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
		Context("just port udp", func() {
			in := []string{"5000/udp"}
			out := parsePortsToLongFormat(in)

			It("has proper struct defined", func() {
				e := []V32ServicePorts{
					V32ServicePorts{
						Target:    5000,
						Published: "",
						Protocol:  "udp",
						Mode:      "host",
					},
				}

				Expect(out).To(Equal(e))
			})
		})
	})
})
