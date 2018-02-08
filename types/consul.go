package types

type ConsulTlsConfig struct {
	CAFile   string `json:"cafile"`
	CAPath   string `json:"capath"`
	CertFile string `json:"certfile"`
	KeyFile  string `json:"keyfile"`
}

type ConsulConfig struct {
	Addr       string          `json:"addr"`
	Datacenter string          `json:"dc"`
	Token      string          `json:"token"`
	Scheme     string          `json:"scheme"`
	Tls        ConsulTlsConfig `json:"tls"`
}
