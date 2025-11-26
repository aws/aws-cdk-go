package previewawsmediaconnectmixins


// Configures settings for the `SilentAudio` metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   silentAudioProperty := &SilentAudioProperty{
//   	State: jsii.String("state"),
//   	ThresholdSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-silentaudio.html
//
type CfnFlowPropsMixin_SilentAudioProperty struct {
	// Indicates whether the `SilentAudio` metric is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-silentaudio.html#cfn-mediaconnect-flow-silentaudio-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// Specifies the number of consecutive seconds of silence that triggers an event or alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-silentaudio.html#cfn-mediaconnect-flow-silentaudio-thresholdseconds
	//
	ThresholdSeconds *float64 `field:"optional" json:"thresholdSeconds" yaml:"thresholdSeconds"`
}

