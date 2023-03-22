package awspinpointemail


// An object that defines an Amazon Kinesis Data Firehose destination for email events.
//
// You can use Amazon Kinesis Data Firehose to stream data to other services, such as Amazon S3 and Amazon Redshift.
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
	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that Amazon Pinpoint sends email events to.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The Amazon Resource Name (ARN) of the IAM role that Amazon Pinpoint uses when sending email events to the Amazon Kinesis Data Firehose stream.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

