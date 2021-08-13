package vpc

type Vpc []struct {
	Name         string `yaml:"Name"`
	Cidr_block   string `yaml:"Cidr_block"`
	DnsHostnames bool   `yaml:"DnsHostnames"`
}
