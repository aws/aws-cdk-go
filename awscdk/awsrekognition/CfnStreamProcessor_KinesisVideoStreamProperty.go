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
//   kinesisVideoStreamProperty := &KinesisVideoStreamProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-kinesisvideostream.html
//
type CfnStreamProcessor_KinesisVideoStreamProperty struct {
	// ARN of the Kinesis video stream stream that streams the source video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rekognition-streamprocessor-kinesisvideostream.html#cfn-rekognition-streamprocessor-kinesisvideostream-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

