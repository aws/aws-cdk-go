package awsmedialive


// Settings to configure the conditions that will define the input as unhealthy and that will make MediaLive fail over to the other input in the input failover pair.
//
// The parent of this entity is InputAttachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automaticInputFailoverSettingsProperty := &automaticInputFailoverSettingsProperty{
//   	errorClearTimeMsec: jsii.Number(123),
//   	failoverConditions: []interface{}{
//   		&failoverConditionProperty{
//   			failoverConditionSettings: &failoverConditionSettingsProperty{
//   				audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   					audioSelectorName: jsii.String("audioSelectorName"),
//   					audioSilenceThresholdMsec: jsii.Number(123),
//   				},
//   				inputLossSettings: &inputLossFailoverSettingsProperty{
//   					inputLossThresholdMsec: jsii.Number(123),
//   				},
//   				videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   					blackDetectThreshold: jsii.Number(123),
//   					videoBlackThresholdMsec: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	inputPreference: jsii.String("inputPreference"),
//   	secondaryInputId: jsii.String("secondaryInputId"),
//   }
//
type CfnChannel_AutomaticInputFailoverSettingsProperty struct {
	// This clear time defines the requirement a recovered input must meet to be considered healthy.
	//
	// The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input_preference for the failover pair is set to PRIMARY_INPUT_PREFERRED, because after this time, MediaLive will switch back to the primary input.
	ErrorClearTimeMsec *float64 `field:"optional" json:"errorClearTimeMsec" yaml:"errorClearTimeMsec"`
	// A list of failover conditions.
	//
	// If any of these conditions occur, MediaLive will perform a failover to the other input.
	FailoverConditions interface{} `field:"optional" json:"failoverConditions" yaml:"failoverConditions"`
	// Input preference when deciding which input to make active when a previously failed input has recovered.
	InputPreference *string `field:"optional" json:"inputPreference" yaml:"inputPreference"`
	// The input ID of the secondary input in the automatic input failover pair.
	SecondaryInputId *string `field:"optional" json:"secondaryInputId" yaml:"secondaryInputId"`
}

