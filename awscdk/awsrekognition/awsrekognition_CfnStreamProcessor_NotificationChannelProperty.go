package awsrekognition


// The Amazon Simple Notification Service topic to which Amazon Rekognition publishes the object detection results and completion status of a video analysis operation.
//
// Amazon Rekognition publishes a notification the first time an object of interest or a person is detected in the video stream. Amazon Rekognition also publishes an an end-of-session notification with a summary when the stream processing session is complete. For more information, see [StreamProcessorNotificationChannel](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorNotificationChannel) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationChannelProperty := &notificationChannelProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnStreamProcessor_NotificationChannelProperty struct {
	// The ARN of the SNS topic that receives notifications.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

