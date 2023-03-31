package awskinesisanalytics


// Identifies an Amazon Kinesis Firehose delivery stream as the streaming source.
//
// You provide the delivery stream's Amazon Resource Name (ARN) and an IAM role ARN that enables Amazon Kinesis Analytics to access the stream on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseInputProperty := &kinesisFirehoseInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplication_KinesisFirehoseInputProperty struct {
	// ARN of the input delivery stream.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the stream on your behalf.
	//
	// You need to make sure that the role has the necessary permissions to access the stream.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

