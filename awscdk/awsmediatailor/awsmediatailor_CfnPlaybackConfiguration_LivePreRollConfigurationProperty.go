package awsmediatailor


// The configuration for pre-roll ad insertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   livePreRollConfigurationProperty := &livePreRollConfigurationProperty{
//   	adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	maxDurationSeconds: jsii.Number(123),
//   }
//
type CfnPlaybackConfiguration_LivePreRollConfigurationProperty struct {
	// The URL for the ad decision server (ADS) for pre-roll ads.
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `field:"optional" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// The maximum allowed duration for the pre-roll ad avail.
	//
	// MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
	MaxDurationSeconds *float64 `field:"optional" json:"maxDurationSeconds" yaml:"maxDurationSeconds"`
}

