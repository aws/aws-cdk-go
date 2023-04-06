package awsquicksight


// The cluster marker configuration of the geospatial map selected point style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterMarkerConfigurationProperty := &ClusterMarkerConfigurationProperty{
//   	ClusterMarker: &ClusterMarkerProperty{
//   		SimpleClusterMarker: &SimpleClusterMarkerProperty{
//   			Color: jsii.String("color"),
//   		},
//   	},
//   }
//
type CfnTemplate_ClusterMarkerConfigurationProperty struct {
	// The cluster marker that is a part of the cluster marker configuration.
	ClusterMarker interface{} `field:"optional" json:"clusterMarker" yaml:"clusterMarker"`
}

