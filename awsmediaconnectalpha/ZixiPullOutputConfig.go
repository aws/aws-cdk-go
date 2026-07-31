package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration options for Zixi Pull outputs.
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
//   var encryptionAlgorithm EncryptionAlgorithm
//   var networkInterface NetworkInterface
//   var role Role
//   var secret Secret
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   zixiPullOutputConfig := &ZixiPullOutputConfig{
//   	CidrAllowList: []*string{
//   		jsii.String("cidrAllowList"),
//   	},
//   	MaxLatency: cdk.Duration_Minutes(jsii.Number(30)),
//   	RemoteId: jsii.String("remoteId"),
//   	StreamId: jsii.String("streamId"),
//
//   	// the properties below are optional
//   	Encryption: &StaticKeyEncryption{
//   		Algorithm: encryptionAlgorithm,
//   		Secret: secret,
//
//   		// the properties below are optional
//   		Role: role,
//   	},
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
type ZixiPullOutputConfig struct {
	// The range of IP addresses that should be allowed to initiate output requests to this flow.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	//
	// Required by the MediaConnect service for Zixi Pull outputs.
	// Experimental.
	CidrAllowList *[]*string `field:"required" json:"cidrAllowList" yaml:"cidrAllowList"`
	// The maximum latency for Zixi-based streams.
	// Experimental.
	MaxLatency awscdk.Duration `field:"required" json:"maxLatency" yaml:"maxLatency"`
	// The remote ID for the Zixi-pull stream.
	// Experimental.
	RemoteId *string `field:"required" json:"remoteId" yaml:"remoteId"`
	// The stream ID that you want to use for this transport.
	// Experimental.
	StreamId *string `field:"required" json:"streamId" yaml:"streamId"`
	// Static key encryption for this output.
	// Default: - no encryption.
	//
	// Experimental.
	Encryption *StaticKeyEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// The VPC interface attachment to use for this output.
	// Default: - no VPC interface attachment.
	//
	// Experimental.
	VpcInterfaceAttachment *VpcInterfaceConfig `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

