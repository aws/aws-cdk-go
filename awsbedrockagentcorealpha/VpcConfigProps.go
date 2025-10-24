package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// VPC configuration properties.
//
// Only used when network mode is VPC.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("testVPC"))
//
//   codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
//   	CodeInterpreterCustomName: jsii.String("my_sandbox_interpreter"),
//   	Description: jsii.String("Code interpreter with isolated network access"),
//   	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
//   		Vpc: vpc,
//   	}),
//   })
//
//   codeInterpreter.connections.AddSecurityGroup(ec2.NewSecurityGroup(this, jsii.String("AdditionalGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   }))
//
// Experimental.
type VpcConfigProps struct {
	// The VPC to deploy the resource to.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to allow the resource to send all network traffic (except ipv6).
	//
	// If set to false, you must individually add traffic rules to allow the
	// resource to connect to network targets.
	//
	// Do not specify this property if the `securityGroups` property is set.
	// Instead, configure `allowAllOutbound` directly on the security group.
	// Default: true.
	//
	// Experimental.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// The list of security groups to associate with the resource's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Default: - If the resource is placed within a VPC and a security group is
	// not specified by this prop, a dedicated security
	// group will be created for this resource.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	//
	// This requires `vpc` to be specified in order for interfaces to actually be
	// placed in the subnets. If `vpc` is not specify, this will raise an error.
	// Default: - the Vpc default strategy if not specified.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

