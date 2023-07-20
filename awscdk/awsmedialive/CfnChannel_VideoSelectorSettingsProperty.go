package awsmedialive


// Information about the video to extract from the input.
//
// The parent of this entity is VideoSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorSettingsProperty := &VideoSelectorSettingsProperty{
//   	VideoSelectorPid: &VideoSelectorPidProperty{
//   		Pid: jsii.Number(123),
//   	},
//   	VideoSelectorProgramId: &VideoSelectorProgramIdProperty{
//   		ProgramId: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorsettings.html
//
type CfnChannel_VideoSelectorSettingsProperty struct {
	// Used to extract video by PID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorsettings.html#cfn-medialive-channel-videoselectorsettings-videoselectorpid
	//
	VideoSelectorPid interface{} `field:"optional" json:"videoSelectorPid" yaml:"videoSelectorPid"`
	// Used to extract video by program ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorsettings.html#cfn-medialive-channel-videoselectorsettings-videoselectorprogramid
	//
	VideoSelectorProgramId interface{} `field:"optional" json:"videoSelectorProgramId" yaml:"videoSelectorProgramId"`
}

