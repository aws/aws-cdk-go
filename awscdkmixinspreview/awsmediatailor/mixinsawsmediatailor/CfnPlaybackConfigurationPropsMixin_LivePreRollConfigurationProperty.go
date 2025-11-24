package mixinsawsmediatailor


// The configuration for pre-roll ad insertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   livePreRollConfigurationProperty := &LivePreRollConfigurationProperty{
//   	AdDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	MaxDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-liveprerollconfiguration.html
//
type CfnPlaybackConfigurationPropsMixin_LivePreRollConfigurationProperty struct {
	// The URL for the ad decision server (ADS) for pre-roll ads.
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-liveprerollconfiguration.html#cfn-mediatailor-playbackconfiguration-liveprerollconfiguration-addecisionserverurl
	//
	AdDecisionServerUrl *string `field:"optional" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// The maximum allowed duration for the pre-roll ad avail.
	//
	// AWS Elemental MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-liveprerollconfiguration.html#cfn-mediatailor-playbackconfiguration-liveprerollconfiguration-maxdurationseconds
	//
	MaxDurationSeconds *float64 `field:"optional" json:"maxDurationSeconds" yaml:"maxDurationSeconds"`
}

