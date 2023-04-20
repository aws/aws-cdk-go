package awsquicksight


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
type CfnDashboard_AxisLabelReferenceOptionsProperty struct {
	// `CfnDashboard.AxisLabelReferenceOptionsProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnDashboard.AxisLabelReferenceOptionsProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}

