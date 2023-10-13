package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPIVisualLayoutOptionsProperty := &KPIVisualLayoutOptionsProperty{
//   	StandardLayout: &KPIVisualStandardLayoutProperty{
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpivisuallayoutoptions.html
//
type CfnDashboard_KPIVisualLayoutOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpivisuallayoutoptions.html#cfn-quicksight-dashboard-kpivisuallayoutoptions-standardlayout
	//
	StandardLayout interface{} `field:"optional" json:"standardLayout" yaml:"standardLayout"`
}

