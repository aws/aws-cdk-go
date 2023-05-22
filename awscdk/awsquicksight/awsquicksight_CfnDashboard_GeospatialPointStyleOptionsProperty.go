package awsquicksight


// The point style of the geospatial map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialPointStyleOptionsProperty := &GeospatialPointStyleOptionsProperty{
//   	ClusterMarkerConfiguration: &ClusterMarkerConfigurationProperty{
//   		ClusterMarker: &ClusterMarkerProperty{
//   			SimpleClusterMarker: &SimpleClusterMarkerProperty{
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   	SelectedPointStyle: jsii.String("selectedPointStyle"),
//   }
//
type CfnDashboard_GeospatialPointStyleOptionsProperty struct {
	// The cluster marker configuration of the geospatial point style.
	ClusterMarkerConfiguration interface{} `field:"optional" json:"clusterMarkerConfiguration" yaml:"clusterMarkerConfiguration"`
	// The selected point styles (point, cluster) of the geospatial map.
	SelectedPointStyle *string `field:"optional" json:"selectedPointStyle" yaml:"selectedPointStyle"`
}

