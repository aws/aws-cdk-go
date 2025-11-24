package mixinsawsses


// The action to publish the email content to an Amazon SNS topic.
//
// When executed, this action will send the email as a notification to the specified SNS topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sNSActionProperty := &SNSActionProperty{
//   	Encoding: jsii.String("encoding"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-snsaction.html
//
type CfnReceiptRulePropsMixin_SNSActionProperty struct {
	// The encoding to use for the email within the Amazon SNS notification.
	//
	// The default value is `UTF-8` . Use `BASE64` if you need to preserve all special characters, especially when the original message uses a different encoding format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-snsaction.html#cfn-ses-receiptrule-snsaction-encoding
	//
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The Amazon Resource Name (ARN) of the Amazon SNS Topic to which notification for the email received will be published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-snsaction.html#cfn-ses-receiptrule-snsaction-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

