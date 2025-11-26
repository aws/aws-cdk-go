package previewawsmedialivemixins


// Sets up Amazon S3 as the destination for this Archive output.
//
// The parent of this entity is ArchiveCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   archiveS3SettingsProperty := &ArchiveS3SettingsProperty{
//   	CannedAcl: jsii.String("cannedAcl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archives3settings.html
//
type CfnChannelPropsMixin_ArchiveS3SettingsProperty struct {
	// Specify the canned ACL to apply to each S3 request.
	//
	// Defaults to none.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archives3settings.html#cfn-medialive-channel-archives3settings-cannedacl
	//
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

