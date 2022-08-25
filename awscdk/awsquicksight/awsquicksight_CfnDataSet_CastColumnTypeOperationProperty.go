package awsquicksight


// A transform operation that casts a column to a different type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   castColumnTypeOperationProperty := &castColumnTypeOperationProperty{
//   	columnName: jsii.String("columnName"),
//   	newColumnType: jsii.String("newColumnType"),
//
//   	// the properties below are optional
//   	format: jsii.String("format"),
//   }
//
type CfnDataSet_CastColumnTypeOperationProperty struct {
	// Column name.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// New column data type.
	NewColumnType *string `field:"required" json:"newColumnType" yaml:"newColumnType"`
	// When casting a column from string to datetime type, you can supply a string in a format supported by Amazon QuickSight to denote the source data format.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

