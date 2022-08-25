package awsmedialive


// The configuration settings that apply to the entire channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalConfigurationProperty := &globalConfigurationProperty{
//   	initialAudioGain: jsii.Number(123),
//   	inputEndAction: jsii.String("inputEndAction"),
//   	inputLossBehavior: &inputLossBehaviorProperty{
//   		blackFrameMsec: jsii.Number(123),
//   		inputLossImageColor: jsii.String("inputLossImageColor"),
//   		inputLossImageSlate: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		inputLossImageType: jsii.String("inputLossImageType"),
//   		repeatFrameMsec: jsii.Number(123),
//   	},
//   	outputLockingMode: jsii.String("outputLockingMode"),
//   	outputTimingSource: jsii.String("outputTimingSource"),
//   	supportLowFramerateInputs: jsii.String("supportLowFramerateInputs"),
//   }
//
type CfnChannel_GlobalConfigurationProperty struct {
	// The value to set the initial audio gain for the channel.
	InitialAudioGain *float64 `field:"optional" json:"initialAudioGain" yaml:"initialAudioGain"`
	// Indicates the action to take when the current input completes (for example, end-of-file).
	//
	// When switchAndLoopInputs is configured, MediaLive restarts at the beginning of the first input. When "none" is configured, MediaLive transcodes either black, a solid color, or a user-specified slate images per the "Input Loss Behavior" configuration until the next input switch occurs (which is controlled through the Channel Schedule API).
	InputEndAction *string `field:"optional" json:"inputEndAction" yaml:"inputEndAction"`
	// The settings for system actions when the input is lost.
	InputLossBehavior interface{} `field:"optional" json:"inputLossBehavior" yaml:"inputLossBehavior"`
	// Indicates how MediaLive pipelines are synchronized.
	//
	// PIPELINELOCKING - MediaLive attempts to synchronize the output of each pipeline to the other. EPOCHLOCKING - MediaLive attempts to synchronize the output of each pipeline to the Unix epoch.
	OutputLockingMode *string `field:"optional" json:"outputLockingMode" yaml:"outputLockingMode"`
	// Indicates whether the rate of frames emitted by the Live encoder should be paced by its system clock (which optionally might be locked to another source through NTP) or should be locked to the clock of the source that is providing the input stream.
	OutputTimingSource *string `field:"optional" json:"outputTimingSource" yaml:"outputTimingSource"`
	// Adjusts the video input buffer for streams with very low video frame rates.
	//
	// This is commonly set to enabled for music channels with less than one video frame per second.
	SupportLowFramerateInputs *string `field:"optional" json:"supportLowFramerateInputs" yaml:"supportLowFramerateInputs"`
}

