package previewawsiotmixins


// Describes an action to publish to an Amazon SNS topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snsActionProperty := &SnsActionProperty{
//   	MessageFormat: jsii.String("messageFormat"),
//   	RoleArn: jsii.String("roleArn"),
//   	TargetArn: jsii.String("targetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html
//
type CfnTopicRulePropsMixin_SnsActionProperty struct {
	// (Optional) The message format of the message to publish.
	//
	// Accepted values are "JSON" and "RAW". The default value of the attribute is "RAW". SNS uses this setting to determine if the payload should be parsed and relevant platform-specific bits of the payload should be extracted. For more information, see [Amazon SNS Message and JSON Formats](https://docs.aws.amazon.com/sns/latest/dg/json-formats.html) in the *Amazon Simple Notification Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-messageformat
	//
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// The ARN of the IAM role that grants access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the SNS topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-targetarn
	//
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

