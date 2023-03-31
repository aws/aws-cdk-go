package awss3


// A container for specifying the configuration for publication of messages to an Amazon Simple Notification Service (Amazon SNS) topic when Amazon S3 detects specified events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicConfigurationProperty := &topicConfigurationProperty{
//   	event: jsii.String("event"),
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	filter: &notificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&filterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_TopicConfigurationProperty struct {
	// The Amazon S3 bucket event about which to send notifications.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it detects events of the specified type.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The filtering rules that determine for which objects to send notifications.
	//
	// For example, you can create a filter so that Amazon S3 sends notifications only when image files with a `.jpg` extension are added to the bucket.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

