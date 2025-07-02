package awss3


// Specifies the configuration for publishing messages to an Amazon Simple Queue Service (Amazon SQS) queue when Amazon S3 detects specified events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueConfigurationProperty := &QueueConfigurationProperty{
//   	Event: jsii.String("event"),
//   	Queue: jsii.String("queue"),
//
//   	// the properties below are optional
//   	Filter: &NotificationFilterProperty{
//   		S3Key: &S3KeyFilterProperty{
//   			Rules: []interface{}{
//   				&FilterRuleProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-queueconfiguration.html
//
type CfnBucket_QueueConfigurationProperty struct {
	// The Amazon S3 bucket event about which you want to publish messages to Amazon SQS.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-queueconfiguration.html#cfn-s3-bucket-queueconfiguration-event
	//
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type.
	//
	// FIFO queues are not allowed when enabling an SQS queue as the event notification destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-queueconfiguration.html#cfn-s3-bucket-queueconfiguration-queue
	//
	Queue *string `field:"required" json:"queue" yaml:"queue"`
	// The filtering rules that determine which objects trigger notifications.
	//
	// For example, you can create a filter so that Amazon S3 sends notifications only when image files with a `.jpg` extension are added to the bucket. For more information, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/user-guide/notification-how-to-filtering.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-queueconfiguration.html#cfn-s3-bucket-queueconfiguration-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

