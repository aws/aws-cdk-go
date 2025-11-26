package previewawsmedialivemixins


// Settings to configure the destination of a Frame Capture output.
//
// The parent of this entity is FrameCaptureGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   frameCaptureCdnSettingsProperty := &FrameCaptureCdnSettingsProperty{
//   	FrameCaptureS3Settings: &FrameCaptureS3SettingsProperty{
//   		CannedAcl: jsii.String("cannedAcl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturecdnsettings.html
//
type CfnChannelPropsMixin_FrameCaptureCdnSettingsProperty struct {
	// Sets up Amazon S3 as the destination for this Frame Capture output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturecdnsettings.html#cfn-medialive-channel-framecapturecdnsettings-framecaptures3settings
	//
	FrameCaptureS3Settings interface{} `field:"optional" json:"frameCaptureS3Settings" yaml:"frameCaptureS3Settings"`
}

