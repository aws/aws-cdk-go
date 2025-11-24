package mixinsawsmedialive


// Sets up Amazon S3 as the destination for this Frame Capture output.
//
// The parent of this entity is FrameCaptureCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   frameCaptureS3SettingsProperty := &FrameCaptureS3SettingsProperty{
//   	CannedAcl: jsii.String("cannedAcl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecaptures3settings.html
//
type CfnChannelPropsMixin_FrameCaptureS3SettingsProperty struct {
	// Specify the canned ACL to apply to each S3 request.
	//
	// Defaults to none.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecaptures3settings.html#cfn-medialive-channel-framecaptures3settings-cannedacl
	//
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

