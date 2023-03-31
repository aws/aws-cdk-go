package awss3


// Specifies the configuration for publishing messages to an Amazon Simple Queue Service (Amazon SQS) queue when Amazon S3 detects specified events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueConfigurationProperty := &queueConfigurationProperty{
//   	event: jsii.String("event"),
//   	queue: jsii.String("queue"),
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
type CfnBucket_QueueConfigurationProperty struct {
	// The Amazon S3 bucket event about which you want to publish messages to Amazon SQS.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type.
	//
	// FIFO queues are not allowed when enabling an SQS queue as the event notification destination.
	Queue *string `field:"required" json:"queue" yaml:"queue"`
	// The filtering rules that determine which objects trigger notifications.
	//
	// For example, you can create a filter so that Amazon S3 sends notifications only when image files with a `.jpg` extension are added to the bucket. For more information, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/user-guide/notification-how-to-filtering.html) in the *Amazon S3 User Guide* .
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

