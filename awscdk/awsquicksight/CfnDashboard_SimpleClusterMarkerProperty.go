package awsquicksight


// The simple cluster marker of the cluster marker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleClusterMarkerProperty := &SimpleClusterMarkerProperty{
//   	Color: jsii.String("color"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-simpleclustermarker.html
//
type CfnDashboard_SimpleClusterMarkerProperty struct {
	// The color of the simple cluster marker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-simpleclustermarker.html#cfn-quicksight-dashboard-simpleclustermarker-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
}

