package awsmedialivealpha


// Configuration for a single audio PID in a PID-based selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var audioPreMixerSettings AudioPreMixerSettings
//   var dolbyEProgramSelection DolbyEProgramSelection
//
//   audioPidConfig := &AudioPidConfig{
//   	Pid: jsii.Number(123),
//
//   	// the properties below are optional
//   	DolbyEDecode: dolbyEProgramSelection,
//   	PremixSettings: audioPreMixerSettings,
//   }
//
// Experimental.
type AudioPidConfig struct {
	// The packet identifier (PID) value from within the source.
	// Experimental.
	Pid *float64 `field:"required" json:"pid" yaml:"pid"`
	// Which Dolby E program to decode from this PID.
	// Default: - no Dolby E decoding.
	//
	// Experimental.
	DolbyEDecode DolbyEProgramSelection `field:"optional" json:"dolbyEDecode" yaml:"dolbyEDecode"`
	// Pre-mixer settings for this PID (gain, channel remix, loudness normalization).
	// Default: - no pre-mixing.
	//
	// Experimental.
	PremixSettings AudioPreMixerSettings `field:"optional" json:"premixSettings" yaml:"premixSettings"`
}

