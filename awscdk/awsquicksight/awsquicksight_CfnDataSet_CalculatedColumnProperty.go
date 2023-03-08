package awsquicksight


// A calculated column for a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedColumnProperty := &calculatedColumnProperty{
//   	columnId: jsii.String("columnId"),
//   	columnName: jsii.String("columnName"),
//   	expression: jsii.String("expression"),
//   }
//
type CfnDataSet_CalculatedColumnProperty struct {
	// A unique ID to identify a calculated column.
	//
	// During a dataset update, if the column ID of a calculated column matches that of an existing calculated column, Amazon QuickSight preserves the existing calculated column.
	ColumnId *string `field:"required" json:"columnId" yaml:"columnId"`
	// Column name.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// An expression that defines the calculated column.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

