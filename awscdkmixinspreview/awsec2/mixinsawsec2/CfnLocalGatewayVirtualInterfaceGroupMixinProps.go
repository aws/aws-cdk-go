package mixinsawsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLocalGatewayVirtualInterfaceGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocalGatewayVirtualInterfaceGroupMixinProps := &CfnLocalGatewayVirtualInterfaceGroupMixinProps{
//   	LocalBgpAsn: jsii.Number(123),
//   	LocalBgpAsnExtended: jsii.Number(123),
//   	LocalGatewayId: jsii.String("localGatewayId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html
//
type CfnLocalGatewayVirtualInterfaceGroupMixinProps struct {
	// The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-localbgpasn
	//
	LocalBgpAsn *float64 `field:"optional" json:"localBgpAsn" yaml:"localBgpAsn"`
	// The extended 32-bit ASN for the local BGP configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-localbgpasnextended
	//
	LocalBgpAsnExtended *float64 `field:"optional" json:"localBgpAsnExtended" yaml:"localBgpAsnExtended"`
	// The ID of the local gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-localgatewayid
	//
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// The tags assigned to the virtual interface group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayvirtualinterfacegroup.html#cfn-ec2-localgatewayvirtualinterfacegroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

