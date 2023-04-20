package awsquicksight


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
type CfnDashboard_ClusterMarkerProperty struct {
	// `CfnDashboard.ClusterMarkerProperty.SimpleClusterMarker`.
	SimpleClusterMarker interface{} `field:"optional" json:"simpleClusterMarker" yaml:"simpleClusterMarker"`
}

