package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocalGatewayRouteTableVPCAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayRouteTableVPCAssociationProps := &cfnLocalGatewayRouteTableVPCAssociationProps{
//   	localGatewayRouteTableId: jsii.String("localGatewayRouteTableId"),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocalGatewayRouteTableVPCAssociationProps struct {
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The tags assigned to the association.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

