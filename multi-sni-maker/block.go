package main

func Block(newReality RealityJson) RealityJson {

	newReality.DNS.Servers = make([]Server, 2)
	newReality.DNS.Servers[0] = Server{
		Address: "tcp://1.1.1.3",
		Detour:  "dns",
	}
	newReality.DNS.Servers[1] = Server{
		Address: "tcp://208.67.222.123",
		Detour:  "dns",
	}
	newReality.DNS.Strategy = "ipv4_only"

	newReality.Route.Rules = make([]Rule, 4)

	newReality.Route.Rules[0] = Rule{
		Geoip:    []string{"cn", "ir", "private"},
		Outbound: "block",
	}
	newReality.Route.Rules[1] = Rule{
		Geosite:  []string{"category-porn", "category-ads-all"},
		Outbound: "block",
	}
	newReality.Route.Rules[2] = Rule{
		IPCidr: []string{"0.0.0.0/8",
			"10.0.0.0/8",
			"100.64.0.0/10",
			"127.0.0.0/8",
			"169.254.0.0/16",
			"172.16.0.0/12",
			"192.0.0.0/24",
			"192.0.2.0/24",
			"192.168.0.0/16",
			"198.18.0.0/15",
			"198.51.100.0/24",
			"203.0.113.0/24",
			"::1/128",
			"fc00::/7",
			"fe80::/10"},
		Outbound: "block",
	}
	newReality.Route.Rules[3] = Rule{
		Network:  "tcp",
		Port:     []int{25, 587, 465, 2525},
		Outbound: "block",
	}

	return newReality

}
