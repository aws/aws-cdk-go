package awsses


// SNSAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sNSActionConfig := &sNSActionConfig{
//   	encoding: jsii.String("encoding"),
//   	topicArn: jsii.String("topicArn"),
//   }
//
type SNSActionConfig struct {
	// The encoding to use for the email within the Amazon SNS notification.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

