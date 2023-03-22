package awskinesisanalyticsv2


// When you configure a SQL-based Kinesis Data Analytics application's output, identifies a Kinesis data stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsOutputProperty := &kinesisStreamsOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationOutput_KinesisStreamsOutputProperty struct {
	// The ARN of the destination Kinesis data stream to write to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

