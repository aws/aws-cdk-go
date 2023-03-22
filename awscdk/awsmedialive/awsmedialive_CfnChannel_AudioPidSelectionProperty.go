package awsmedialive


// Used to extract audio by The PID.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioPidSelectionProperty := &audioPidSelectionProperty{
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_AudioPidSelectionProperty struct {
	// Select the audio by this PID.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

