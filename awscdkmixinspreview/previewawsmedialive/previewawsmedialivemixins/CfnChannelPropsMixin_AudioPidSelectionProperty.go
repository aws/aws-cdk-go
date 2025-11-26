package previewawsmedialivemixins


// Used to extract audio by The PID.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioPidSelectionProperty := &AudioPidSelectionProperty{
//   	Pid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopidselection.html
//
type CfnChannelPropsMixin_AudioPidSelectionProperty struct {
	// Select the audio by this PID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopidselection.html#cfn-medialive-channel-audiopidselection-pid
	//
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

