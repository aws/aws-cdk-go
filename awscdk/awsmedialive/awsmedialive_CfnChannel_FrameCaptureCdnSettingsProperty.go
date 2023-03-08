package awsmedialive


// Settings to configure the destination of a Frame Capture output.
//
// The parent of this entity is FrameCaptureGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureCdnSettingsProperty := &frameCaptureCdnSettingsProperty{
//   	frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   		cannedAcl: jsii.String("cannedAcl"),
//   	},
//   }
//
type CfnChannel_FrameCaptureCdnSettingsProperty struct {
	// Sets up Amazon S3 as the destination for this Frame Capture output.
	FrameCaptureS3Settings interface{} `field:"optional" json:"frameCaptureS3Settings" yaml:"frameCaptureS3Settings"`
}

