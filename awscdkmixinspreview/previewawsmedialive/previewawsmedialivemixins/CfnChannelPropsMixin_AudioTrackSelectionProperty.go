package previewawsmedialivemixins


// Information about the audio track to extract.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrackselection.html
//
type CfnChannelPropsMixin_AudioTrackSelectionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrackselection.html#cfn-medialive-channel-audiotrackselection-dolbyedecode
	//
	DolbyEDecode interface{} `field:"optional" json:"dolbyEDecode" yaml:"dolbyEDecode"`
	// Selects one or more unique audio tracks from within a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrackselection.html#cfn-medialive-channel-audiotrackselection-tracks
	//
	Tracks interface{} `field:"optional" json:"tracks" yaml:"tracks"`
}

