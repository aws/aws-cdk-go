package awsmediapackagev2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannelGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelGroupProps := &CfnChannelGroupProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelgroup.html
//
type CfnChannelGroupProps struct {
	// The name of the channel group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelgroup.html#cfn-mediapackagev2-channelgroup-channelgroupname
	//
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The configuration for a MediaPackage V2 channel group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelgroup.html#cfn-mediapackagev2-channelgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with the channel group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelgroup.html#cfn-mediapackagev2-channelgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

