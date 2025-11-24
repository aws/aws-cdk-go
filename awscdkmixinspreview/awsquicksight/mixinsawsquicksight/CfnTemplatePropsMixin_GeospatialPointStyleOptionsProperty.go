package mixinsawsquicksight


// The point style of the geospatial map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialPointStyleOptionsProperty := &GeospatialPointStyleOptionsProperty{
//   	ClusterMarkerConfiguration: &ClusterMarkerConfigurationProperty{
//   		ClusterMarker: &ClusterMarkerProperty{
//   			SimpleClusterMarker: &SimpleClusterMarkerProperty{
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   	HeatmapConfiguration: &GeospatialHeatmapConfigurationProperty{
//   		HeatmapColor: &GeospatialHeatmapColorScaleProperty{
//   			Colors: []interface{}{
//   				&GeospatialHeatmapDataColorProperty{
//   					Color: jsii.String("color"),
//   				},
//   			},
//   		},
//   	},
//   	SelectedPointStyle: jsii.String("selectedPointStyle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialpointstyleoptions.html
//
type CfnTemplatePropsMixin_GeospatialPointStyleOptionsProperty struct {
	// The cluster marker configuration of the geospatial point style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialpointstyleoptions.html#cfn-quicksight-template-geospatialpointstyleoptions-clustermarkerconfiguration
	//
	ClusterMarkerConfiguration interface{} `field:"optional" json:"clusterMarkerConfiguration" yaml:"clusterMarkerConfiguration"`
	// The heatmap configuration of the geospatial point style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialpointstyleoptions.html#cfn-quicksight-template-geospatialpointstyleoptions-heatmapconfiguration
	//
	HeatmapConfiguration interface{} `field:"optional" json:"heatmapConfiguration" yaml:"heatmapConfiguration"`
	// The selected point styles (point, cluster) of the geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialpointstyleoptions.html#cfn-quicksight-template-geospatialpointstyleoptions-selectedpointstyle
	//
	SelectedPointStyle *string `field:"optional" json:"selectedPointStyle" yaml:"selectedPointStyle"`
}

