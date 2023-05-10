package awsquicksight


// The reference that specifies where the axis label is applied to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisLabelReferenceOptionsProperty := &AxisLabelReferenceOptionsProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FieldId: jsii.String("fieldId"),
//   }
//
type CfnTemplate_AxisLabelReferenceOptionsProperty struct {
	// The column that the axis label is targeted to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The field that the axis label is targeted to.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}

