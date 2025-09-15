package awsmediatailor


// Defines where AWS Elemental MediaTailor sends logs for the playback configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	PercentEnabled: jsii.Number(123),
//
//   	// the properties below are optional
//   	AdsInteractionLog: &AdsInteractionLogProperty{
//   		ExcludeEventTypes: []*string{
//   			jsii.String("excludeEventTypes"),
//   		},
//   		PublishOptInEventTypes: []*string{
//   			jsii.String("publishOptInEventTypes"),
//   		},
//   	},
//   	EnabledLoggingStrategies: []*string{
//   		jsii.String("enabledLoggingStrategies"),
//   	},
//   	ManifestServiceInteractionLog: &ManifestServiceInteractionLogProperty{
//   		ExcludeEventTypes: []*string{
//   			jsii.String("excludeEventTypes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-logconfiguration.html
//
type CfnPlaybackConfiguration_LogConfigurationProperty struct {
	// The percentage of session logs that MediaTailor sends to your configured log destination.
	//
	// For example, if your playback configuration has 1000 sessions and `percentEnabled` is set to `60` , MediaTailor sends logs for 600 of the sessions to CloudWatch Logs. MediaTailor decides at random which of the playback configuration sessions to send logs for. If you want to view logs for a specific session, you can use the [debug log mode](https://docs.aws.amazon.com/mediatailor/latest/ug/debug-log-mode.html) .
	//
	// Valid values: `0` - `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-logconfiguration.html#cfn-mediatailor-playbackconfiguration-logconfiguration-percentenabled
	//
	PercentEnabled *float64 `field:"required" json:"percentEnabled" yaml:"percentEnabled"`
	// Settings for customizing what events are included in logs for interactions with the ad decision server (ADS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-logconfiguration.html#cfn-mediatailor-playbackconfiguration-logconfiguration-adsinteractionlog
	//
	AdsInteractionLog interface{} `field:"optional" json:"adsInteractionLog" yaml:"adsInteractionLog"`
	// The method used for collecting logs from AWS Elemental MediaTailor.
	//
	// `LEGACY_CLOUDWATCH` indicates that MediaTailor is sending logs directly to Amazon CloudWatch Logs. `VENDED_LOGS` indicates that MediaTailor is sending logs to CloudWatch, which then vends the logs to your destination of choice. Supported destinations are CloudWatch Logs log group, Amazon S3 bucket, and Amazon Data Firehose stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-logconfiguration.html#cfn-mediatailor-playbackconfiguration-logconfiguration-enabledloggingstrategies
	//
	EnabledLoggingStrategies *[]*string `field:"optional" json:"enabledLoggingStrategies" yaml:"enabledLoggingStrategies"`
	// Settings for customizing what events are included in logs for interactions with the origin server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-logconfiguration.html#cfn-mediatailor-playbackconfiguration-logconfiguration-manifestserviceinteractionlog
	//
	ManifestServiceInteractionLog interface{} `field:"optional" json:"manifestServiceInteractionLog" yaml:"manifestServiceInteractionLog"`
}

