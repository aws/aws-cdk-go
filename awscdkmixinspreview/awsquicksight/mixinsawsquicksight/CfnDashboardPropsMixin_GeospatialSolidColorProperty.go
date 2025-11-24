package mixinsawsquicksight


// The definition for a solid color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialSolidColorProperty := &GeospatialSolidColorProperty{
//   	Color: jsii.String("color"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialsolidcolor.html
//
type CfnDashboardPropsMixin_GeospatialSolidColorProperty struct {
	// The color and opacity values for the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialsolidcolor.html#cfn-quicksight-dashboard-geospatialsolidcolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Enables and disables the view state of the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialsolidcolor.html#cfn-quicksight-dashboard-geospatialsolidcolor-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

