package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexM2tsSettingsProperty := &MultiplexM2tsSettingsProperty{
//   	AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   	Arib: jsii.String("arib"),
//   	AudioBufferModel: jsii.String("audioBufferModel"),
//   	AudioFramesPerPes: jsii.Number(123),
//   	AudioStreamType: jsii.String("audioStreamType"),
//   	CcDescriptor: jsii.String("ccDescriptor"),
//   	Ebif: jsii.String("ebif"),
//   	EsRateInPes: jsii.String("esRateInPes"),
//   	Klv: jsii.String("klv"),
//   	NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	PcrControl: jsii.String("pcrControl"),
//   	PcrPeriod: jsii.Number(123),
//   	Scte35Control: jsii.String("scte35Control"),
//   	Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html
//
type CfnChannelPropsMixin_MultiplexM2tsSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-absentinputaudiobehavior
	//
	AbsentInputAudioBehavior *string `field:"optional" json:"absentInputAudioBehavior" yaml:"absentInputAudioBehavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-arib
	//
	Arib *string `field:"optional" json:"arib" yaml:"arib"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-audiobuffermodel
	//
	AudioBufferModel *string `field:"optional" json:"audioBufferModel" yaml:"audioBufferModel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-audioframesperpes
	//
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-audiostreamtype
	//
	AudioStreamType *string `field:"optional" json:"audioStreamType" yaml:"audioStreamType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-ccdescriptor
	//
	CcDescriptor *string `field:"optional" json:"ccDescriptor" yaml:"ccDescriptor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-ebif
	//
	Ebif *string `field:"optional" json:"ebif" yaml:"ebif"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-esrateinpes
	//
	EsRateInPes *string `field:"optional" json:"esRateInPes" yaml:"esRateInPes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-klv
	//
	Klv *string `field:"optional" json:"klv" yaml:"klv"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-nielsenid3behavior
	//
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-pcrcontrol
	//
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-pcrperiod
	//
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-scte35control
	//
	Scte35Control *string `field:"optional" json:"scte35Control" yaml:"scte35Control"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexm2tssettings.html#cfn-medialive-channel-multiplexm2tssettings-scte35prerollpullupmilliseconds
	//
	Scte35PrerollPullupMilliseconds *float64 `field:"optional" json:"scte35PrerollPullupMilliseconds" yaml:"scte35PrerollPullupMilliseconds"`
}

