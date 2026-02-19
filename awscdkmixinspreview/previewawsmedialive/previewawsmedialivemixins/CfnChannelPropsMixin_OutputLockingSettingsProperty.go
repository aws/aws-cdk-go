package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputLockingSettingsProperty := &OutputLockingSettingsProperty{
//   	DisabledLockingSettings: &DisabledLockingSettingsProperty{
//   		CustomEpoch: jsii.String("customEpoch"),
//   	},
//   	EpochLockingSettings: &EpochLockingSettingsProperty{
//   		CustomEpoch: jsii.String("customEpoch"),
//   		JamSyncTime: jsii.String("jamSyncTime"),
//   	},
//   	PipelineLockingSettings: &PipelineLockingSettingsProperty{
//   		CustomEpoch: jsii.String("customEpoch"),
//   		PipelineLockingMethod: jsii.String("pipelineLockingMethod"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputlockingsettings.html
//
type CfnChannelPropsMixin_OutputLockingSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputlockingsettings.html#cfn-medialive-channel-outputlockingsettings-disabledlockingsettings
	//
	DisabledLockingSettings interface{} `field:"optional" json:"disabledLockingSettings" yaml:"disabledLockingSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputlockingsettings.html#cfn-medialive-channel-outputlockingsettings-epochlockingsettings
	//
	EpochLockingSettings interface{} `field:"optional" json:"epochLockingSettings" yaml:"epochLockingSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputlockingsettings.html#cfn-medialive-channel-outputlockingsettings-pipelinelockingsettings
	//
	PipelineLockingSettings interface{} `field:"optional" json:"pipelineLockingSettings" yaml:"pipelineLockingSettings"`
}

