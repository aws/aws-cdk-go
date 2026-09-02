package awsmedialivealpha


// A mapping from input channels to an output channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioChannelMapping := &AudioChannelMapping{
//   	InputChannelLevels: []InputChannelLevel{
//   		&InputChannelLevel{
//   			InputChannel: jsii.Number(123),
//
//   			// the properties below are optional
//   			Gain: jsii.Number(123),
//   		},
//   	},
//   	OutputChannel: jsii.Number(123),
//   }
//
// Experimental.
type AudioChannelMapping struct {
	// The input channels and their gain levels to mix into this output channel.
	// Experimental.
	InputChannelLevels *[]*InputChannelLevel `field:"required" json:"inputChannelLevels" yaml:"inputChannelLevels"`
	// The index of the output channel being produced.
	// Experimental.
	OutputChannel *float64 `field:"required" json:"outputChannel" yaml:"outputChannel"`
}

