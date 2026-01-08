package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVPNConcentratorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConcentratorMixinProps := &CfnVPNConcentratorMixinProps{
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html
//
type CfnVPNConcentratorMixinProps struct {
	// Any tags assigned to the VPN concentrator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html#cfn-ec2-vpnconcentrator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the transit gateway associated with the VPN concentrator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html#cfn-ec2-vpnconcentrator-transitgatewayid
	//
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The type of VPN concentrator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html#cfn-ec2-vpnconcentrator-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

