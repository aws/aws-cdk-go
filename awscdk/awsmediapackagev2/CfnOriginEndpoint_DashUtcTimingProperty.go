package awsmediapackagev2


// <p>Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashUtcTimingProperty := &DashUtcTimingProperty{
//   	TimingMode: jsii.String("timingMode"),
//   	TimingSource: jsii.String("timingSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashutctiming.html
//
type CfnOriginEndpoint_DashUtcTimingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashutctiming.html#cfn-mediapackagev2-originendpoint-dashutctiming-timingmode
	//
	TimingMode *string `field:"optional" json:"timingMode" yaml:"timingMode"`
	// <p>The the method that the player uses to synchronize to coordinated universal time (UTC) wall clock time.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashutctiming.html#cfn-mediapackagev2-originendpoint-dashutctiming-timingsource
	//
	TimingSource *string `field:"optional" json:"timingSource" yaml:"timingSource"`
}

