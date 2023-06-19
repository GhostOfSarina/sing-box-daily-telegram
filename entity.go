package main


type RealityJson struct {
	Log struct {
		Level     string `json:"level"`
		Timestamp bool   `json:"timestamp"`
	} `json:"log"`
	Inbounds  []Inbound `json:"inbounds"`
	Outbounds []struct {
		Type string `json:"type"`
		Tag  string `json:"tag"`
	} `json:"outbounds"`
}

type Inbound struct {
	Type                     string `json:"type"`
	Tag                      string `json:"tag"`
	Listen                   string `json:"listen"`
	ListenPort               int    `json:"listen_port"`
	Sniff                    bool   `json:"sniff"`
	SniffOverrideDestination bool   `json:"sniff_override_destination"`
	DomainStrategy           string `json:"domain_strategy"`
	Users                    []struct {
		UUID string `json:"uuid"`
		Flow string `json:"flow"`
	} `json:"users"`
	TLS struct {
		Enabled    bool   `json:"enabled"`
		ServerName string `json:"server_name"`
		Reality    struct {
			Enabled   bool `json:"enabled"`
			Handshake struct {
				Server     string `json:"server"`
				ServerPort int    `json:"server_port"`
			} `json:"handshake"`
			PrivateKey string   `json:"private_key"`
			ShortID    []string `json:"short_id"`
		} `json:"reality"`
	} `json:"tls"`
}
