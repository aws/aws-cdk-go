package awsmedialivealpha


// Configuration for a single audio track in a track-based selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var audioPreMixerSettings AudioPreMixerSettings
//
//   audioTrackConfig := &AudioTrackConfig{
//   	Track: jsii.Number(123),
//
//   	// the properties below are optional
//   	PremixSettings: audioPreMixerSettings,
//   }
//
// Experimental.
type AudioTrackConfig struct {
	// The 1-based track number to extract.
	// Experimental.
	Track *float64 `field:"required" json:"track" yaml:"track"`
	// Pre-mixer settings for this track (gain, channel remix, loudness normalization).
	// Default: - no pre-mixing.
	//
	// Experimental.
	PremixSettings AudioPreMixerSettings `field:"optional" json:"premixSettings" yaml:"premixSettings"`
}

