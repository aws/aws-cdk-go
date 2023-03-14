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
//   videoSelectorSettingsProperty := &videoSelectorSettingsProperty{
//   	videoSelectorPid: &videoSelectorPidProperty{
//   		pid: jsii.Number(123),
//   	},
//   	videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   		programId: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_VideoSelectorSettingsProperty struct {
	// Used to extract video by PID.
	VideoSelectorPid interface{} `field:"optional" json:"videoSelectorPid" yaml:"videoSelectorPid"`
	// Used to extract video by program ID.
	VideoSelectorProgramId interface{} `field:"optional" json:"videoSelectorProgramId" yaml:"videoSelectorProgramId"`
}

