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
//   audioCodecSettingsProperty := &audioCodecSettingsProperty{
//   	aacSettings: &aacSettingsProperty{
//   		bitrate: jsii.Number(123),
//   		codingMode: jsii.String("codingMode"),
//   		inputType: jsii.String("inputType"),
//   		profile: jsii.String("profile"),
//   		rateControlMode: jsii.String("rateControlMode"),
//   		rawFormat: jsii.String("rawFormat"),
//   		sampleRate: jsii.Number(123),
//   		spec: jsii.String("spec"),
//   		vbrQuality: jsii.String("vbrQuality"),
//   	},
//   	ac3Settings: &ac3SettingsProperty{
//   		bitrate: jsii.Number(123),
//   		bitstreamMode: jsii.String("bitstreamMode"),
//   		codingMode: jsii.String("codingMode"),
//   		dialnorm: jsii.Number(123),
//   		drcProfile: jsii.String("drcProfile"),
//   		lfeFilter: jsii.String("lfeFilter"),
//   		metadataControl: jsii.String("metadataControl"),
//   	},
//   	eac3Settings: &eac3SettingsProperty{
//   		attenuationControl: jsii.String("attenuationControl"),
//   		bitrate: jsii.Number(123),
//   		bitstreamMode: jsii.String("bitstreamMode"),
//   		codingMode: jsii.String("codingMode"),
//   		dcFilter: jsii.String("dcFilter"),
//   		dialnorm: jsii.Number(123),
//   		drcLine: jsii.String("drcLine"),
//   		drcRf: jsii.String("drcRf"),
//   		lfeControl: jsii.String("lfeControl"),
//   		lfeFilter: jsii.String("lfeFilter"),
//   		loRoCenterMixLevel: jsii.Number(123),
//   		loRoSurroundMixLevel: jsii.Number(123),
//   		ltRtCenterMixLevel: jsii.Number(123),
//   		ltRtSurroundMixLevel: jsii.Number(123),
//   		metadataControl: jsii.String("metadataControl"),
//   		passthroughControl: jsii.String("passthroughControl"),
//   		phaseControl: jsii.String("phaseControl"),
//   		stereoDownmix: jsii.String("stereoDownmix"),
//   		surroundExMode: jsii.String("surroundExMode"),
//   		surroundMode: jsii.String("surroundMode"),
//   	},
//   	mp2Settings: &mp2SettingsProperty{
//   		bitrate: jsii.Number(123),
//   		codingMode: jsii.String("codingMode"),
//   		sampleRate: jsii.Number(123),
//   	},
//   	passThroughSettings: &passThroughSettingsProperty{
//   	},
//   	wavSettings: &wavSettingsProperty{
//   		bitDepth: jsii.Number(123),
//   		codingMode: jsii.String("codingMode"),
//   		sampleRate: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_AudioCodecSettingsProperty struct {
	// The setup of the AAC audio codec in the output.
	AacSettings interface{} `field:"optional" json:"aacSettings" yaml:"aacSettings"`
	// The setup of an AC3 audio codec in the output.
	Ac3Settings interface{} `field:"optional" json:"ac3Settings" yaml:"ac3Settings"`
	// The setup of an EAC3 audio codec in the output.
	Eac3Settings interface{} `field:"optional" json:"eac3Settings" yaml:"eac3Settings"`
	// The setup of an MP2 audio codec in the output.
	Mp2Settings interface{} `field:"optional" json:"mp2Settings" yaml:"mp2Settings"`
	// The setup to pass through the Dolby audio codec to the output.
	PassThroughSettings interface{} `field:"optional" json:"passThroughSettings" yaml:"passThroughSettings"`
	// Settings for audio encoded with the WAV codec.
	WavSettings interface{} `field:"optional" json:"wavSettings" yaml:"wavSettings"`
}

