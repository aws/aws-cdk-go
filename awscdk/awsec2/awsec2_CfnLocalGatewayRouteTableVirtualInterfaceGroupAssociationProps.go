package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps := &cfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps{
//   	localGatewayRouteTableId: jsii.String("localGatewayRouteTableId"),
//   	localGatewayVirtualInterfaceGroupId: jsii.String("localGatewayVirtualInterfaceGroupId"),
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
type CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps struct {
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
	// The ID of the virtual interface group.
	LocalGatewayVirtualInterfaceGroupId *string `field:"required" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
	// The tags assigned to the association.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

