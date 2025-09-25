package awsmediatailor


// A reference to a PlaybackConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   playbackConfigurationReference := &PlaybackConfigurationReference{
//   	PlaybackConfigurationArn: jsii.String("playbackConfigurationArn"),
//   	PlaybackConfigurationName: jsii.String("playbackConfigurationName"),
//   }
//
type PlaybackConfigurationReference struct {
	// The ARN of the PlaybackConfiguration resource.
	PlaybackConfigurationArn *string `field:"required" json:"playbackConfigurationArn" yaml:"playbackConfigurationArn"`
	// The Name of the PlaybackConfiguration resource.
	PlaybackConfigurationName *string `field:"required" json:"playbackConfigurationName" yaml:"playbackConfigurationName"`
}

