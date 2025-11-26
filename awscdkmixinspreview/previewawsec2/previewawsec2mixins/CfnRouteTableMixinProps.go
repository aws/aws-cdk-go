package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRouteTablePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteTableMixinProps := &CfnRouteTableMixinProps{
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html
//
type CfnRouteTableMixinProps struct {
	// Any tags assigned to the route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html#cfn-ec2-routetable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html#cfn-ec2-routetable-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

