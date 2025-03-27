package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   epochLockingSettingsProperty := &EpochLockingSettingsProperty{
//   	CustomEpoch: jsii.String("customEpoch"),
//   	JamSyncTime: jsii.String("jamSyncTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-epochlockingsettings.html
//
type CfnChannel_EpochLockingSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-epochlockingsettings.html#cfn-medialive-channel-epochlockingsettings-customepoch
	//
	CustomEpoch *string `field:"optional" json:"customEpoch" yaml:"customEpoch"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-epochlockingsettings.html#cfn-medialive-channel-epochlockingsettings-jamsynctime
	//
	JamSyncTime *string `field:"optional" json:"jamSyncTime" yaml:"jamSyncTime"`
}

