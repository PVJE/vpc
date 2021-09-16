package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateVPC(vpc_name string, vpc_assign_generated_ipv6_cidr_block bool, vpc_cidr string, vpc_enable_dns_support bool, vpc_dns_hostanme bool, vpc_instance_tenancy string, tags_key []string, tags_values []string, ctx *pulumi.Context) (*ec2.Vpc, error) {
	tagsMap := make(pulumi.StringMap)
	for k := range tags_key {
		tagsMap[tags_key[k]] = pulumi.String(tags_values[k])
	}
	my_vpc, err := ec2.NewVpc(ctx, string(vpc_name), &ec2.VpcArgs{
		AssignGeneratedIpv6CidrBlock: pulumi.Bool(bool(vpc_assign_generated_ipv6_cidr_block)),
		CidrBlock:                    pulumi.String(string(vpc_cidr)),
		EnableDnsSupport:             pulumi.Bool(bool(vpc_enable_dns_support)),
		EnableDnsHostnames:           pulumi.Bool(bool(vpc_dns_hostanme)),
		InstanceTenancy:              pulumi.String(string(vpc_instance_tenancy)),
		Tags:                         pulumi.StringMap(tagsMap),
	}, pulumi.Protect(false))
	if err != nil {
		return nil, err
	}
	return my_vpc, nil
}
