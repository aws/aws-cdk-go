package awskinesisanalyticsv2


// Identifies a Kinesis data stream as the streaming source.
//
// You provide the stream's Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsInputProperty := &kinesisStreamsInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplication_KinesisStreamsInputProperty struct {
	// The ARN of the input Kinesis data stream to read.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

