package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnTransitGatewayMulticastDomainProps := &cfnTransitGatewayMulticastDomainProps{
//   	transitGatewayId: jsii.String("transitGatewayId"),
//
//   	// the properties below are optional
//   	options: options,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayMulticastDomainProps struct {
	// The ID of the transit gateway.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The options for the transit gateway multicast domain.
	//
	// - AutoAcceptSharedAssociations (enable | disable)
	// - Igmpv2Support (enable | disable)
	// - StaticSourcesSupport (enable | disable).
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The tags for the transit gateway multicast domain.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

