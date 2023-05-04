package awsmedialive


// Information about the audio track to extract.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioTrackSelectionProperty := &AudioTrackSelectionProperty{
//   	DolbyEDecode: &AudioDolbyEDecodeProperty{
//   		ProgramSelection: jsii.String("programSelection"),
//   	},
//   	Tracks: []interface{}{
//   		&AudioTrackProperty{
//   			Track: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_AudioTrackSelectionProperty struct {
	// `CfnChannel.AudioTrackSelectionProperty.DolbyEDecode`.
	DolbyEDecode interface{} `field:"optional" json:"dolbyEDecode" yaml:"dolbyEDecode"`
	// Selects one or more unique audio tracks from within a source.
	Tracks interface{} `field:"optional" json:"tracks" yaml:"tracks"`
}

