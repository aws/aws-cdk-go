package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aspectRatioProperty := &AspectRatioProperty{
//   	Height: jsii.Number(123),
//   	Width: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-aspectratio.html
//
type CfnFeed_AspectRatioProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-aspectratio.html#cfn-elementalinference-feed-aspectratio-height
	//
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-aspectratio.html#cfn-elementalinference-feed-aspectratio-width
	//
	Width *float64 `field:"required" json:"width" yaml:"width"`
}

