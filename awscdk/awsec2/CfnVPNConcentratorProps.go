package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVPNConcentrator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNConcentratorProps := &CfnVPNConcentratorProps{
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html
//
type CfnVPNConcentratorProps struct {
	// The ID of the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html#cfn-ec2-vpnconcentrator-transitgatewayid
	//
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The type of VPN concentrator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html#cfn-ec2-vpnconcentrator-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Any tags assigned to the VPN concentrator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconcentrator.html#cfn-ec2-vpnconcentrator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

