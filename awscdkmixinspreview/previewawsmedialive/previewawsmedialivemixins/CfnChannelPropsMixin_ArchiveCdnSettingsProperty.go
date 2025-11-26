package previewawsmedialivemixins


// Settings to configure the destination of an Archive output.
//
// The parent of this entity is ArchiveGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   archiveCdnSettingsProperty := &ArchiveCdnSettingsProperty{
//   	ArchiveS3Settings: &ArchiveS3SettingsProperty{
//   		CannedAcl: jsii.String("cannedAcl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecdnsettings.html
//
type CfnChannelPropsMixin_ArchiveCdnSettingsProperty struct {
	// Sets up Amazon S3 as the destination for this Archive output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecdnsettings.html#cfn-medialive-channel-archivecdnsettings-archives3settings
	//
	ArchiveS3Settings interface{} `field:"optional" json:"archiveS3Settings" yaml:"archiveS3Settings"`
}

