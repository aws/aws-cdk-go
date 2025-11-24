package mixinsawsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnChannelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelMixinProps := &CfnChannelMixinProps{
//   	Audiences: []*string{
//   		jsii.String("audiences"),
//   	},
//   	ChannelName: jsii.String("channelName"),
//   	FillerSlate: &SlateSourceProperty{
//   		SourceLocationName: jsii.String("sourceLocationName"),
//   		VodSourceName: jsii.String("vodSourceName"),
//   	},
//   	LogConfiguration: &LogConfigurationForChannelProperty{
//   		LogTypes: []*string{
//   			jsii.String("logTypes"),
//   		},
//   	},
//   	Outputs: []interface{}{
//   		&RequestOutputItemProperty{
//   			DashPlaylistSettings: &DashPlaylistSettingsProperty{
//   				ManifestWindowSeconds: jsii.Number(123),
//   				MinBufferTimeSeconds: jsii.Number(123),
//   				MinUpdatePeriodSeconds: jsii.Number(123),
//   				SuggestedPresentationDelaySeconds: jsii.Number(123),
//   			},
//   			HlsPlaylistSettings: &HlsPlaylistSettingsProperty{
//   				AdMarkupType: []*string{
//   					jsii.String("adMarkupType"),
//   				},
//   				ManifestWindowSeconds: jsii.Number(123),
//   			},
//   			ManifestName: jsii.String("manifestName"),
//   			SourceGroup: jsii.String("sourceGroup"),
//   		},
//   	},
//   	PlaybackMode: jsii.String("playbackMode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tier: jsii.String("tier"),
//   	TimeShiftConfiguration: &TimeShiftConfigurationProperty{
//   		MaxTimeDelaySeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html
//
type CfnChannelMixinProps struct {
	// The list of audiences defined in channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-audiences
	//
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// The name of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The slate used to fill gaps between programs in the schedule.
	//
	// You must configure filler slate if your channel uses the `LINEAR` `PlaybackMode` . MediaTailor doesn't support filler slate for channels using the `LOOP` `PlaybackMode` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-fillerslate
	//
	FillerSlate interface{} `field:"optional" json:"fillerSlate" yaml:"fillerSlate"`
	// The log configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The channel's output properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// The type of playback mode for this channel.
	//
	// `LINEAR` - Programs play back-to-back only once.
	//
	// `LOOP` - Programs play back-to-back in an endless loop. When the last program in the schedule plays, playback loops back to the first program in the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-playbackmode
	//
	PlaybackMode *string `field:"optional" json:"playbackMode" yaml:"playbackMode"`
	// The tags to assign to the channel.
	//
	// Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The tier for this channel.
	//
	// STANDARD tier channels can contain live programs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// The configuration for time-shifted viewing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html#cfn-mediatailor-channel-timeshiftconfiguration
	//
	TimeShiftConfiguration interface{} `field:"optional" json:"timeShiftConfiguration" yaml:"timeShiftConfiguration"`
}

