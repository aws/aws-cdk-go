package awsmedialive


// Sets up Amazon S3 as the destination for this Frame Capture output.
//
// The parent of this entity is FrameCaptureCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureS3SettingsProperty := &frameCaptureS3SettingsProperty{
//   	cannedAcl: jsii.String("cannedAcl"),
//   }
//
type CfnChannel_FrameCaptureS3SettingsProperty struct {
	// Specify the canned ACL to apply to each S3 request.
	//
	// Defaults to none.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

