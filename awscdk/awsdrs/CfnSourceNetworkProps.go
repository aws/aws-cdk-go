package awsdrs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSourceNetwork`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSourceNetworkProps := &CfnSourceNetworkProps{
//   	OriginAccountId: jsii.String("originAccountId"),
//   	OriginRegion: jsii.String("originRegion"),
//   	VpcId: jsii.String("vpcId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html
//
type CfnSourceNetworkProps struct {
	// The account ID containing the VPC to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-originaccountid
	//
	OriginAccountId *string `field:"required" json:"originAccountId" yaml:"originAccountId"`
	// The region containing the VPC to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-originregion
	//
	OriginRegion *string `field:"required" json:"originRegion" yaml:"originRegion"`
	// The VPC ID to protect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// A set of tags associated with the Source Network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-sourcenetwork.html#cfn-drs-sourcenetwork-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

