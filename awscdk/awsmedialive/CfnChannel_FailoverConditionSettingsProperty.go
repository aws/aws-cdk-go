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
//   failoverConditionSettingsProperty := &FailoverConditionSettingsProperty{
//   	AudioSilenceSettings: &AudioSilenceFailoverSettingsProperty{
//   		AudioSelectorName: jsii.String("audioSelectorName"),
//   		AudioSilenceThresholdMsec: jsii.Number(123),
//   	},
//   	InputLossSettings: &InputLossFailoverSettingsProperty{
//   		InputLossThresholdMsec: jsii.Number(123),
//   	},
//   	VideoBlackSettings: &VideoBlackFailoverSettingsProperty{
//   		BlackDetectThreshold: jsii.Number(123),
//   		VideoBlackThresholdMsec: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failoverconditionsettings.html
//
type CfnChannel_FailoverConditionSettingsProperty struct {
	// MediaLive will perform a failover if the specified audio selector is silent for the specified period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failoverconditionsettings.html#cfn-medialive-channel-failoverconditionsettings-audiosilencesettings
	//
	AudioSilenceSettings interface{} `field:"optional" json:"audioSilenceSettings" yaml:"audioSilenceSettings"`
	// MediaLive will perform a failover if content is not detected in this input for the specified period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failoverconditionsettings.html#cfn-medialive-channel-failoverconditionsettings-inputlosssettings
	//
	InputLossSettings interface{} `field:"optional" json:"inputLossSettings" yaml:"inputLossSettings"`
	// MediaLive will perform a failover if content is considered black for the specified period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failoverconditionsettings.html#cfn-medialive-channel-failoverconditionsettings-videoblacksettings
	//
	VideoBlackSettings interface{} `field:"optional" json:"videoBlackSettings" yaml:"videoBlackSettings"`
}

