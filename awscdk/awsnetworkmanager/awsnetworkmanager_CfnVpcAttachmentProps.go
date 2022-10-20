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
//   cfnVpcAttachmentProps := &cfnVpcAttachmentProps{
//   	coreNetworkId: jsii.String("coreNetworkId"),
//   	options: &vpcOptionsProperty{
//   		ipv6Support: jsii.Boolean(false),
//   	},
//   	subnetArns: []*string{
//   		jsii.String("subnetArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcArn: jsii.String("vpcArn"),
//   }
//
type CfnVpcAttachmentProps struct {
	// The core network ID.
	CoreNetworkId *string `field:"optional" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Options for creating the VPC attachment.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The subnet ARNs.
	SubnetArns *[]*string `field:"optional" json:"subnetArns" yaml:"subnetArns"`
	// The tags associated with the VPC attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the VPC attachment.
	VpcArn *string `field:"optional" json:"vpcArn" yaml:"vpcArn"`
}

