package awskinesisanalytics


// Identifies an Amazon Kinesis stream as the streaming source.
//
// You provide the stream's Amazon Resource Name (ARN) and an IAM role ARN that enables Amazon Kinesis Analytics to access the stream on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsInputProperty := &kinesisStreamsInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplication_KinesisStreamsInputProperty struct {
	// ARN of the input Amazon Kinesis stream to read.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the stream on your behalf.
	//
	// You need to grant the necessary permissions to this role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

