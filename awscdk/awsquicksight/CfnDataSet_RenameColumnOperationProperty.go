package awsquicksight


// A transform operation that renames a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameColumnOperationProperty := &RenameColumnOperationProperty{
//   	ColumnName: jsii.String("columnName"),
//
//   	// the properties below are optional
//   	NewColumnName: jsii.String("newColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-renamecolumnoperation.html
//
type CfnDataSet_RenameColumnOperationProperty struct {
	// The name of the column to be renamed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-renamecolumnoperation.html#cfn-quicksight-dataset-renamecolumnoperation-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The new name for the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-renamecolumnoperation.html#cfn-quicksight-dataset-renamecolumnoperation-newcolumnname
	//
	NewColumnName *string `field:"optional" json:"newColumnName" yaml:"newColumnName"`
}

