package awsrekognition


// Amazon Rekognition Video Stream Processor take as input a Kinesis video stream (Input) and a Kinesis data stream (Output).
//
// This is the Amazon Kinesis Data Streams instance to which the Amazon Rekognition stream processor streams the analysis results. This must be created within the constraints specified at [KinesisDataStream](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_KinesisDataStream) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisDataStreamProperty := &kinesisDataStreamProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnStreamProcessor_KinesisDataStreamProperty struct {
	// ARN of the output Amazon Kinesis Data Streams stream.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

