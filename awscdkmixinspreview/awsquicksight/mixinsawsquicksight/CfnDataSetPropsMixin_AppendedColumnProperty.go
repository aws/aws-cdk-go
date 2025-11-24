package mixinsawsquicksight


// Represents a column that will be included in the result of an append operation, combining data from multiple sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   appendedColumnProperty := &AppendedColumnProperty{
//   	ColumnName: jsii.String("columnName"),
//   	NewColumnId: jsii.String("newColumnId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendedcolumn.html
//
type CfnDataSetPropsMixin_AppendedColumnProperty struct {
	// The name of the column to include in the appended result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendedcolumn.html#cfn-quicksight-dataset-appendedcolumn-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// A unique identifier for the column in the appended result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendedcolumn.html#cfn-quicksight-dataset-appendedcolumn-newcolumnid
	//
	NewColumnId *string `field:"optional" json:"newColumnId" yaml:"newColumnId"`
}

