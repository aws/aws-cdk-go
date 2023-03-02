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
	// `AWS::MediaTailor::PlaybackConfiguration.AdDecisionServerUrl`.
	AdDecisionServerUrl *string `field:"required" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// `AWS::MediaTailor::PlaybackConfiguration.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::MediaTailor::PlaybackConfiguration.VideoContentSourceUrl`.
	VideoContentSourceUrl *string `field:"required" json:"videoContentSourceUrl" yaml:"videoContentSourceUrl"`
	// `AWS::MediaTailor::PlaybackConfiguration.AvailSuppression`.
	AvailSuppression interface{} `field:"optional" json:"availSuppression" yaml:"availSuppression"`
	// `AWS::MediaTailor::PlaybackConfiguration.Bumper`.
	Bumper interface{} `field:"optional" json:"bumper" yaml:"bumper"`
	// `AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration`.
	CdnConfiguration interface{} `field:"optional" json:"cdnConfiguration" yaml:"cdnConfiguration"`
	// The player parameters and aliases used as dynamic variables during session initialization.
	//
	// For more information, see [Domain Variables](https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html) .
	ConfigurationAliases interface{} `field:"optional" json:"configurationAliases" yaml:"configurationAliases"`
	// `AWS::MediaTailor::PlaybackConfiguration.DashConfiguration`.
	DashConfiguration interface{} `field:"optional" json:"dashConfiguration" yaml:"dashConfiguration"`
	// The configuration for HLS content.
	HlsConfiguration interface{} `field:"optional" json:"hlsConfiguration" yaml:"hlsConfiguration"`
	// `AWS::MediaTailor::PlaybackConfiguration.LivePreRollConfiguration`.
	LivePreRollConfiguration interface{} `field:"optional" json:"livePreRollConfiguration" yaml:"livePreRollConfiguration"`
	// `AWS::MediaTailor::PlaybackConfiguration.ManifestProcessingRules`.
	ManifestProcessingRules interface{} `field:"optional" json:"manifestProcessingRules" yaml:"manifestProcessingRules"`
	// `AWS::MediaTailor::PlaybackConfiguration.PersonalizationThresholdSeconds`.
	PersonalizationThresholdSeconds *float64 `field:"optional" json:"personalizationThresholdSeconds" yaml:"personalizationThresholdSeconds"`
	// `AWS::MediaTailor::PlaybackConfiguration.SlateAdUrl`.
	SlateAdUrl *string `field:"optional" json:"slateAdUrl" yaml:"slateAdUrl"`
	// `AWS::MediaTailor::PlaybackConfiguration.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::MediaTailor::PlaybackConfiguration.TranscodeProfileName`.
	TranscodeProfileName *string `field:"optional" json:"transcodeProfileName" yaml:"transcodeProfileName"`
}

