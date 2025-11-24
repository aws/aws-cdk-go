package mixinsawsquicksight


// The cluster marker that is a part of the cluster marker configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterMarkerProperty := &ClusterMarkerProperty{
//   	SimpleClusterMarker: &SimpleClusterMarkerProperty{
//   		Color: jsii.String("color"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-clustermarker.html
//
type CfnDashboardPropsMixin_ClusterMarkerProperty struct {
	// The simple cluster marker of the cluster marker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-clustermarker.html#cfn-quicksight-dashboard-clustermarker-simpleclustermarker
	//
	SimpleClusterMarker interface{} `field:"optional" json:"simpleClusterMarker" yaml:"simpleClusterMarker"`
}

