package awsquicksight


// The map state properties for a map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialMapStateProperty := &GeospatialMapStateProperty{
//   	Bounds: &GeospatialCoordinateBoundsProperty{
//   		East: jsii.Number(123),
//   		North: jsii.Number(123),
//   		South: jsii.Number(123),
//   		West: jsii.Number(123),
//   	},
//   	MapNavigation: jsii.String("mapNavigation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstate.html
//
type CfnAnalysis_GeospatialMapStateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstate.html#cfn-quicksight-analysis-geospatialmapstate-bounds
	//
	Bounds interface{} `field:"optional" json:"bounds" yaml:"bounds"`
	// Enables or disables map navigation for a map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstate.html#cfn-quicksight-analysis-geospatialmapstate-mapnavigation
	//
	MapNavigation *string `field:"optional" json:"mapNavigation" yaml:"mapNavigation"`
}

