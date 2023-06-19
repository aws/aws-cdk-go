package awsses


// SNSAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sNSActionConfig := &SNSActionConfig{
//   	Encoding: jsii.String("encoding"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// Experimental.
type SNSActionConfig struct {
	// The encoding to use for the email within the Amazon SNS notification.
	// Experimental.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify.
	// Experimental.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

