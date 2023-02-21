package awsmedialive


// Settings to configure the destination of an Archive output.
//
// The parent of this entity is ArchiveGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveCdnSettingsProperty := &archiveCdnSettingsProperty{
//   	archiveS3Settings: &archiveS3SettingsProperty{
//   		cannedAcl: jsii.String("cannedAcl"),
//   	},
//   }
//
type CfnChannel_ArchiveCdnSettingsProperty struct {
	// Sets up Amazon S3 as the destination for this Archive output.
	ArchiveS3Settings interface{} `field:"optional" json:"archiveS3Settings" yaml:"archiveS3Settings"`
}

