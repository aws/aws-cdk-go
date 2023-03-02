package awsmedialive


// The settings for an archive output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveGroupSettingsProperty := &archiveGroupSettingsProperty{
//   	archiveCdnSettings: &archiveCdnSettingsProperty{
//   		archiveS3Settings: &archiveS3SettingsProperty{
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   	},
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	rolloverInterval: jsii.Number(123),
//   }
//
type CfnChannel_ArchiveGroupSettingsProperty struct {
	// Settings to configure the destination of an Archive output.
	ArchiveCdnSettings interface{} `field:"optional" json:"archiveCdnSettings" yaml:"archiveCdnSettings"`
	// A directory and base file name where archive files should be written.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The number of seconds to write to an archive file before closing and starting a new one.
	RolloverInterval *float64 `field:"optional" json:"rolloverInterval" yaml:"rolloverInterval"`
}

