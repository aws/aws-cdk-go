package awsmedialive


// The settings for remixing audio.
//
// The parent of this entity is RemixSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioChannelMappingProperty := &audioChannelMappingProperty{
//   	inputChannelLevels: []interface{}{
//   		&inputChannelLevelProperty{
//   			gain: jsii.Number(123),
//   			inputChannel: jsii.Number(123),
//   		},
//   	},
//   	outputChannel: jsii.Number(123),
//   }
//
type CfnChannel_AudioChannelMappingProperty struct {
	// The indices and gain values for each input channel that should be remixed into this output channel.
	InputChannelLevels interface{} `field:"optional" json:"inputChannelLevels" yaml:"inputChannelLevels"`
	// The index of the output channel that is being produced.
	OutputChannel *float64 `field:"optional" json:"outputChannel" yaml:"outputChannel"`
}

