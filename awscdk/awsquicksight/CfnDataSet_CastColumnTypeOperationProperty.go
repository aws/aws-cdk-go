package awsquicksight


// A transform operation that casts a column to a different type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   castColumnTypeOperationProperty := &CastColumnTypeOperationProperty{
//   	ColumnName: jsii.String("columnName"),
//   	NewColumnType: jsii.String("newColumnType"),
//
//   	// the properties below are optional
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html
//
type CfnDataSet_CastColumnTypeOperationProperty struct {
	// Column name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html#cfn-quicksight-dataset-castcolumntypeoperation-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// New column data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html#cfn-quicksight-dataset-castcolumntypeoperation-newcolumntype
	//
	NewColumnType *string `field:"required" json:"newColumnType" yaml:"newColumnType"`
	// When casting a column from string to datetime type, you can supply a string in a format supported by Amazon QuickSight to denote the source data format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html#cfn-quicksight-dataset-castcolumntypeoperation-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

