package mixinsawsmedialive


// The configuration settings that apply to the entire channel.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var pipelineLockingSettings interface{}
//
//   globalConfigurationProperty := &GlobalConfigurationProperty{
//   	InitialAudioGain: jsii.Number(123),
//   	InputEndAction: jsii.String("inputEndAction"),
//   	InputLossBehavior: &InputLossBehaviorProperty{
//   		BlackFrameMsec: jsii.Number(123),
//   		InputLossImageColor: jsii.String("inputLossImageColor"),
//   		InputLossImageSlate: &InputLocationProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			Uri: jsii.String("uri"),
//   			Username: jsii.String("username"),
//   		},
//   		InputLossImageType: jsii.String("inputLossImageType"),
//   		RepeatFrameMsec: jsii.Number(123),
//   	},
//   	OutputLockingMode: jsii.String("outputLockingMode"),
//   	OutputLockingSettings: &OutputLockingSettingsProperty{
//   		EpochLockingSettings: &EpochLockingSettingsProperty{
//   			CustomEpoch: jsii.String("customEpoch"),
//   			JamSyncTime: jsii.String("jamSyncTime"),
//   		},
//   		PipelineLockingSettings: pipelineLockingSettings,
//   	},
//   	OutputTimingSource: jsii.String("outputTimingSource"),
//   	SupportLowFramerateInputs: jsii.String("supportLowFramerateInputs"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html
//
type CfnChannelPropsMixin_GlobalConfigurationProperty struct {
	// The value to set the initial audio gain for the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html#cfn-medialive-channel-globalconfiguration-initialaudiogain
	//
	InitialAudioGain *float64 `field:"optional" json:"initialAudioGain" yaml:"initialAudioGain"`
	// Indicates the action to take when the current input completes (for example, end-of-file).
	//
	// When switchAndLoopInputs is configured, MediaLive restarts at the beginning of the first input. When "none" is configured, MediaLive transcodes either black, a solid color, or a user-specified slate images per the "Input Loss Behavior" configuration until the next input switch occurs (which is controlled through the Channel Schedule API).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html#cfn-medialive-channel-globalconfiguration-inputendaction
	//
	InputEndAction *string `field:"optional" json:"inputEndAction" yaml:"inputEndAction"`
	// The settings for system actions when the input is lost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html#cfn-medialive-channel-globalconfiguration-inputlossbehavior
	//
	InputLossBehavior interface{} `field:"optional" json:"inputLossBehavior" yaml:"inputLossBehavior"`
	// Indicates how MediaLive pipelines are synchronized.
	//
	// PIPELINELOCKING - MediaLive attempts to synchronize the output of each pipeline to the other. EPOCHLOCKING - MediaLive attempts to synchronize the output of each pipeline to the Unix epoch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html#cfn-medialive-channel-globalconfiguration-outputlockingmode
	//
	OutputLockingMode *string `field:"optional" json:"outputLockingMode" yaml:"outputLockingMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html#cfn-medialive-channel-globalconfiguration-outputlockingsettings
	//
	OutputLockingSettings interface{} `field:"optional" json:"outputLockingSettings" yaml:"outputLockingSettings"`
	// Indicates whether the rate of frames emitted by the Live encoder should be paced by its system clock (which optionally might be locked to another source through NTP) or should be locked to the clock of the source that is providing the input stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html#cfn-medialive-channel-globalconfiguration-outputtimingsource
	//
	OutputTimingSource *string `field:"optional" json:"outputTimingSource" yaml:"outputTimingSource"`
	// Adjusts the video input buffer for streams with very low video frame rates.
	//
	// This is commonly set to enabled for music channels with less than one video frame per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-globalconfiguration.html#cfn-medialive-channel-globalconfiguration-supportlowframerateinputs
	//
	SupportLowFramerateInputs *string `field:"optional" json:"supportLowFramerateInputs" yaml:"supportLowFramerateInputs"`
}

