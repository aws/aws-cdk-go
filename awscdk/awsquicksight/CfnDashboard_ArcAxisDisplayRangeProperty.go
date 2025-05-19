package awsquicksight


// The arc axis range of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcAxisDisplayRangeProperty := &ArcAxisDisplayRangeProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-arcaxisdisplayrange.html
//
type CfnDashboard_ArcAxisDisplayRangeProperty struct {
	// The maximum value of the arc axis range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-arcaxisdisplayrange.html#cfn-quicksight-dashboard-arcaxisdisplayrange-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum value of the arc axis range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-arcaxisdisplayrange.html#cfn-quicksight-dashboard-arcaxisdisplayrange-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

