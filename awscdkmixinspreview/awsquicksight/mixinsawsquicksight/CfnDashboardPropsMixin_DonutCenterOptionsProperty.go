package mixinsawsquicksight


// The label options of the label that is displayed in the center of a donut chart.
//
// This option isn't available for pie charts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   donutCenterOptionsProperty := &DonutCenterOptionsProperty{
//   	LabelVisibility: jsii.String("labelVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-donutcenteroptions.html
//
type CfnDashboardPropsMixin_DonutCenterOptionsProperty struct {
	// Determines the visibility of the label in a donut chart.
	//
	// In the Quick Sight console, this option is called `'Show total'` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-donutcenteroptions.html#cfn-quicksight-dashboard-donutcenteroptions-labelvisibility
	//
	LabelVisibility *string `field:"optional" json:"labelVisibility" yaml:"labelVisibility"`
}

