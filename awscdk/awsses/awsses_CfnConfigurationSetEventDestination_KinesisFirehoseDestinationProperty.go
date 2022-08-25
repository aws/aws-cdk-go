package awsses


// Contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
//
// Event destinations, such as Amazon Kinesis Firehose, are associated with configuration sets, which enable you to publish email sending events. For information about using configuration sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseDestinationProperty := &kinesisFirehoseDestinationProperty{
//   	deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   }
//
type CfnConfigurationSetEventDestination_KinesisFirehoseDestinationProperty struct {
	// The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

