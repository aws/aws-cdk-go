package previewawsquicksightmixins


// The heatmap configuration of the geospatial point style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialHeatmapConfigurationProperty := &GeospatialHeatmapConfigurationProperty{
//   	HeatmapColor: &GeospatialHeatmapColorScaleProperty{
//   		Colors: []interface{}{
//   			&GeospatialHeatmapDataColorProperty{
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialheatmapconfiguration.html
//
type CfnDashboardPropsMixin_GeospatialHeatmapConfigurationProperty struct {
	// The color scale specification for the heatmap point style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialheatmapconfiguration.html#cfn-quicksight-dashboard-geospatialheatmapconfiguration-heatmapcolor
	//
	HeatmapColor interface{} `field:"optional" json:"heatmapColor" yaml:"heatmapColor"`
}

