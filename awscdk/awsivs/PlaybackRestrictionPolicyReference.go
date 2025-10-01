package awsivs


// A reference to a PlaybackRestrictionPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   playbackRestrictionPolicyReference := &PlaybackRestrictionPolicyReference{
//   	PlaybackRestrictionPolicyArn: jsii.String("playbackRestrictionPolicyArn"),
//   }
//
type PlaybackRestrictionPolicyReference struct {
	// The Arn of the PlaybackRestrictionPolicy resource.
	PlaybackRestrictionPolicyArn *string `field:"required" json:"playbackRestrictionPolicyArn" yaml:"playbackRestrictionPolicyArn"`
}

