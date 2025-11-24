package mixinsawsquicksight


// The color to be used in the heatmap point style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialHeatmapDataColorProperty := &GeospatialHeatmapDataColorProperty{
//   	Color: jsii.String("color"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialheatmapdatacolor.html
//
type CfnDashboardPropsMixin_GeospatialHeatmapDataColorProperty struct {
	// The hex color to be used in the heatmap point style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialheatmapdatacolor.html#cfn-quicksight-dashboard-geospatialheatmapdatacolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
}

