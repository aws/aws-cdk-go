package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps := &CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps{
//   	LocalGatewayRouteTableId: jsii.String("localGatewayRouteTableId"),
//   	LocalGatewayVirtualInterfaceGroupId: jsii.String("localGatewayVirtualInterfaceGroupId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevirtualinterfacegroupassociation.html
//
type CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps struct {
	// The ID of the local gateway route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevirtualinterfacegroupassociation.html#cfn-ec2-localgatewayroutetablevirtualinterfacegroupassociation-localgatewayroutetableid
	//
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
	// The ID of the virtual interface group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevirtualinterfacegroupassociation.html#cfn-ec2-localgatewayroutetablevirtualinterfacegroupassociation-localgatewayvirtualinterfacegroupid
	//
	LocalGatewayVirtualInterfaceGroupId *string `field:"required" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
	// The tags assigned to the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevirtualinterfacegroupassociation.html#cfn-ec2-localgatewayroutetablevirtualinterfacegroupassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

