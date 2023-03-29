package awsquicksight


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
type CfnTemplate_GeospatialPointStyleOptionsProperty struct {
	// `CfnTemplate.GeospatialPointStyleOptionsProperty.ClusterMarkerConfiguration`.
	ClusterMarkerConfiguration interface{} `field:"optional" json:"clusterMarkerConfiguration" yaml:"clusterMarkerConfiguration"`
	// `CfnTemplate.GeospatialPointStyleOptionsProperty.SelectedPointStyle`.
	SelectedPointStyle *string `field:"optional" json:"selectedPointStyle" yaml:"selectedPointStyle"`
}

