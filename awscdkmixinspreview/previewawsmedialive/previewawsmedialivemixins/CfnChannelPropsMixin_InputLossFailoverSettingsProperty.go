package previewawsmedialivemixins


// MediaLive will perform a failover if content is not detected in this input for the specified period.
//
// The parent of this entity is FailoverConditionSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputLossFailoverSettingsProperty := &InputLossFailoverSettingsProperty{
//   	InputLossThresholdMsec: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossfailoversettings.html
//
type CfnChannelPropsMixin_InputLossFailoverSettingsProperty struct {
	// The amount of time (in milliseconds) that no input is detected.
	//
	// After that time, an input failover will occur.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossfailoversettings.html#cfn-medialive-channel-inputlossfailoversettings-inputlossthresholdmsec
	//
	InputLossThresholdMsec *float64 `field:"optional" json:"inputLossThresholdMsec" yaml:"inputLossThresholdMsec"`
}

