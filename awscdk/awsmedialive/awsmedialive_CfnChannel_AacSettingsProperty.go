package awsmedialive


// The settings for an AAC audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aacSettingsProperty := &aacSettingsProperty{
//   	bitrate: jsii.Number(123),
//   	codingMode: jsii.String("codingMode"),
//   	inputType: jsii.String("inputType"),
//   	profile: jsii.String("profile"),
//   	rateControlMode: jsii.String("rateControlMode"),
//   	rawFormat: jsii.String("rawFormat"),
//   	sampleRate: jsii.Number(123),
//   	spec: jsii.String("spec"),
//   	vbrQuality: jsii.String("vbrQuality"),
//   }
//
type CfnChannel_AacSettingsProperty struct {
	// The average bitrate in bits/second.
	//
	// Valid values depend on the rate control mode and profile.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Mono, stereo, or 5.1 channel layout. Valid values depend on the rate control mode and profile. The adReceiverMix setting receives a stereo description plus control track, and emits a mono AAC encode of the description track, with control data emitted in the PES header as per ETSI TS 101 154 Annex E.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Set to broadcasterMixedAd when the input contains pre-mixed main audio + AD (narration) as a stereo pair.
	//
	// The Audio Type field (audioType) will be set to 3, which signals to downstream systems that this stream contains broadcaster mixed AD. Note that the input received by the encoder must contain pre-mixed audio; MediaLive does not perform the mixing. The values in audioTypeControl and audioType (in AudioDescription) are ignored when set to broadcasterMixedAd. Leave this set to normal when the input does not contain pre-mixed audio + AD.
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// The AAC profile.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The rate control mode.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the LATM/LOAS AAC output for raw containers.
	RawFormat *string `field:"optional" json:"rawFormat" yaml:"rawFormat"`
	// The sample rate in Hz.
	//
	// Valid values depend on the rate control mode and profile.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
	// Uses MPEG-2 AAC audio instead of MPEG-4 AAC audio for raw or MPEG-2 Transport Stream containers.
	Spec *string `field:"optional" json:"spec" yaml:"spec"`
	// The VBR quality level.
	//
	// This is used only if rateControlMode is VBR.
	VbrQuality *string `field:"optional" json:"vbrQuality" yaml:"vbrQuality"`
}

