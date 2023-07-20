package awsquicksight


// The selected field options for the pivot table field options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldOptionProperty := &PivotTableFieldOptionProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	CustomLabel: jsii.String("customLabel"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoption.html
//
type CfnDashboard_PivotTableFieldOptionProperty struct {
	// The field ID of the pivot table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoption.html#cfn-quicksight-dashboard-pivottablefieldoption-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The custom label of the pivot table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoption.html#cfn-quicksight-dashboard-pivottablefieldoption-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The visibility of the pivot table field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoption.html#cfn-quicksight-dashboard-pivottablefieldoption-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

