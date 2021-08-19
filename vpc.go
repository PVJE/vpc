package vpc

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vpc struct {
	Name         string
	Cidr_block   string
	DnsHostnames bool
}

func CreateVPC(file_name string, tags map[string]string, ctx *pulumi.Context, vpc Vpc) {
	//pulumi.Run(func(ctx *pulumi.Context) error {
	// id := make([]string, len(Vpc))
	ec2.NewVpc(ctx, string(vpc.Name), &ec2.VpcArgs{
		AssignGeneratedIpv6CidrBlock: pulumi.Bool(false),
		CidrBlock:                    pulumi.String(string(vpc.Cidr_block)),
		//CidrBlock:          pulumi.String(string("10.9.48.64/27")),
		EnableDnsSupport:   pulumi.Bool(true),
		EnableDnsHostnames: pulumi.Bool(bool(vpc.DnsHostnames)),
		InstanceTenancy:    pulumi.String("default"),
		Tags: pulumi.StringMap{
			"Name":        pulumi.String(string(vpc.Name)),
			"Project":     pulumi.String(string(tags["Project"])),
			"Project-env": pulumi.String(string(tags["Project_env"])),
			// "Name": pulumi.String(string(Vpc[i].Name)),
			// "Project":     pulumi.String(string(Tags.Project)),
			// "Project-env": pulumi.String(string(Tags.Project_env)),
		},
	}, pulumi.Protect(false))
	// if err != nil {
	// 	return err
	// }
	//ctx.Export("VPC_id", id.ID())

	// 	return nil
	// })

	// pulumi.Run(func(ctx *pulumi.Context) {
	// 	var cfg Config
	// 	// var id [len(Vpc)]string
	// 	readFile(&cfg, file_name)
	// 	// id := make([]string, len(Vpc))
	// 	for i := range Vpc {
	// 		id, err := ec2.NewVpc(ctx, string("My_VPC_"+strconv.Itoa(i)), &ec2.VpcArgs{
	// 			AssignGeneratedIpv6CidrBlock: pulumi.Bool(false),
	// 			CidrBlock:                    pulumi.String(string(Vpc[0].Cidr_block)),
	// 			//CidrBlock:          pulumi.String(string("10.9.48.64/27")),
	// 			EnableDnsSupport:   pulumi.Bool(true),
	// 			EnableDnsHostnames: pulumi.Bool(bool(Vpc[i].DnsHostnames)),
	// 			InstanceTenancy:    pulumi.String("default"),
	// 			Tags: pulumi.StringMap{
	// 				"Name":        pulumi.String(string(Vpc[i].Name)),
	// 				"Project":     pulumi.String(string(Tags.Project)),
	// 				"Project-env": pulumi.String(string(Tags.Project_env)),
	// 			},
	// 		}, pulumi.Protect(false))
	// 		if err != nil {
	// 			return err
	// 		}

	// 	}
	// 	return nil, id
	// })

}
