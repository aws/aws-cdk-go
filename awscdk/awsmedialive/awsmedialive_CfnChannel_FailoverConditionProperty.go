package awsmedialive


// Failover Condition settings. There can be multiple failover conditions inside AutomaticInputFailoverSettings.
//
// The parent of this entity is AutomaticInputFailoverSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConditionProperty := &failoverConditionProperty{
//   	failoverConditionSettings: &failoverConditionSettingsProperty{
//   		audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   			audioSelectorName: jsii.String("audioSelectorName"),
//   			audioSilenceThresholdMsec: jsii.Number(123),
//   		},
//   		inputLossSettings: &inputLossFailoverSettingsProperty{
//   			inputLossThresholdMsec: jsii.Number(123),
//   		},
//   		videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   			blackDetectThreshold: jsii.Number(123),
//   			videoBlackThresholdMsec: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_FailoverConditionProperty struct {
	// Settings for a specific failover condition.
	FailoverConditionSettings interface{} `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

