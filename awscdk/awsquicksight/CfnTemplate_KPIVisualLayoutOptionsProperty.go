package awsquicksight


// The options that determine the layout a KPI visual.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisuallayoutoptions.html
//
type CfnTemplate_KPIVisualLayoutOptionsProperty struct {
	// The standard layout of the KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisuallayoutoptions.html#cfn-quicksight-template-kpivisuallayoutoptions-standardlayout
	//
	StandardLayout interface{} `field:"optional" json:"standardLayout" yaml:"standardLayout"`
}

