package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnPlaybackConfigurationProps := &cfnPlaybackConfigurationProps{
//   	adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	name: jsii.String("name"),
//   	videoContentSourceUrl: jsii.String("videoContentSourceUrl"),
//
//   	// the properties below are optional
//   	availSuppression: &availSuppressionProperty{
//   		mode: jsii.String("mode"),
//   		value: jsii.String("value"),
//   	},
//   	bumper: &bumperProperty{
//   		endUrl: jsii.String("endUrl"),
//   		startUrl: jsii.String("startUrl"),
//   	},
//   	cdnConfiguration: &cdnConfigurationProperty{
//   		adSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   		contentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   	},
//   	configurationAliases: map[string]interface{}{
//   		"configurationAliasesKey": configurationAliases,
//   	},
//   	dashConfiguration: &dashConfigurationProperty{
//   		manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   		mpdLocation: jsii.String("mpdLocation"),
//   		originManifestType: jsii.String("originManifestType"),
//   	},
//   	hlsConfiguration: &hlsConfigurationProperty{
//   		manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   	},
//   	livePreRollConfiguration: &livePreRollConfigurationProperty{
//   		adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   		maxDurationSeconds: jsii.Number(123),
//   	},
//   	manifestProcessingRules: &manifestProcessingRulesProperty{
//   		adMarkerPassthrough: &adMarkerPassthroughProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	personalizationThresholdSeconds: jsii.Number(123),
//   	slateAdUrl: jsii.String("slateAdUrl"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transcodeProfileName: jsii.String("transcodeProfileName"),
//   }
//
type CfnPlaybackConfigurationProps struct {
	// The URL for the ad decision server (ADS).
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `field:"required" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// The identifier for the playback configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL prefix for the parent manifest for the stream, minus the asset ID.
	//
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `field:"required" json:"videoContentSourceUrl" yaml:"videoContentSourceUrl"`
	// The configuration for avail suppression, also known as ad suppression.
	//
	// For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	AvailSuppression interface{} `field:"optional" json:"availSuppression" yaml:"availSuppression"`
	// The configuration for bumpers.
	//
	// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html) .
	Bumper interface{} `field:"optional" json:"bumper" yaml:"bumper"`
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration interface{} `field:"optional" json:"cdnConfiguration" yaml:"cdnConfiguration"`
	// The player parameters and aliases used as dynamic variables during session initialization.
	//
	// For more information, see [Domain Variables](https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html) .
	ConfigurationAliases interface{} `field:"optional" json:"configurationAliases" yaml:"configurationAliases"`
	// The configuration for DASH content.
	DashConfiguration interface{} `field:"optional" json:"dashConfiguration" yaml:"dashConfiguration"`
	// `AWS::MediaTailor::PlaybackConfiguration.HlsConfiguration`.
	HlsConfiguration interface{} `field:"optional" json:"hlsConfiguration" yaml:"hlsConfiguration"`
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration interface{} `field:"optional" json:"livePreRollConfiguration" yaml:"livePreRollConfiguration"`
	// The configuration for manifest processing rules.
	//
	// Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules interface{} `field:"optional" json:"manifestProcessingRules" yaml:"manifestProcessingRules"`
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	//
	// If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to *ad replacement* in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see [Ad Behavior in MediaTailor](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	PersonalizationThresholdSeconds *float64 `field:"optional" json:"personalizationThresholdSeconds" yaml:"personalizationThresholdSeconds"`
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads.
	//
	// MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `field:"optional" json:"slateAdUrl" yaml:"slateAdUrl"`
	// The tags to assign to the playback configuration.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name that is used to associate this playback configuration with a custom transcode profile.
	//
	// This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName *string `field:"optional" json:"transcodeProfileName" yaml:"transcodeProfileName"`
}

