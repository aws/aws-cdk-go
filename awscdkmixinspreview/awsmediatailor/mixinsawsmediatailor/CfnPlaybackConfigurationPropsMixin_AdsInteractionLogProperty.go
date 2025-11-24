package mixinsawsmediatailor


// Settings for customizing what events are included in logs for interactions with the ad decision server (ADS).
//
// For more information about ADS logs, inlcuding descriptions of the event types, see [MediaTailor ADS logs description and event types](https://docs.aws.amazon.com/mediatailor/latest/ug/ads-log-format.html) in AWS Elemental MediaTailor User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   adsInteractionLogProperty := &AdsInteractionLogProperty{
//   	ExcludeEventTypes: []*string{
//   		jsii.String("excludeEventTypes"),
//   	},
//   	PublishOptInEventTypes: []*string{
//   		jsii.String("publishOptInEventTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adsinteractionlog.html
//
type CfnPlaybackConfigurationPropsMixin_AdsInteractionLogProperty struct {
	// Indicates that MediaTailor won't emit the selected events in the logs for playback sessions that are initialized with this configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adsinteractionlog.html#cfn-mediatailor-playbackconfiguration-adsinteractionlog-excludeeventtypes
	//
	ExcludeEventTypes *[]*string `field:"optional" json:"excludeEventTypes" yaml:"excludeEventTypes"`
	// Indicates that MediaTailor emits `RAW_ADS_RESPONSE` logs for playback sessions that are initialized with this configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adsinteractionlog.html#cfn-mediatailor-playbackconfiguration-adsinteractionlog-publishoptineventtypes
	//
	PublishOptInEventTypes *[]*string `field:"optional" json:"publishOptInEventTypes" yaml:"publishOptInEventTypes"`
}

