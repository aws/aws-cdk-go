package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableRowsLabelOptionsProperty := &PivotTableRowsLabelOptionsProperty{
//   	CustomLabel: jsii.String("customLabel"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablerowslabeloptions.html
//
type CfnDashboard_PivotTableRowsLabelOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablerowslabeloptions.html#cfn-quicksight-dashboard-pivottablerowslabeloptions-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablerowslabeloptions.html#cfn-quicksight-dashboard-pivottablerowslabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

