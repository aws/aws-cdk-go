package awsiot


// Describes an action to publish to an Amazon SNS topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsActionProperty := &snsActionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	messageFormat: jsii.String("messageFormat"),
//   }
//
type CfnTopicRule_SnsActionProperty struct {
	// The ARN of the IAM role that grants access.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the SNS topic.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// (Optional) The message format of the message to publish.
	//
	// Accepted values are "JSON" and "RAW". The default value of the attribute is "RAW". SNS uses this setting to determine if the payload should be parsed and relevant platform-specific bits of the payload should be extracted. For more information, see [Amazon SNS Message and JSON Formats](https://docs.aws.amazon.com/sns/latest/dg/json-formats.html) in the *Amazon Simple Notification Service Developer Guide* .
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
}

