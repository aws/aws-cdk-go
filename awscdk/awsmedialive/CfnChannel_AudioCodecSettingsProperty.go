package awsmedialive


// The configuration of the audio codec in the audio output.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioCodecSettingsProperty := &AudioCodecSettingsProperty{
//   	AacSettings: &AacSettingsProperty{
//   		Bitrate: jsii.Number(123),
//   		CodingMode: jsii.String("codingMode"),
//   		InputType: jsii.String("inputType"),
//   		Profile: jsii.String("profile"),
//   		RateControlMode: jsii.String("rateControlMode"),
//   		RawFormat: jsii.String("rawFormat"),
//   		SampleRate: jsii.Number(123),
//   		Spec: jsii.String("spec"),
//   		VbrQuality: jsii.String("vbrQuality"),
//   	},
//   	Ac3Settings: &Ac3SettingsProperty{
//   		Bitrate: jsii.Number(123),
//   		BitstreamMode: jsii.String("bitstreamMode"),
//   		CodingMode: jsii.String("codingMode"),
//   		Dialnorm: jsii.Number(123),
//   		DrcProfile: jsii.String("drcProfile"),
//   		LfeFilter: jsii.String("lfeFilter"),
//   		MetadataControl: jsii.String("metadataControl"),
//   	},
//   	Eac3AtmosSettings: &Eac3AtmosSettingsProperty{
//   		Bitrate: jsii.Number(123),
//   		CodingMode: jsii.String("codingMode"),
//   		Dialnorm: jsii.Number(123),
//   		DrcLine: jsii.String("drcLine"),
//   		DrcRf: jsii.String("drcRf"),
//   		HeightTrim: jsii.Number(123),
//   		SurroundTrim: jsii.Number(123),
//   	},
//   	Eac3Settings: &Eac3SettingsProperty{
//   		AttenuationControl: jsii.String("attenuationControl"),
//   		Bitrate: jsii.Number(123),
//   		BitstreamMode: jsii.String("bitstreamMode"),
//   		CodingMode: jsii.String("codingMode"),
//   		DcFilter: jsii.String("dcFilter"),
//   		Dialnorm: jsii.Number(123),
//   		DrcLine: jsii.String("drcLine"),
//   		DrcRf: jsii.String("drcRf"),
//   		LfeControl: jsii.String("lfeControl"),
//   		LfeFilter: jsii.String("lfeFilter"),
//   		LoRoCenterMixLevel: jsii.Number(123),
//   		LoRoSurroundMixLevel: jsii.Number(123),
//   		LtRtCenterMixLevel: jsii.Number(123),
//   		LtRtSurroundMixLevel: jsii.Number(123),
//   		MetadataControl: jsii.String("metadataControl"),
//   		PassthroughControl: jsii.String("passthroughControl"),
//   		PhaseControl: jsii.String("phaseControl"),
//   		StereoDownmix: jsii.String("stereoDownmix"),
//   		SurroundExMode: jsii.String("surroundExMode"),
//   		SurroundMode: jsii.String("surroundMode"),
//   	},
//   	Mp2Settings: &Mp2SettingsProperty{
//   		Bitrate: jsii.Number(123),
//   		CodingMode: jsii.String("codingMode"),
//   		SampleRate: jsii.Number(123),
//   	},
//   	PassThroughSettings: &PassThroughSettingsProperty{
//   	},
//   	WavSettings: &WavSettingsProperty{
//   		BitDepth: jsii.Number(123),
//   		CodingMode: jsii.String("codingMode"),
//   		SampleRate: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html
//
type CfnChannel_AudioCodecSettingsProperty struct {
	// The setup of the AAC audio codec in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-aacsettings
	//
	AacSettings interface{} `field:"optional" json:"aacSettings" yaml:"aacSettings"`
	// The setup of an AC3 audio codec in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-ac3settings
	//
	Ac3Settings interface{} `field:"optional" json:"ac3Settings" yaml:"ac3Settings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-eac3atmossettings
	//
	Eac3AtmosSettings interface{} `field:"optional" json:"eac3AtmosSettings" yaml:"eac3AtmosSettings"`
	// The setup of an EAC3 audio codec in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-eac3settings
	//
	Eac3Settings interface{} `field:"optional" json:"eac3Settings" yaml:"eac3Settings"`
	// The setup of an MP2 audio codec in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-mp2settings
	//
	Mp2Settings interface{} `field:"optional" json:"mp2Settings" yaml:"mp2Settings"`
	// The setup to pass through the Dolby audio codec to the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-passthroughsettings
	//
	PassThroughSettings interface{} `field:"optional" json:"passThroughSettings" yaml:"passThroughSettings"`
	// Settings for audio encoded with the WAV codec.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-wavsettings
	//
	WavSettings interface{} `field:"optional" json:"wavSettings" yaml:"wavSettings"`
}

