package mixinsawsmedialive


// Failover Condition settings. There can be multiple failover conditions inside AutomaticInputFailoverSettings.
//
// The parent of this entity is AutomaticInputFailoverSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   failoverConditionProperty := &FailoverConditionProperty{
//   	FailoverConditionSettings: &FailoverConditionSettingsProperty{
//   		AudioSilenceSettings: &AudioSilenceFailoverSettingsProperty{
//   			AudioSelectorName: jsii.String("audioSelectorName"),
//   			AudioSilenceThresholdMsec: jsii.Number(123),
//   		},
//   		InputLossSettings: &InputLossFailoverSettingsProperty{
//   			InputLossThresholdMsec: jsii.Number(123),
//   		},
//   		VideoBlackSettings: &VideoBlackFailoverSettingsProperty{
//   			BlackDetectThreshold: jsii.Number(123),
//   			VideoBlackThresholdMsec: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failovercondition.html
//
type CfnChannelPropsMixin_FailoverConditionProperty struct {
	// Settings for a specific failover condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failovercondition.html#cfn-medialive-channel-failovercondition-failoverconditionsettings
	//
	FailoverConditionSettings interface{} `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

