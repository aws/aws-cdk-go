package awskinesisanalytics


// For a SQL-based Kinesis Data Analytics application, identifies a Kinesis Data Firehose delivery stream as the streaming source.
//
// You provide the delivery stream's Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseInputProperty := &kinesisFirehoseInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationV2_KinesisFirehoseInputProperty struct {
	// The Amazon Resource Name (ARN) of the delivery stream.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

