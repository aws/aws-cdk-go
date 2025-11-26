package previewawsmedialivemixins


// The settings for an AC3 audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ac3SettingsProperty := &Ac3SettingsProperty{
//   	AttenuationControl: jsii.String("attenuationControl"),
//   	Bitrate: jsii.Number(123),
//   	BitstreamMode: jsii.String("bitstreamMode"),
//   	CodingMode: jsii.String("codingMode"),
//   	Dialnorm: jsii.Number(123),
//   	DrcProfile: jsii.String("drcProfile"),
//   	LfeFilter: jsii.String("lfeFilter"),
//   	MetadataControl: jsii.String("metadataControl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html
//
type CfnChannelPropsMixin_Ac3SettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-attenuationcontrol
	//
	AttenuationControl *string `field:"optional" json:"attenuationControl" yaml:"attenuationControl"`
	// The average bitrate in bits/second.
	//
	// Valid bitrates depend on the coding mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-bitrate
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Specifies the bitstream mode (bsmod) for the emitted AC-3 stream.
	//
	// For more information about these values, see ATSC A/52-2012.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-bitstreammode
	//
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// The Dolby Digital coding mode.
	//
	// This determines the number of channels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-codingmode
	//
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Sets the dialnorm for the output.
	//
	// If excluded and the input audio is Dolby Digital, dialnorm is passed through.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-dialnorm
	//
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// If set to filmStandard, adds dynamic range compression signaling to the output bitstream as defined in the Dolby Digital specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-drcprofile
	//
	DrcProfile *string `field:"optional" json:"drcProfile" yaml:"drcProfile"`
	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	//
	// This is valid only in codingMode32Lfe mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-lfefilter
	//
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this audio data.
	//
	// If the audio is supplied from one of these streams, the static metadata settings are used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ac3settings.html#cfn-medialive-channel-ac3settings-metadatacontrol
	//
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
}

