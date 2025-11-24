package mixinsawsquicksight


// The standard layout of the KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kPIVisualStandardLayoutProperty := &KPIVisualStandardLayoutProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpivisualstandardlayout.html
//
type CfnDashboardPropsMixin_KPIVisualStandardLayoutProperty struct {
	// The standard layout type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpivisualstandardlayout.html#cfn-quicksight-dashboard-kpivisualstandardlayout-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

