package awsmedialive


// The settings for an EAC3 audio encode in the output.
//
// The parent of this entity is AudioCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eac3SettingsProperty := &Eac3SettingsProperty{
//   	AttenuationControl: jsii.String("attenuationControl"),
//   	Bitrate: jsii.Number(123),
//   	BitstreamMode: jsii.String("bitstreamMode"),
//   	CodingMode: jsii.String("codingMode"),
//   	DcFilter: jsii.String("dcFilter"),
//   	Dialnorm: jsii.Number(123),
//   	DrcLine: jsii.String("drcLine"),
//   	DrcRf: jsii.String("drcRf"),
//   	LfeControl: jsii.String("lfeControl"),
//   	LfeFilter: jsii.String("lfeFilter"),
//   	LoRoCenterMixLevel: jsii.Number(123),
//   	LoRoSurroundMixLevel: jsii.Number(123),
//   	LtRtCenterMixLevel: jsii.Number(123),
//   	LtRtSurroundMixLevel: jsii.Number(123),
//   	MetadataControl: jsii.String("metadataControl"),
//   	PassthroughControl: jsii.String("passthroughControl"),
//   	PhaseControl: jsii.String("phaseControl"),
//   	StereoDownmix: jsii.String("stereoDownmix"),
//   	SurroundExMode: jsii.String("surroundExMode"),
//   	SurroundMode: jsii.String("surroundMode"),
//   }
//
type CfnChannel_Eac3SettingsProperty struct {
	// When set to attenuate3Db, applies a 3 dB attenuation to the surround channels.
	//
	// Used only for the 3/2 coding mode.
	AttenuationControl *string `field:"optional" json:"attenuationControl" yaml:"attenuationControl"`
	// The average bitrate in bits/second.
	//
	// Valid bitrates depend on the coding mode.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Specifies the bitstream mode (bsmod) for the emitted E-AC-3 stream.
	//
	// For more information, see ATSC A/52-2012 (Annex E).
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// The Dolby Digital Plus coding mode.
	//
	// This mode determines the number of channels.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// When set to enabled, activates a DC highpass filter for all input channels.
	DcFilter *string `field:"optional" json:"dcFilter" yaml:"dcFilter"`
	// Sets the dialnorm for the output.
	//
	// If blank and the input audio is Dolby Digital Plus, dialnorm will be passed through.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// Sets the Dolby dynamic range compression profile.
	DrcLine *string `field:"optional" json:"drcLine" yaml:"drcLine"`
	// Sets the profile for heavy Dolby dynamic range compression, ensuring that the instantaneous signal peaks do not exceed specified levels.
	DrcRf *string `field:"optional" json:"drcRf" yaml:"drcRf"`
	// When encoding 3/2 audio, setting to lfe enables the LFE channel.
	LfeControl *string `field:"optional" json:"lfeControl" yaml:"lfeControl"`
	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	//
	// Valid only with a codingMode32 coding mode.
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// The Left only/Right only center mix level.
	//
	// Used only for the 3/2 coding mode.
	LoRoCenterMixLevel *float64 `field:"optional" json:"loRoCenterMixLevel" yaml:"loRoCenterMixLevel"`
	// The Left only/Right only surround mix level.
	//
	// Used only for a 3/2 coding mode.
	LoRoSurroundMixLevel *float64 `field:"optional" json:"loRoSurroundMixLevel" yaml:"loRoSurroundMixLevel"`
	// The Left total/Right total center mix level.
	//
	// Used only for a 3/2 coding mode.
	LtRtCenterMixLevel *float64 `field:"optional" json:"ltRtCenterMixLevel" yaml:"ltRtCenterMixLevel"`
	// The Left total/Right total surround mix level.
	//
	// Used only for the 3/2 coding mode.
	LtRtSurroundMixLevel *float64 `field:"optional" json:"ltRtSurroundMixLevel" yaml:"ltRtSurroundMixLevel"`
	// When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this audio data.
	//
	// If the audio is not supplied from one of these streams, then the static metadata settings are used.
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
	// When set to whenPossible, input DD+ audio will be passed through if it is present on the input.
	//
	// This detection is dynamic over the life of the transcode. Inputs that alternate between DD+ and non-DD+ content will have a consistent DD+ output as the system alternates between passthrough and encoding.
	PassthroughControl *string `field:"optional" json:"passthroughControl" yaml:"passthroughControl"`
	// When set to shift90Degrees, applies a 90-degree phase shift to the surround channels.
	//
	// Used only for a 3/2 coding mode.
	PhaseControl *string `field:"optional" json:"phaseControl" yaml:"phaseControl"`
	// A stereo downmix preference.
	//
	// Used only for the 3/2 coding mode.
	StereoDownmix *string `field:"optional" json:"stereoDownmix" yaml:"stereoDownmix"`
	// When encoding 3/2 audio, sets whether an extra center back surround channel is matrix encoded into the left and right surround channels.
	SurroundExMode *string `field:"optional" json:"surroundExMode" yaml:"surroundExMode"`
	// When encoding 2/0 audio, sets whether Dolby Surround is matrix-encoded into the two channels.
	SurroundMode *string `field:"optional" json:"surroundMode" yaml:"surroundMode"`
}

