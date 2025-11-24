package mixinsawsmedialive


// Selects a specific PID from within a video source.
//
// The parent of this entity is VideoSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   videoSelectorPidProperty := &VideoSelectorPidProperty{
//   	Pid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorpid.html
//
type CfnChannelPropsMixin_VideoSelectorPidProperty struct {
	// Selects a specific PID from within a video source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorpid.html#cfn-medialive-channel-videoselectorpid-pid
	//
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

