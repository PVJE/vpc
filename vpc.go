package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateVPC(vpc_name string, vpc_cidr string, vpc_dnshostanme bool, tags_key []string, tags_values []string, ctx *pulumi.Context) (*ec2.Vpc, error) {
	var my_vpc *ec2.Vpc
	var err error
	for i := range tags_key {
		my_vpc, err = ec2.NewVpc(ctx, string(vpc_name), &ec2.VpcArgs{
			AssignGeneratedIpv6CidrBlock: pulumi.Bool(false),
			CidrBlock:                    pulumi.String(string(vpc_cidr)),
			EnableDnsSupport:             pulumi.Bool(true),
			EnableDnsHostnames:           pulumi.Bool(bool(vpc_dnshostanme)),
			InstanceTenancy:              pulumi.String("default"),
			Tags: pulumi.StringMap{
				tags_key[i]: pulumi.String(string(tags_values[i])),
				"pulumi":    pulumi.String(string(vpc_name)),
				// "Project":     pulumi.String(string(tags["Project"])),
				// "Project-env": pulumi.String(string(tags["Project_env"])),
			},
		}, pulumi.Protect(false))
		if err != nil {
			return nil, err
		}
	}
	return my_vpc, nil
}
