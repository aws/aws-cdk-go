package awsmediatailor

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
//   	ChannelName: jsii.String("channelName"),
//   	Outputs: []interface{}{
//   		&RequestOutputItemProperty{
//   			ManifestName: jsii.String("manifestName"),
//   			SourceGroup: jsii.String("sourceGroup"),
//
//   			// the properties below are optional
//   			DashPlaylistSettings: &DashPlaylistSettingsProperty{
//   				ManifestWindowSeconds: jsii.Number(123),
//   				MinBufferTimeSeconds: jsii.Number(123),
//   				MinUpdatePeriodSeconds: jsii.Number(123),
//   				SuggestedPresentationDelaySeconds: jsii.Number(123),
//   			},
//   			HlsPlaylistSettings: &HlsPlaylistSettingsProperty{
//   				ManifestWindowSeconds: jsii.Number(123),
//   			},
//   		},
//   	},
//   	PlaybackMode: jsii.String("playbackMode"),
//
//   	// the properties below are optional
//   	FillerSlate: &SlateSourceProperty{
//   		SourceLocationName: jsii.String("sourceLocationName"),
//   		VodSourceName: jsii.String("vodSourceName"),
//   	},
//   	LogConfiguration: &LogConfigurationForChannelProperty{
//   		LogTypes: []*string{
//   			jsii.String("logTypes"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tier: jsii.String("tier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html
//
type CfnChannelProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// <p>The channel's output properties.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-outputs
	//
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-playbackmode
	//
	PlaybackMode *string `field:"required" json:"playbackMode" yaml:"playbackMode"`
	// <p>Slate VOD source configuration.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-fillerslate
	//
	FillerSlate interface{} `field:"optional" json:"fillerSlate" yaml:"fillerSlate"`
	// <p>The log configuration for the channel.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The tags to assign to the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

