package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTransitGatewayPeering`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayPeeringProps := &CfnTransitGatewayPeeringProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	TransitGatewayArn: jsii.String("transitGatewayArn"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayPeeringProps struct {
	// The ID of the core network.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The ARN of the transit gateway.
	TransitGatewayArn *string `field:"required" json:"transitGatewayArn" yaml:"transitGatewayArn"`
	// The list of key-value tags associated with the peering.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

