package awsmediapackagev2


// <p>The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSwitchConfigurationProperty := &InputSwitchConfigurationProperty{
//   	MqcsInputSwitching: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-inputswitchconfiguration.html
//
type CfnChannel_InputSwitchConfigurationProperty struct {
	// <p>When true, AWS Elemental MediaPackage performs input switching based on the MQCS.
	//
	// Default is true. This setting is valid only when <code>InputType</code> is <code>CMAF</code>.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-inputswitchconfiguration.html#cfn-mediapackagev2-channel-inputswitchconfiguration-mqcsinputswitching
	//
	MqcsInputSwitching interface{} `field:"optional" json:"mqcsInputSwitching" yaml:"mqcsInputSwitching"`
}

