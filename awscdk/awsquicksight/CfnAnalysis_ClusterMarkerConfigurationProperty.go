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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-clustermarkerconfiguration.html
//
type CfnAnalysis_ClusterMarkerConfigurationProperty struct {
	// The cluster marker that is a part of the cluster marker configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-clustermarkerconfiguration.html#cfn-quicksight-analysis-clustermarkerconfiguration-clustermarker
	//
	ClusterMarker interface{} `field:"optional" json:"clusterMarker" yaml:"clusterMarker"`
}

