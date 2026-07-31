package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for RIST.
//
// Example:
//   var stack Stack
//   var securityGroup ISecurityGroup
//   var subnet ISubnet
//   var role IRole
//
//
//   vpcInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("my-vpc-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   })
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Rist(&SourceRist{
//   		FlowSourceName: jsii.String("vpc-source"),
//   		Description: jsii.String("VPC-based source"),
//   		Port: jsii.Number(5000),
//   		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
//   		Network: awsmediaconnectalpha.NetworkConfiguration_Vpc(vpcInterface),
//   	}),
//   	VpcInterfaces: []VpcInterfaceConfig{
//   		vpcInterface,
//   	},
//   })
//
// Experimental.
type SourceRist struct {
	// A description of the source.
	//
	// This description appears only on the MediaConnect console and will not be seen by the end user.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the source.
	// Default: - a name is generated automatically.
	//
	// Experimental.
	FlowSourceName *string `field:"optional" json:"flowSourceName" yaml:"flowSourceName"`
	// Defines networking configuration.
	// Experimental.
	Network NetworkConfiguration `field:"required" json:"network" yaml:"network"`
	// The port that the flow will be listening on for incoming content.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	// Default: - no maximum bitrate.
	//
	// Experimental.
	MaxBitrate awscdk.Bitrate `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum latency in milliseconds for a RIST or Zixi-based source.
	// Default: - undefined; when omitted, MediaConnect applies its service default maximum latency.
	//
	// Experimental.
	MaxLatency awscdk.Duration `field:"optional" json:"maxLatency" yaml:"maxLatency"`
}

