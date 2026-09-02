package awsdrs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSourceNetworkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSourceNetworkMixinProps := &CfnSourceNetworkMixinProps{
//   	OriginAccountId: jsii.String("originAccountId"),
//   	OriginRegion: jsii.String("originRegion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html
//
type CfnSourceNetworkMixinProps struct {
	// The account ID containing the VPC to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-originaccountid
	//
	OriginAccountId *string `field:"optional" json:"originAccountId" yaml:"originAccountId"`
	// The region containing the VPC to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-originregion
	//
	OriginRegion *string `field:"optional" json:"originRegion" yaml:"originRegion"`
	// A set of tags associated with the Source Network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC ID to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

