package awsmedialivealpha


// An input channel level for audio remixing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   inputChannelLevel := &InputChannelLevel{
//   	InputChannel: jsii.Number(123),
//
//   	// the properties below are optional
//   	Gain: jsii.Number(123),
//   }
//
// Experimental.
type InputChannelLevel struct {
	// The index of the input channel to use as a source.
	// Experimental.
	InputChannel *float64 `field:"required" json:"inputChannel" yaml:"inputChannel"`
	// The remixing gain in dB (-60 to 6).
	// Default: 0.
	//
	// Experimental.
	Gain *float64 `field:"optional" json:"gain" yaml:"gain"`
}

