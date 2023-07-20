package awsquicksight


// The heatmap configuration of the geospatial point style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialheatmapconfiguration.html
//
type CfnAnalysis_GeospatialHeatmapConfigurationProperty struct {
	// The color scale specification for the heatmap point style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialheatmapconfiguration.html#cfn-quicksight-analysis-geospatialheatmapconfiguration-heatmapcolor
	//
	HeatmapColor interface{} `field:"optional" json:"heatmapColor" yaml:"heatmapColor"`
}

