package awsnimblestudio


// Configures how streaming sessions are backed up when launched from this launch profile.
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
type CfnLaunchProfile_StreamConfigurationSessionBackupProperty struct {
	// The maximum number of backups that each streaming session created from this launch profile can have.
	MaxBackupsToRetain *float64 `field:"optional" json:"maxBackupsToRetain" yaml:"maxBackupsToRetain"`
	// Specifies how artists sessions are backed up.
	//
	// Configures backups for streaming sessions launched with this launch profile. The default value is `DEACTIVATED` , which means that backups are deactivated. To allow backups, set this value to `AUTOMATIC` .
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

