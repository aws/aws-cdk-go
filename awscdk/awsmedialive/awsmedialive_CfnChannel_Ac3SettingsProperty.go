package awsmedialive


// The settings for an AC3 audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ac3SettingsProperty := &ac3SettingsProperty{
//   	bitrate: jsii.Number(123),
//   	bitstreamMode: jsii.String("bitstreamMode"),
//   	codingMode: jsii.String("codingMode"),
//   	dialnorm: jsii.Number(123),
//   	drcProfile: jsii.String("drcProfile"),
//   	lfeFilter: jsii.String("lfeFilter"),
//   	metadataControl: jsii.String("metadataControl"),
//   }
//
type CfnChannel_Ac3SettingsProperty struct {
	// The average bitrate in bits/second.
	//
	// Valid bitrates depend on the coding mode.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Specifies the bitstream mode (bsmod) for the emitted AC-3 stream.
	//
	// For more information about these values, see ATSC A/52-2012.
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// The Dolby Digital coding mode.
	//
	// This determines the number of channels.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Sets the dialnorm for the output.
	//
	// If excluded and the input audio is Dolby Digital, dialnorm is passed through.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// If set to filmStandard, adds dynamic range compression signaling to the output bitstream as defined in the Dolby Digital specification.
	DrcProfile *string `field:"optional" json:"drcProfile" yaml:"drcProfile"`
	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	//
	// This is valid only in codingMode32Lfe mode.
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this audio data.
	//
	// If the audio is supplied from one of these streams, the static metadata settings are used.
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
}

