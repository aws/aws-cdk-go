package awsmedialive


// Settings for one failover condition.
//
// The parent of this entity is FailoverCondition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConditionSettingsProperty := &failoverConditionSettingsProperty{
//   	audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   		audioSelectorName: jsii.String("audioSelectorName"),
//   		audioSilenceThresholdMsec: jsii.Number(123),
//   	},
//   	inputLossSettings: &inputLossFailoverSettingsProperty{
//   		inputLossThresholdMsec: jsii.Number(123),
//   	},
//   	videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   		blackDetectThreshold: jsii.Number(123),
//   		videoBlackThresholdMsec: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_FailoverConditionSettingsProperty struct {
	// MediaLive will perform a failover if the specified audio selector is silent for the specified period.
	AudioSilenceSettings interface{} `field:"optional" json:"audioSilenceSettings" yaml:"audioSilenceSettings"`
	// MediaLive will perform a failover if content is not detected in this input for the specified period.
	InputLossSettings interface{} `field:"optional" json:"inputLossSettings" yaml:"inputLossSettings"`
	// MediaLive will perform a failover if content is considered black for the specified period.
	VideoBlackSettings interface{} `field:"optional" json:"videoBlackSettings" yaml:"videoBlackSettings"`
}

