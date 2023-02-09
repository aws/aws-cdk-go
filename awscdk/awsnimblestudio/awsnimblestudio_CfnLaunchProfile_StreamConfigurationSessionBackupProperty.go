package awsnimblestudio


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationSessionBackupProperty := &streamConfigurationSessionBackupProperty{
//   	maxBackupsToRetain: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   }
//
type CfnLaunchProfile_StreamConfigurationSessionBackupProperty struct {
	// `CfnLaunchProfile.StreamConfigurationSessionBackupProperty.MaxBackupsToRetain`.
	MaxBackupsToRetain *float64 `field:"optional" json:"maxBackupsToRetain" yaml:"maxBackupsToRetain"`
	// `CfnLaunchProfile.StreamConfigurationSessionBackupProperty.Mode`.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

