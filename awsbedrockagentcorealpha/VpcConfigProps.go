package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// VPC configuration properties.
//
// Only used when network mode is VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup SecurityGroup
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var vpc Vpc
//
//   vpcConfigProps := &VpcConfigProps{
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	AllowAllOutbound: jsii.Boolean(false),
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []SubnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []ISubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type VpcConfigProps struct {
	// The VPC to deploy the resource to.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
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
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// The list of security groups to associate with the resource's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Default: - If the resource is placed within a VPC and a security group is
	// not specified by this prop, a dedicated security
	// group will be created for this resource.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	//
	// This requires `vpc` to be specified in order for interfaces to actually be
	// placed in the subnets. If `vpc` is not specify, this will raise an error.
	// Default: - the Vpc default strategy if not specified.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

