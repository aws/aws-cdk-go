package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPlaybackConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationAliases interface{}
//
//   cfnPlaybackConfigurationProps := &CfnPlaybackConfigurationProps{
//   	AdDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	Name: jsii.String("name"),
//   	VideoContentSourceUrl: jsii.String("videoContentSourceUrl"),
//
//   	// the properties below are optional
//   	AdConditioningConfiguration: &AdConditioningConfigurationProperty{
//   		StreamingMediaFileConditioning: jsii.String("streamingMediaFileConditioning"),
//   	},
//   	AvailSuppression: &AvailSuppressionProperty{
//   		FillPolicy: jsii.String("fillPolicy"),
//   		Mode: jsii.String("mode"),
//   		Value: jsii.String("value"),
//   	},
//   	Bumper: &BumperProperty{
//   		EndUrl: jsii.String("endUrl"),
//   		StartUrl: jsii.String("startUrl"),
//   	},
//   	CdnConfiguration: &CdnConfigurationProperty{
//   		AdSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   		ContentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   	},
//   	ConfigurationAliases: map[string]interface{}{
//   		"configurationAliasesKey": configurationAliases,
//   	},
//   	DashConfiguration: &DashConfigurationProperty{
//   		ManifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   		MpdLocation: jsii.String("mpdLocation"),
//   		OriginManifestType: jsii.String("originManifestType"),
//   	},
//   	HlsConfiguration: &HlsConfigurationProperty{
//   		ManifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   	},
//   	InsertionMode: jsii.String("insertionMode"),
//   	LivePreRollConfiguration: &LivePreRollConfigurationProperty{
//   		AdDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   		MaxDurationSeconds: jsii.Number(123),
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		PercentEnabled: jsii.Number(123),
//
//   		// the properties below are optional
//   		AdsInteractionLog: &AdsInteractionLogProperty{
//   			ExcludeEventTypes: []*string{
//   				jsii.String("excludeEventTypes"),
//   			},
//   			PublishOptInEventTypes: []*string{
//   				jsii.String("publishOptInEventTypes"),
//   			},
//   		},
//   		EnabledLoggingStrategies: []*string{
//   			jsii.String("enabledLoggingStrategies"),
//   		},
//   		ManifestServiceInteractionLog: &ManifestServiceInteractionLogProperty{
//   			ExcludeEventTypes: []*string{
//   				jsii.String("excludeEventTypes"),
//   			},
//   		},
//   	},
//   	ManifestProcessingRules: &ManifestProcessingRulesProperty{
//   		AdMarkerPassthrough: &AdMarkerPassthroughProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	PersonalizationThresholdSeconds: jsii.Number(123),
//   	SlateAdUrl: jsii.String("slateAdUrl"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TranscodeProfileName: jsii.String("transcodeProfileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html
//
type CfnPlaybackConfigurationProps struct {
	// The URL for the ad decision server (ADS).
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-addecisionserverurl
	//
	AdDecisionServerUrl *string `field:"required" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// The identifier for the playback configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL prefix for the parent manifest for the stream, minus the asset ID.
	//
	// The maximum length is 512 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-videocontentsourceurl
	//
	VideoContentSourceUrl *string `field:"required" json:"videoContentSourceUrl" yaml:"videoContentSourceUrl"`
	// The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns, and what priority MediaTailor uses when inserting ads.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-adconditioningconfiguration
	//
	AdConditioningConfiguration interface{} `field:"optional" json:"adConditioningConfiguration" yaml:"adConditioningConfiguration"`
	// The configuration for avail suppression, also known as ad suppression.
	//
	// For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-availsuppression
	//
	AvailSuppression interface{} `field:"optional" json:"availSuppression" yaml:"availSuppression"`
	// The configuration for bumpers.
	//
	// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-bumper
	//
	Bumper interface{} `field:"optional" json:"bumper" yaml:"bumper"`
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-cdnconfiguration
	//
	CdnConfiguration interface{} `field:"optional" json:"cdnConfiguration" yaml:"cdnConfiguration"`
	// The player parameters and aliases used as dynamic variables during session initialization.
	//
	// For more information, see [Domain Variables](https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-configurationaliases
	//
	ConfigurationAliases interface{} `field:"optional" json:"configurationAliases" yaml:"configurationAliases"`
	// The configuration for a DASH source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-dashconfiguration
	//
	DashConfiguration interface{} `field:"optional" json:"dashConfiguration" yaml:"dashConfiguration"`
	// The configuration for HLS content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-hlsconfiguration
	//
	HlsConfiguration interface{} `field:"optional" json:"hlsConfiguration" yaml:"hlsConfiguration"`
	// The setting that controls whether players can use stitched or guided ad insertion.
	//
	// The default, `STITCHED_ONLY` , forces all player sessions to use stitched (server-side) ad insertion. Choosing `PLAYER_SELECT` allows players to select either stitched or guided ad insertion at session-initialization time. The default for players that do not specify an insertion mode is stitched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-insertionmode
	//
	InsertionMode *string `field:"optional" json:"insertionMode" yaml:"insertionMode"`
	// The configuration for pre-roll ad insertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-liveprerollconfiguration
	//
	LivePreRollConfiguration interface{} `field:"optional" json:"livePreRollConfiguration" yaml:"livePreRollConfiguration"`
	// Defines where AWS Elemental MediaTailor sends logs for the playback configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The configuration for manifest processing rules.
	//
	// Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-manifestprocessingrules
	//
	ManifestProcessingRules interface{} `field:"optional" json:"manifestProcessingRules" yaml:"manifestProcessingRules"`
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	//
	// If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to *ad replacement* in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see [Ad Behavior in AWS Elemental MediaTailor](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-personalizationthresholdseconds
	//
	PersonalizationThresholdSeconds *float64 `field:"optional" json:"personalizationThresholdSeconds" yaml:"personalizationThresholdSeconds"`
	// The URL for a video asset to transcode and use to fill in time that's not used by ads.
	//
	// AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID playback configurations. For VPAID, the slate is required because MediaTailor provides it in the slots designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-slateadurl
	//
	SlateAdUrl *string `field:"optional" json:"slateAdUrl" yaml:"slateAdUrl"`
	// The tags to assign to the playback configuration.
	//
	// Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name that is used to associate this playback configuration with a custom transcode profile.
	//
	// This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html#cfn-mediatailor-playbackconfiguration-transcodeprofilename
	//
	TranscodeProfileName *string `field:"optional" json:"transcodeProfileName" yaml:"transcodeProfileName"`
}

