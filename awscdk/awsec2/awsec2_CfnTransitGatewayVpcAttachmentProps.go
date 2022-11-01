package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTransitGatewayVpcAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   cfnTransitGatewayVpcAttachmentProps := &cfnTransitGatewayVpcAttachmentProps{
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	transitGatewayId: jsii.String("transitGatewayId"),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	addSubnetIds: []*string{
//   		jsii.String("addSubnetIds"),
//   	},
//   	options: options,
//   	removeSubnetIds: []*string{
//   		jsii.String("removeSubnetIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayVpcAttachmentProps struct {
	// The IDs of the subnets.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the transit gateway.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The IDs of one or more subnets to add.
	//
	// You can specify at most one subnet per Availability Zone.
	AddSubnetIds *[]*string `field:"optional" json:"addSubnetIds" yaml:"addSubnetIds"`
	// The VPC attachment options in JSON or YAML.
	//
	// - DnsSupport (enable | disable)
	// - Ipv6Support (enable| disable)
	// - ApplianceModeSupport (enable | disable).
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The IDs of one or more subnets to remove.
	RemoveSubnetIds *[]*string `field:"optional" json:"removeSubnetIds" yaml:"removeSubnetIds"`
	// The tags for the VPC attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

