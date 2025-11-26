package previewawsmedialivemixins


// Settings to configure the conditions that will define the input as unhealthy and that will make MediaLive fail over to the other input in the input failover pair.
//
// The parent of this entity is InputAttachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   automaticInputFailoverSettingsProperty := &AutomaticInputFailoverSettingsProperty{
//   	ErrorClearTimeMsec: jsii.Number(123),
//   	FailoverConditions: []interface{}{
//   		&FailoverConditionProperty{
//   			FailoverConditionSettings: &FailoverConditionSettingsProperty{
//   				AudioSilenceSettings: &AudioSilenceFailoverSettingsProperty{
//   					AudioSelectorName: jsii.String("audioSelectorName"),
//   					AudioSilenceThresholdMsec: jsii.Number(123),
//   				},
//   				InputLossSettings: &InputLossFailoverSettingsProperty{
//   					InputLossThresholdMsec: jsii.Number(123),
//   				},
//   				VideoBlackSettings: &VideoBlackFailoverSettingsProperty{
//   					BlackDetectThreshold: jsii.Number(123),
//   					VideoBlackThresholdMsec: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	InputPreference: jsii.String("inputPreference"),
//   	SecondaryInputId: jsii.String("secondaryInputId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html
//
type CfnChannelPropsMixin_AutomaticInputFailoverSettingsProperty struct {
	// This clear time defines the requirement a recovered input must meet to be considered healthy.
	//
	// The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input_preference for the failover pair is set to PRIMARY_INPUT_PREFERRED, because after this time, MediaLive will switch back to the primary input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-errorcleartimemsec
	//
	ErrorClearTimeMsec *float64 `field:"optional" json:"errorClearTimeMsec" yaml:"errorClearTimeMsec"`
	// A list of failover conditions.
	//
	// If any of these conditions occur, MediaLive will perform a failover to the other input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-failoverconditions
	//
	FailoverConditions interface{} `field:"optional" json:"failoverConditions" yaml:"failoverConditions"`
	// Input preference when deciding which input to make active when a previously failed input has recovered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-inputpreference
	//
	InputPreference *string `field:"optional" json:"inputPreference" yaml:"inputPreference"`
	// The input ID of the secondary input in the automatic input failover pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-secondaryinputid
	//
	SecondaryInputId *string `field:"optional" json:"secondaryInputId" yaml:"secondaryInputId"`
}

