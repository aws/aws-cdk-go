package previewawspinpointemailmixins


// An object that defines an Amazon Kinesis Data Firehose destination for email events.
//
// You can use Amazon Kinesis Data Firehose to stream data to other services, such as Amazon S3 and Amazon Redshift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisFirehoseDestinationProperty := &KinesisFirehoseDestinationProperty{
//   	DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-kinesisfirehosedestination.html
//
type CfnConfigurationSetEventDestinationPropsMixin_KinesisFirehoseDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that Amazon Pinpoint sends email events to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-kinesisfirehosedestination.html#cfn-pinpointemail-configurationseteventdestination-kinesisfirehosedestination-deliverystreamarn
	//
	DeliveryStreamArn *string `field:"optional" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The Amazon Resource Name (ARN) of the IAM role that Amazon Pinpoint uses when sending email events to the Amazon Kinesis Data Firehose stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-kinesisfirehosedestination.html#cfn-pinpointemail-configurationseteventdestination-kinesisfirehosedestination-iamrolearn
	//
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
}

