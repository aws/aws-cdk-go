package previewawsmediaconnectmixins


// Specifies the configuration for audio stream metrics monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioMonitoringSettingProperty := &AudioMonitoringSettingProperty{
//   	SilentAudio: &SilentAudioProperty{
//   		State: jsii.String("state"),
//   		ThresholdSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-audiomonitoringsetting.html
//
type CfnFlowPropsMixin_AudioMonitoringSettingProperty struct {
	// Detects periods of silence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-audiomonitoringsetting.html#cfn-mediaconnect-flow-audiomonitoringsetting-silentaudio
	//
	SilentAudio interface{} `field:"optional" json:"silentAudio" yaml:"silentAudio"`
}

