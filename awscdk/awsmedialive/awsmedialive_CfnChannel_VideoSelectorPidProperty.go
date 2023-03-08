package awsmedialive


// Selects a specific PID from within a video source.
//
// The parent of this entity is VideoSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorPidProperty := &videoSelectorPidProperty{
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_VideoSelectorPidProperty struct {
	// Selects a specific PID from within a video source.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

