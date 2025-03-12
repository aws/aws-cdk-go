package awsquicksight


// The color to be used in the heatmap point style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialHeatmapDataColorProperty := &GeospatialHeatmapDataColorProperty{
//   	Color: jsii.String("color"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialheatmapdatacolor.html
//
type CfnAnalysis_GeospatialHeatmapDataColorProperty struct {
	// The hex color to be used in the heatmap point style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialheatmapdatacolor.html#cfn-quicksight-analysis-geospatialheatmapdatacolor-color
	//
	Color *string `field:"required" json:"color" yaml:"color"`
}

