package awsmedialive


// Information about one audio track to extract. You can select multiple tracks.
//
// The parent of this entity is AudioTrackSelection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioTrackProperty := &audioTrackProperty{
//   	track: jsii.Number(123),
//   }
//
type CfnChannel_AudioTrackProperty struct {
	// 1-based integer value that maps to a specific audio track.
	Track *float64 `field:"optional" json:"track" yaml:"track"`
}

