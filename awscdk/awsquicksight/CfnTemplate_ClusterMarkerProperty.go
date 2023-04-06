package awsquicksight


// The cluster marker that is a part of the cluster marker configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterMarkerProperty := &ClusterMarkerProperty{
//   	SimpleClusterMarker: &SimpleClusterMarkerProperty{
//   		Color: jsii.String("color"),
//   	},
//   }
//
type CfnTemplate_ClusterMarkerProperty struct {
	// The simple cluster marker of the cluster marker.
	SimpleClusterMarker interface{} `field:"optional" json:"simpleClusterMarker" yaml:"simpleClusterMarker"`
}

