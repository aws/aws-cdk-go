package awsmedialive


// The settings for a frame capture output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureGroupSettingsProperty := &frameCaptureGroupSettingsProperty{
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	frameCaptureCdnSettings: &frameCaptureCdnSettingsProperty{
//   		frameCaptureS3Settings: &frameCaptureS3SettingsProperty{
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   	},
//   }
//
type CfnChannel_FrameCaptureGroupSettingsProperty struct {
	// The destination for the frame capture files.
	//
	// The destination is either the URI for an Amazon S3 bucket and object, plus a file name prefix (for example, s3ssl://sportsDelivery/highlights/20180820/curling_) or the URI for a MediaStore container, plus a file name prefix (for example, mediastoressl://sportsDelivery/20180820/curling_). The final file names consist of the prefix from the destination field (for example, "curling_") + name modifier + the counter (5 digits, starting from 00001) + extension (which is always .jpg). For example, curlingLow.00001.jpg.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Settings to configure the destination of a Frame Capture output.
	FrameCaptureCdnSettings interface{} `field:"optional" json:"frameCaptureCdnSettings" yaml:"frameCaptureCdnSettings"`
}

