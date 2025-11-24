package mixinsawsquicksight


// The reference that specifies where the axis label is applied to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   axisLabelReferenceOptionsProperty := &AxisLabelReferenceOptionsProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FieldId: jsii.String("fieldId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislabelreferenceoptions.html
//
type CfnDashboardPropsMixin_AxisLabelReferenceOptionsProperty struct {
	// The column that the axis label is targeted to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislabelreferenceoptions.html#cfn-quicksight-dashboard-axislabelreferenceoptions-column
	//
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// The field that the axis label is targeted to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislabelreferenceoptions.html#cfn-quicksight-dashboard-axislabelreferenceoptions-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

