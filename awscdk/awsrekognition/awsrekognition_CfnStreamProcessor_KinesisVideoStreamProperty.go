package awsrekognition


// The Kinesis video stream that provides the source of the streaming video for an Amazon Rekognition Video stream processor.
//
// For more information, see [KinesisVideoStream](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_KinesisVideoStream) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisVideoStreamProperty := &kinesisVideoStreamProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnStreamProcessor_KinesisVideoStreamProperty struct {
	// ARN of the Kinesis video stream stream that streams the source video.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

