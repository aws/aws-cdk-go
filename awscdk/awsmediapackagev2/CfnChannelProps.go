package awsmediapackagev2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &CfnChannelProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	InputSwitchConfiguration: &InputSwitchConfigurationProperty{
//   		MqcsInputSwitching: jsii.Boolean(false),
//   	},
//   	InputType: jsii.String("inputType"),
//   	OutputHeaderConfiguration: &OutputHeaderConfigurationProperty{
//   		PublishMqcs: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html
//
type CfnChannelProps struct {
	// The name of the channel group associated with the channel configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html#cfn-mediapackagev2-channel-channelgroupname
	//
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The name of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html#cfn-mediapackagev2-channel-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The description of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html#cfn-mediapackagev2-channel-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html#cfn-mediapackagev2-channel-inputswitchconfiguration
	//
	InputSwitchConfiguration interface{} `field:"optional" json:"inputSwitchConfiguration" yaml:"inputSwitchConfiguration"`
	// The input type will be an immutable field which will be used to define whether the channel will allow CMAF ingest or HLS ingest.
	//
	// If unprovided, it will default to HLS to preserve current behavior.
	//
	// The allowed values are:
	//
	// - `HLS` - The HLS streaming specification (which defines M3U8 manifests and TS segments).
	// - `CMAF` - The DASH-IF CMAF Ingest specification (which defines CMAF segments with optional DASH manifests).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html#cfn-mediapackagev2-channel-inputtype
	//
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html#cfn-mediapackagev2-channel-outputheaderconfiguration
	//
	OutputHeaderConfiguration interface{} `field:"optional" json:"outputHeaderConfiguration" yaml:"outputHeaderConfiguration"`
	// The tags associated with the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html#cfn-mediapackagev2-channel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

