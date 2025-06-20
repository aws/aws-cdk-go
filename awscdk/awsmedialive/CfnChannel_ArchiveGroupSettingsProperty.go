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
//   archiveGroupSettingsProperty := &ArchiveGroupSettingsProperty{
//   	ArchiveCdnSettings: &ArchiveCdnSettingsProperty{
//   		ArchiveS3Settings: &ArchiveS3SettingsProperty{
//   			CannedAcl: jsii.String("cannedAcl"),
//   		},
//   	},
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	RolloverInterval: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivegroupsettings.html
//
type CfnChannel_ArchiveGroupSettingsProperty struct {
	// Settings to configure the destination of an Archive output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivegroupsettings.html#cfn-medialive-channel-archivegroupsettings-archivecdnsettings
	//
	ArchiveCdnSettings interface{} `field:"optional" json:"archiveCdnSettings" yaml:"archiveCdnSettings"`
	// A directory and base file name where archive files should be written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivegroupsettings.html#cfn-medialive-channel-archivegroupsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The number of seconds to write to an archive file before closing and starting a new one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivegroupsettings.html#cfn-medialive-channel-archivegroupsettings-rolloverinterval
	//
	RolloverInterval *float64 `field:"optional" json:"rolloverInterval" yaml:"rolloverInterval"`
}

