package awskinesisanalyticsv2


// For a SQL-based Kinesis Data Analytics application, when configuring application output, identifies a Kinesis Data Firehose delivery stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN) of the delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseOutputProperty := &kinesisFirehoseOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationOutput_KinesisFirehoseOutputProperty struct {
	// The ARN of the destination delivery stream to write to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

