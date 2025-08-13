package awsmediapackagev2


// The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSwitchConfigurationProperty := &InputSwitchConfigurationProperty{
//   	MqcsInputSwitching: jsii.Boolean(false),
//   	PreferredInput: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-inputswitchconfiguration.html
//
type CfnChannel_InputSwitchConfigurationProperty struct {
	// When true, AWS Elemental MediaPackage performs input switching based on the MQCS.
	//
	// Default is true. This setting is valid only when `InputType` is `CMAF` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-inputswitchconfiguration.html#cfn-mediapackagev2-channel-inputswitchconfiguration-mqcsinputswitching
	//
	MqcsInputSwitching interface{} `field:"optional" json:"mqcsInputSwitching" yaml:"mqcsInputSwitching"`
	// For CMAF inputs, indicates which input MediaPackage should prefer when both inputs have equal MQCS scores.
	//
	// Select `1` to prefer the first ingest endpoint, or `2` to prefer the second ingest endpoint. If you don't specify a preferred input, MediaPackage uses its default switching behavior when MQCS scores are equal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-inputswitchconfiguration.html#cfn-mediapackagev2-channel-inputswitchconfiguration-preferredinput
	//
	PreferredInput *float64 `field:"optional" json:"preferredInput" yaml:"preferredInput"`
}

