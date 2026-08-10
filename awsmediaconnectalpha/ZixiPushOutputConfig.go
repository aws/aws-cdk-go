package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration options for Zixi Push outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryptionAlgorithm EncryptionAlgorithm
//   var role Role
//   var secret Secret
//
//   zixiPushOutputConfig := &ZixiPushOutputConfig{
//   	Destination: jsii.String("destination"),
//   	MaxLatency: cdk.Duration_Minutes(jsii.Number(30)),
//   	Port: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//
//   	// the properties below are optional
//   	CidrAllowList: []*string{
//   		jsii.String("cidrAllowList"),
//   	},
//   	Encryption: &StaticKeyEncryption{
//   		Algorithm: encryptionAlgorithm,
//   		Secret: secret,
//
//   		// the properties below are optional
//   		Role: role,
//   	},
//   	VpcInterfaceAttachmentName: jsii.String("vpcInterfaceAttachmentName"),
//   }
//
// Experimental.
type ZixiPushOutputConfig struct {
	// The IP address where you want to send the output.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The maximum latency for Zixi-based streams.
	// Experimental.
	MaxLatency awscdk.Duration `field:"required" json:"maxLatency" yaml:"maxLatency"`
	// The port to use when content is distributed to this output.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The stream ID that you want to use for this transport.
	// Experimental.
	StreamId *string `field:"required" json:"streamId" yaml:"streamId"`
	// The range of IP addresses that should be allowed to initiate output requests to this flow.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	// Default: - no CIDR allow list.
	//
	// Experimental.
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// Static key encryption for this output.
	// Default: - no encryption.
	//
	// Experimental.
	Encryption *StaticKeyEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// The name of the VPC interface attachment to use for this output.
	// Default: - no VPC interface attachment.
	//
	// Experimental.
	VpcInterfaceAttachmentName *string `field:"optional" json:"vpcInterfaceAttachmentName" yaml:"vpcInterfaceAttachmentName"`
}

