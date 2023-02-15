package awsmediatailor


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
	// `CfnPlaybackConfiguration.LivePreRollConfigurationProperty.AdDecisionServerUrl`.
	AdDecisionServerUrl *string `field:"optional" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// `CfnPlaybackConfiguration.LivePreRollConfigurationProperty.MaxDurationSeconds`.
	MaxDurationSeconds *float64 `field:"optional" json:"maxDurationSeconds" yaml:"maxDurationSeconds"`
}

