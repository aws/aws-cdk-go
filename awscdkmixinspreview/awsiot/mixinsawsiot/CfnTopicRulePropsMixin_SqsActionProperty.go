package mixinsawsiot


// Describes an action to publish data to an Amazon SQS queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sqsActionProperty := &SqsActionProperty{
//   	QueueUrl: jsii.String("queueUrl"),
//   	RoleArn: jsii.String("roleArn"),
//   	UseBase64: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html
//
type CfnTopicRulePropsMixin_SqsActionProperty struct {
	// The URL of the Amazon SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html#cfn-iot-topicrule-sqsaction-queueurl
	//
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
	// The ARN of the IAM role that grants access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html#cfn-iot-topicrule-sqsaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Specifies whether to use Base64 encoding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html#cfn-iot-topicrule-sqsaction-usebase64
	//
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

