package awsmediatailor


// The configuration for bumpers.
//
// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bumperProperty := &BumperProperty{
//   	EndUrl: jsii.String("endUrl"),
//   	StartUrl: jsii.String("startUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-bumper.html
//
type CfnPlaybackConfiguration_BumperProperty struct {
	// The URL for the end bumper asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-bumper.html#cfn-mediatailor-playbackconfiguration-bumper-endurl
	//
	EndUrl *string `field:"optional" json:"endUrl" yaml:"endUrl"`
	// The URL for the start bumper asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-bumper.html#cfn-mediatailor-playbackconfiguration-bumper-starturl
	//
	StartUrl *string `field:"optional" json:"startUrl" yaml:"startUrl"`
}

