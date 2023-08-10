package awsnimblestudio


// <p>Configures how streaming sessions are backed up when launched from this launch             profile.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationSessionBackupProperty := &StreamConfigurationSessionBackupProperty{
//   	MaxBackupsToRetain: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionbackup.html
//
type CfnLaunchProfile_StreamConfigurationSessionBackupProperty struct {
	// <p>The maximum number of backups that each streaming session created from this launch             profile can have.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionbackup.html#cfn-nimblestudio-launchprofile-streamconfigurationsessionbackup-maxbackupstoretain
	//
	MaxBackupsToRetain *float64 `field:"optional" json:"maxBackupsToRetain" yaml:"maxBackupsToRetain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionbackup.html#cfn-nimblestudio-launchprofile-streamconfigurationsessionbackup-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

