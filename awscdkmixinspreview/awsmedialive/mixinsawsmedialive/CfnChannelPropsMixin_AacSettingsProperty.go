package mixinsawsmedialive


// The settings for an AAC audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aacSettingsProperty := &AacSettingsProperty{
//   	Bitrate: jsii.Number(123),
//   	CodingMode: jsii.String("codingMode"),
//   	InputType: jsii.String("inputType"),
//   	Profile: jsii.String("profile"),
//   	RateControlMode: jsii.String("rateControlMode"),
//   	RawFormat: jsii.String("rawFormat"),
//   	SampleRate: jsii.Number(123),
//   	Spec: jsii.String("spec"),
//   	VbrQuality: jsii.String("vbrQuality"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html
//
type CfnChannelPropsMixin_AacSettingsProperty struct {
	// The average bitrate in bits/second.
	//
	// Valid values depend on the rate control mode and profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-bitrate
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Mono, stereo, or 5.1 channel layout. Valid values depend on the rate control mode and profile. The adReceiverMix setting receives a stereo description plus control track, and emits a mono AAC encode of the description track, with control data emitted in the PES header as per ETSI TS 101 154 Annex E.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-codingmode
	//
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Set to broadcasterMixedAd when the input contains pre-mixed main audio + AD (narration) as a stereo pair.
	//
	// The Audio Type field (audioType) will be set to 3, which signals to downstream systems that this stream contains broadcaster mixed AD. Note that the input received by the encoder must contain pre-mixed audio; MediaLive does not perform the mixing. The values in audioTypeControl and audioType (in AudioDescription) are ignored when set to broadcasterMixedAd. Leave this set to normal when the input does not contain pre-mixed audio + AD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-inputtype
	//
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// The AAC profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-profile
	//
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// The rate control mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-ratecontrolmode
	//
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the LATM/LOAS AAC output for raw containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-rawformat
	//
	RawFormat *string `field:"optional" json:"rawFormat" yaml:"rawFormat"`
	// The sample rate in Hz.
	//
	// Valid values depend on the rate control mode and profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-samplerate
	//
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
	// Uses MPEG-2 AAC audio instead of MPEG-4 AAC audio for raw or MPEG-2 Transport Stream containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-spec
	//
	Spec *string `field:"optional" json:"spec" yaml:"spec"`
	// The VBR quality level.
	//
	// This is used only if rateControlMode is VBR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-vbrquality
	//
	VbrQuality *string `field:"optional" json:"vbrQuality" yaml:"vbrQuality"`
}

