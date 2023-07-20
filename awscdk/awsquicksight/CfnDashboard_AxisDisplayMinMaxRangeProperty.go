package awsquicksight


// The minimum and maximum setup for an axis display range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisDisplayMinMaxRangeProperty := &AxisDisplayMinMaxRangeProperty{
//   	Maximum: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayminmaxrange.html
//
type CfnDashboard_AxisDisplayMinMaxRangeProperty struct {
	// The maximum setup for an axis display range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayminmaxrange.html#cfn-quicksight-dashboard-axisdisplayminmaxrange-maximum
	//
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// The minimum setup for an axis display range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisdisplayminmaxrange.html#cfn-quicksight-dashboard-axisdisplayminmaxrange-minimum
	//
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

