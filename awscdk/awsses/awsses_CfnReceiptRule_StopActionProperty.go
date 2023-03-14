package awsses


// When included in a receipt rule, this action terminates the evaluation of the receipt rule set and, optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).
//
// For information about setting a stop action in a receipt rule, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-stop.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stopActionProperty := &stopActionProperty{
//   	scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnReceiptRule_StopActionProperty struct {
	// The scope of the StopAction.
	//
	// The only acceptable value is `RuleSet` .
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the stop action is taken.
	//
	// You can find the ARN of a topic by using the [ListTopics](https://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html) Amazon SNS operation.
	//
	// For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) .
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

