package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocalGatewayVirtualInterfaceGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocalGatewayVirtualInterfaceGroupProps := &CfnLocalGatewayVirtualInterfaceGroupProps{
//   	LocalGatewayId: jsii.String("localGatewayId"),
//
//   	// the properties below are optional
//   	LocalBgpAsn: jsii.Number(123),
//   	LocalBgpAsnExtended: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html
//
type CfnLocalGatewayVirtualInterfaceGroupProps struct {
	// The ID of the local gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-localgatewayid
	//
	LocalGatewayId *string `field:"required" json:"localGatewayId" yaml:"localGatewayId"`
	// The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-localbgpasn
	//
	LocalBgpAsn *float64 `field:"optional" json:"localBgpAsn" yaml:"localBgpAsn"`
	// The extended 32-bit ASN for the local BGP configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-localbgpasnextended
	//
	LocalBgpAsnExtended *float64 `field:"optional" json:"localBgpAsnExtended" yaml:"localBgpAsnExtended"`
	// The tags assigned to the virtual interface group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

