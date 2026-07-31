package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration options for SRT Listener outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkInterface NetworkInterface
//   var role Role
//   var secret Secret
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   srtListenerOutputConfig := &SrtListenerOutputConfig{
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	CidrAllowList: []*string{
//   		jsii.String("cidrAllowList"),
//   	},
//   	Encryption: &SrtPasswordEncryption{
//   		Secret: secret,
//
//   		// the properties below are optional
//   		Role: role,
//   	},
//   	MinLatency: cdk.Duration_Minutes(jsii.Number(30)),
//   	VpcInterfaceAttachment: &VpcInterfaceConfig{
//   		Name: jsii.String("name"),
//   		Role: role,
//   		SecurityGroups: []ISecurityGroup{
//   			securityGroup,
//   		},
//   		Subnet: subnet,
//
//   		// the properties below are optional
//   		NetworkInterfaceIds: []*string{
//   			jsii.String("networkInterfaceIds"),
//   		},
//   		NetworkInterfaceType: networkInterface,
//   	},
//   }
//
// Experimental.
type SrtListenerOutputConfig struct {
	// The port to use when content is distributed to this output.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The range of IP addresses that should be allowed to initiate output requests to this flow.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	// Default: - no CIDR allow list.
	//
	// Experimental.
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// SRT password encryption for this output.
	// Default: - no encryption.
	//
	// Experimental.
	Encryption *SrtPasswordEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// The value you set on your
	// MediaConnect output represents the minimal potential latency of that connection. The
	// latency of the stream is set to the higher of the sender's minimum latency and the
	// receiver's minimum latency.
	// Default: - no minimum latency.
	//
	// Experimental.
	MinLatency awscdk.Duration `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The VPC interface attachment to use for this output.
	// Default: - no VPC interface attachment.
	//
	// Experimental.
	VpcInterfaceAttachment *VpcInterfaceConfig `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

