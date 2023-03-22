package awskinesisanalytics


// When configuring application output, identifies an Amazon Kinesis stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN) and also an IAM role ARN that Amazon Kinesis Analytics can use to write to the stream on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsOutputProperty := &kinesisStreamsOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplicationOutput_KinesisStreamsOutputProperty struct {
	// ARN of the destination Amazon Kinesis stream to write to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the destination stream on your behalf.
	//
	// You need to grant the necessary permissions to this role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

