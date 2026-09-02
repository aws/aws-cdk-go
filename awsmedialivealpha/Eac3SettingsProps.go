package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for EAC3 codec settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var bitrate Bitrate
//   var eac3AttenuationControl Eac3AttenuationControl
//   var eac3BitstreamMode Eac3BitstreamMode
//   var eac3CodingMode Eac3CodingMode
//   var eac3DcFilter Eac3DcFilter
//   var eac3DrcLine Eac3DrcLine
//   var eac3DrcRf Eac3DrcRf
//   var eac3LfeControl Eac3LfeControl
//   var eac3LfeFilter Eac3LfeFilter
//   var eac3MetadataControl Eac3MetadataControl
//   var eac3PassthroughControl Eac3PassthroughControl
//   var eac3PhaseControl Eac3PhaseControl
//   var eac3StereoDownmix Eac3StereoDownmix
//   var eac3SurroundExMode Eac3SurroundExMode
//   var eac3SurroundMode Eac3SurroundMode
//
//   eac3SettingsProps := &Eac3SettingsProps{
//   	AttenuationControl: eac3AttenuationControl,
//   	Bitrate: bitrate,
//   	BitstreamMode: eac3BitstreamMode,
//   	CodingMode: eac3CodingMode,
//   	DcFilter: eac3DcFilter,
//   	DialNorm: jsii.Number(123),
//   	DrcLine: eac3DrcLine,
//   	DrcRf: eac3DrcRf,
//   	LfeControl: eac3LfeControl,
//   	LfeFilter: eac3LfeFilter,
//   	LoRoCenterMixLevel: jsii.Number(123),
//   	LoRoSurroundMixLevel: jsii.Number(123),
//   	LtRtCenterMixLevel: jsii.Number(123),
//   	LtRtSurroundMixLevel: jsii.Number(123),
//   	MetadataControl: eac3MetadataControl,
//   	PassthroughControl: eac3PassthroughControl,
//   	PhaseControl: eac3PhaseControl,
//   	StereoDownmix: eac3StereoDownmix,
//   	SurroundExMode: eac3SurroundExMode,
//   	SurroundMode: eac3SurroundMode,
//   }
//
// Experimental.
type Eac3SettingsProps struct {
	// When set to attenuate3Db, applies a 3 dB attenuation to the surround channels.
	//
	// Used only for the 3/2 coding mode.
	// Default: Eac3AttenuationControl.NONE
	//
	// Experimental.
	AttenuationControl Eac3AttenuationControl `field:"optional" json:"attenuationControl" yaml:"attenuationControl"`
	// The average bitrate.
	// Default: - service default.
	//
	// Experimental.
	Bitrate awscdk.Bitrate `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Specifies the bitstream mode (bsmod) for the emitted E-AC-3 stream.
	// Default: Eac3BitstreamMode.COMPLETE_MAIN
	//
	// Experimental.
	BitstreamMode Eac3BitstreamMode `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// The Dolby Digital Plus coding mode.
	// Default: Eac3CodingMode.CODING_MODE_3_2
	//
	// Experimental.
	CodingMode Eac3CodingMode `field:"optional" json:"codingMode" yaml:"codingMode"`
	// When set to enabled, activates a DC highpass filter for all input channels.
	// Default: - service default.
	//
	// Experimental.
	DcFilter Eac3DcFilter `field:"optional" json:"dcFilter" yaml:"dcFilter"`
	// The dialogue normalization level (1–31).
	// Default: - service default.
	//
	// Experimental.
	DialNorm *float64 `field:"optional" json:"dialNorm" yaml:"dialNorm"`
	// Sets the Dolby dynamic range compression profile.
	// Default: - service default.
	//
	// Experimental.
	DrcLine Eac3DrcLine `field:"optional" json:"drcLine" yaml:"drcLine"`
	// Sets the profile for heavy Dolby dynamic range compression, ensuring that the instantaneous signal peaks do not exceed specified levels.
	// Default: - service default.
	//
	// Experimental.
	DrcRf Eac3DrcRf `field:"optional" json:"drcRf" yaml:"drcRf"`
	// When encoding 3/2 audio, setting to lfe enables the LFE channel.
	// Default: - service default.
	//
	// Experimental.
	LfeControl Eac3LfeControl `field:"optional" json:"lfeControl" yaml:"lfeControl"`
	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	//
	// Valid only with a codingMode32 coding mode.
	// Default: - service default.
	//
	// Experimental.
	LfeFilter Eac3LfeFilter `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// The Left only/Right only center mix level.
	//
	// Used only for the 3/2 coding mode.
	// Default: - service default.
	//
	// Experimental.
	LoRoCenterMixLevel *float64 `field:"optional" json:"loRoCenterMixLevel" yaml:"loRoCenterMixLevel"`
	// The Left only/Right only surround mix level.
	//
	// Used only for a 3/2 coding mode.
	// Default: - service default.
	//
	// Experimental.
	LoRoSurroundMixLevel *float64 `field:"optional" json:"loRoSurroundMixLevel" yaml:"loRoSurroundMixLevel"`
	// The Left total/Right total center mix level.
	//
	// Used only for a 3/2 coding mode.
	// Default: - service default.
	//
	// Experimental.
	LtRtCenterMixLevel *float64 `field:"optional" json:"ltRtCenterMixLevel" yaml:"ltRtCenterMixLevel"`
	// The Left total/Right total surround mix level.
	//
	// Used only for the 3/2 coding mode.
	// Default: - service default.
	//
	// Experimental.
	LtRtSurroundMixLevel *float64 `field:"optional" json:"ltRtSurroundMixLevel" yaml:"ltRtSurroundMixLevel"`
	// When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this audio data.
	// Default: - service default.
	//
	// Experimental.
	MetadataControl Eac3MetadataControl `field:"optional" json:"metadataControl" yaml:"metadataControl"`
	// When set to whenPossible, input DD+ audio will be passed through if it is present on the input.
	// Default: - service default.
	//
	// Experimental.
	PassthroughControl Eac3PassthroughControl `field:"optional" json:"passthroughControl" yaml:"passthroughControl"`
	// When set to shift90Degrees, applies a 90-degree phase shift to the surround channels.
	//
	// Used only for a 3/2 coding mode.
	// Default: - service default.
	//
	// Experimental.
	PhaseControl Eac3PhaseControl `field:"optional" json:"phaseControl" yaml:"phaseControl"`
	// A stereo downmix preference.
	//
	// Used only for the 3/2 coding mode.
	// Default: - service default.
	//
	// Experimental.
	StereoDownmix Eac3StereoDownmix `field:"optional" json:"stereoDownmix" yaml:"stereoDownmix"`
	// When encoding 3/2 audio, sets whether an extra center back surround channel is matrix encoded into the left and right surround channels.
	// Default: - service default.
	//
	// Experimental.
	SurroundExMode Eac3SurroundExMode `field:"optional" json:"surroundExMode" yaml:"surroundExMode"`
	// When encoding 2/0 audio, sets whether Dolby Surround is matrix-encoded into the two channels.
	// Default: - service default.
	//
	// Experimental.
	SurroundMode Eac3SurroundMode `field:"optional" json:"surroundMode" yaml:"surroundMode"`
}

