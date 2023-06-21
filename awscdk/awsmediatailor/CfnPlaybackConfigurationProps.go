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
//   	AvailSuppression: &AvailSuppressionProperty{
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
//   	LivePreRollConfiguration: &LivePreRollConfigurationProperty{
//   		AdDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   		MaxDurationSeconds: jsii.Number(123),
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

