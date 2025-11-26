package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLocalGatewayRouteTablePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocalGatewayRouteTableMixinProps := &CfnLocalGatewayRouteTableMixinProps{
//   	LocalGatewayId: jsii.String("localGatewayId"),
//   	Mode: jsii.String("mode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetable.html
//
type CfnLocalGatewayRouteTableMixinProps struct {
	// The ID of the local gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetable.html#cfn-ec2-localgatewayroutetable-localgatewayid
	//
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// The mode of the local gateway route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetable.html#cfn-ec2-localgatewayroutetable-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The tags assigned to the local gateway route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetable.html#cfn-ec2-localgatewayroutetable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

