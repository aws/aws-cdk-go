package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitGatewayMulticastDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   cfnTransitGatewayMulticastDomainProps := &CfnTransitGatewayMulticastDomainProps{
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//
//   	// the properties below are optional
//   	Options: options,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomain.html
//
type CfnTransitGatewayMulticastDomainProps struct {
	// The ID of the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomain.html#cfn-ec2-transitgatewaymulticastdomain-transitgatewayid
	//
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The options for the transit gateway multicast domain.
	//
	// - AutoAcceptSharedAssociations (enable | disable)
	// - Igmpv2Support (enable | disable)
	// - StaticSourcesSupport (enable | disable).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomain.html#cfn-ec2-transitgatewaymulticastdomain-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The tags for the transit gateway multicast domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastdomain.html#cfn-ec2-transitgatewaymulticastdomain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

