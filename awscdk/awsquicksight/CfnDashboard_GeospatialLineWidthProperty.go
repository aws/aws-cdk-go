package awsquicksight


// The width properties for a line.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialLineWidthProperty := &GeospatialLineWidthProperty{
//   	LineWidth: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallinewidth.html
//
type CfnDashboard_GeospatialLineWidthProperty struct {
	// The positive value for the width of a line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallinewidth.html#cfn-quicksight-dashboard-geospatiallinewidth-linewidth
	//
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}

