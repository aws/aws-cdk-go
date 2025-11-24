package mixinsawsquicksight


// The options for the label thta is located above the row headers.
//
// This option is only applicable when `RowsLayout` is set to `HIERARCHY` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableRowsLabelOptionsProperty := &PivotTableRowsLabelOptionsProperty{
//   	CustomLabel: jsii.String("customLabel"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablerowslabeloptions.html
//
type CfnTemplatePropsMixin_PivotTableRowsLabelOptionsProperty struct {
	// The custom label string for the rows label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablerowslabeloptions.html#cfn-quicksight-template-pivottablerowslabeloptions-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The visibility of the rows label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablerowslabeloptions.html#cfn-quicksight-template-pivottablerowslabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

