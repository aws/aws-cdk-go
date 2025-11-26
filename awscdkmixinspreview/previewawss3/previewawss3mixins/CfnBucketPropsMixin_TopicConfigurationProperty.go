package previewawss3mixins


// A container for specifying the configuration for publication of messages to an Amazon Simple Notification Service (Amazon SNS) topic when Amazon S3 detects specified events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicConfigurationProperty := &TopicConfigurationProperty{
//   	Event: jsii.String("event"),
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
//   	Topic: jsii.String("topic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-topicconfiguration.html
//
type CfnBucketPropsMixin_TopicConfigurationProperty struct {
	// The Amazon S3 bucket event about which to send notifications.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-topicconfiguration.html#cfn-s3-bucket-topicconfiguration-event
	//
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The filtering rules that determine for which objects to send notifications.
	//
	// For example, you can create a filter so that Amazon S3 sends notifications only when image files with a `.jpg` extension are added to the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-topicconfiguration.html#cfn-s3-bucket-topicconfiguration-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it detects events of the specified type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-topicconfiguration.html#cfn-s3-bucket-topicconfiguration-topic
	//
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

