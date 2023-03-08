package awspinpointemail


// An object that defines an Amazon SNS destination for email events.
//
// You can use Amazon SNS to send notification when certain email events occur.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snsDestinationProperty := &snsDestinationProperty{
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnConfigurationSetEventDestination_SnsDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish email events to.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

