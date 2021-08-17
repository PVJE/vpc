package vpc

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Tags struct {
		Project     string `yaml:"Project"`
		Project_env string `yaml:"Project_env"`
	}

	Vpc []struct {
		Name         string `yaml:"Name"`
		Cidr_block   string `yaml:"Cidr_block"`
		DnsHostnames bool   `yaml:"DnsHostnames"`
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config, file_name_read string) {
	f, err := os.Open(file_name_read)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func MyVPC(file_name string) {
	pulumi.Run(func(ctx *pulumi.Context) error {
		var cfg Config
		readFile(&cfg, file_name)
		// var ctx *pulumi.Context
		for i := range cfg.Vpc {
			ec2.NewVpc(ctx, string(cfg.Vpc[i].Name), &ec2.VpcArgs{
				AssignGeneratedIpv6CidrBlock: pulumi.Bool(false),
				CidrBlock:                    pulumi.String(string(cfg.Vpc[0].Cidr_block)),
				//CidrBlock:          pulumi.String(string("10.9.48.64/27")),
				EnableDnsSupport:   pulumi.Bool(true),
				EnableDnsHostnames: pulumi.Bool(bool(cfg.Vpc[i].DnsHostnames)),
				InstanceTenancy:    pulumi.String("default"),
				Tags: pulumi.StringMap{
					"Name":        pulumi.String(string(cfg.Vpc[i].Name)),
					"Project":     pulumi.String(string(cfg.Tags.Project)),
					"Project-env": pulumi.String(string(cfg.Tags.Project_env)),
				},
			}, pulumi.Protect(false))

		}
		return nil
	})

}
