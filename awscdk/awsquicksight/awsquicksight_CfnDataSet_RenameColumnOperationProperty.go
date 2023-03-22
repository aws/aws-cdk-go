package awsquicksight


// A transform operation that renames a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameColumnOperationProperty := &renameColumnOperationProperty{
//   	columnName: jsii.String("columnName"),
//   	newColumnName: jsii.String("newColumnName"),
//   }
//
type CfnDataSet_RenameColumnOperationProperty struct {
	// The name of the column to be renamed.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The new name for the column.
	NewColumnName *string `field:"required" json:"newColumnName" yaml:"newColumnName"`
}

