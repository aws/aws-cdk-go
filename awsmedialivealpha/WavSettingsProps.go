package awsmedialivealpha


// Properties for WAV codec settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var audioBitDepth AudioBitDepth
//   var audioSampleRate AudioSampleRate
//   var wavCodingMode WavCodingMode
//
//   wavSettingsProps := &WavSettingsProps{
//   	BitDepth: audioBitDepth,
//   	CodingMode: wavCodingMode,
//   	SampleRate: audioSampleRate,
//   }
//
// Experimental.
type WavSettingsProps struct {
	// The bit depth of the WAV output.
	// Default: AudioBitDepth.DEPTH_16
	//
	// Experimental.
	BitDepth AudioBitDepth `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// The audio coding mode for the WAV audio.
	// Default: WavCodingMode.CODING_MODE_2_0
	//
	// Experimental.
	CodingMode WavCodingMode `field:"optional" json:"codingMode" yaml:"codingMode"`
	// The sample rate.
	// Default: AudioSampleRate.HZ_48000
	//
	// Experimental.
	SampleRate AudioSampleRate `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

