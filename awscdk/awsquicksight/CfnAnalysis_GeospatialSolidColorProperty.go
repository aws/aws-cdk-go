package awsquicksight


// The definition for a solid color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialSolidColorProperty := &GeospatialSolidColorProperty{
//   	Color: jsii.String("color"),
//
//   	// the properties below are optional
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialsolidcolor.html
//
type CfnAnalysis_GeospatialSolidColorProperty struct {
	// The color and opacity values for the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialsolidcolor.html#cfn-quicksight-analysis-geospatialsolidcolor-color
	//
	Color *string `field:"required" json:"color" yaml:"color"`
	// Enables and disables the view state of the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialsolidcolor.html#cfn-quicksight-analysis-geospatialsolidcolor-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

