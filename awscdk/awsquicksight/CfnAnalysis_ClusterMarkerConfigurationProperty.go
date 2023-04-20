package awsquicksight


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
type CfnAnalysis_ClusterMarkerConfigurationProperty struct {
	// `CfnAnalysis.ClusterMarkerConfigurationProperty.ClusterMarker`.
	ClusterMarker interface{} `field:"optional" json:"clusterMarker" yaml:"clusterMarker"`
}

