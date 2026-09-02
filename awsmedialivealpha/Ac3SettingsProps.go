package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for AC3 codec settings.
//
// Example:
//   // AAC stereo
//   aac := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
//   	Name: jsii.String("aac_stereo"),
//   	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
//   		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
//   		CodingMode: medialive.AacCodingMode_CODING_MODE_2_0(),
//   	}),
//   })
//
//   // AC3 5.1
//   ac3 := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
//   	Name: jsii.String("ac3_surround"),
//   	Codec: medialive.AudioCodecSettings_Ac3(&Ac3SettingsProps{
//   		Bitrate: awscdk.Bitrate_*Kbps(jsii.Number(384)),
//   		CodingMode: medialive.Ac3CodingMode_CODING_MODE_3_2_LFE(),
//   	}),
//   })
//
// Experimental.
type Ac3SettingsProps struct {
	// Applies a 3 dB attenuation to the surround channels.
	//
	// Used only for the 3/2 coding mode.
	// Default: - service default.
	//
	// Experimental.
	AttenuationControl Ac3AttenuationControl `field:"optional" json:"attenuationControl" yaml:"attenuationControl"`
	// The average bitrate.
	// Default: - service default.
	//
	// Experimental.
	Bitrate awscdk.Bitrate `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Specifies the bitstream mode (bsmod) for the emitted AC-3 stream.
	// Default: Ac3BitstreamMode.COMPLETE_MAIN
	//
	// Experimental.
	BitstreamMode Ac3BitstreamMode `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// The Dolby Digital coding mode.
	// Default: Ac3CodingMode.CODING_MODE_2_0
	//
	// Experimental.
	CodingMode Ac3CodingMode `field:"optional" json:"codingMode" yaml:"codingMode"`
	// The dialogue normalization level (1–31).
	// Default: - service default.
	//
	// Experimental.
	DialNorm *float64 `field:"optional" json:"dialNorm" yaml:"dialNorm"`
	// If set to filmStandard, adds dynamic range compression signaling to the output bitstream.
	// Default: - service default.
	//
	// Experimental.
	DrcProfile Ac3DrcProfile `field:"optional" json:"drcProfile" yaml:"drcProfile"`
	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	//
	// Valid only in codingMode32Lfe mode.
	// Default: Ac3LfeFilter.DISABLED
	//
	// Experimental.
	LfeFilter Ac3LfeFilter `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this audio data.
	// Default: - service default.
	//
	// Experimental.
	MetadataControl Ac3MetadataControl `field:"optional" json:"metadataControl" yaml:"metadataControl"`
}

