package awsmediatailor


// The configuration for bumpers.
//
// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bumperProperty := &bumperProperty{
//   	endUrl: jsii.String("endUrl"),
//   	startUrl: jsii.String("startUrl"),
//   }
//
type CfnPlaybackConfiguration_BumperProperty struct {
	// The URL for the end bumper asset.
	EndUrl *string `field:"optional" json:"endUrl" yaml:"endUrl"`
	// The URL for the start bumper asset.
	StartUrl *string `field:"optional" json:"startUrl" yaml:"startUrl"`
}

