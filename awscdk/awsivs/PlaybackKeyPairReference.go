package awsivs


// A reference to a PlaybackKeyPair resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   playbackKeyPairReference := &PlaybackKeyPairReference{
//   	PlaybackKeyPairArn: jsii.String("playbackKeyPairArn"),
//   }
//
type PlaybackKeyPairReference struct {
	// The Arn of the PlaybackKeyPair resource.
	PlaybackKeyPairArn *string `field:"required" json:"playbackKeyPairArn" yaml:"playbackKeyPairArn"`
}

