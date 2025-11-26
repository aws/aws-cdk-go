package previewawsmedialivemixins


// Information about one audio track to extract. You can select multiple tracks.
//
// The parent of this entity is AudioTrackSelection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioTrackProperty := &AudioTrackProperty{
//   	Track: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrack.html
//
type CfnChannelPropsMixin_AudioTrackProperty struct {
	// 1-based integer value that maps to a specific audio track.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiotrack.html#cfn-medialive-channel-audiotrack-track
	//
	Track *float64 `field:"optional" json:"track" yaml:"track"`
}

