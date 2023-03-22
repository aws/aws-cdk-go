package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVpcAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcAttachmentProps := &CfnVpcAttachmentProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	SubnetArns: []*string{
//   		jsii.String("subnetArns"),
//   	},
//   	VpcArn: jsii.String("vpcArn"),
//
//   	// the properties below are optional
//   	Options: &VpcOptionsProperty{
//   		ApplianceModeSupport: jsii.Boolean(false),
//   		Ipv6Support: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVpcAttachmentProps struct {
	// The core network ID.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The subnet ARNs.
	SubnetArns *[]*string `field:"required" json:"subnetArns" yaml:"subnetArns"`
	// The ARN of the VPC attachment.
	VpcArn *string `field:"required" json:"vpcArn" yaml:"vpcArn"`
	// Options for creating the VPC attachment.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The tags associated with the VPC attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

