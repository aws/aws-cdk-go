package awsquicksight


// The column option of the transposed table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transposedTableOptionProperty := &TransposedTableOptionProperty{
//   	ColumnType: jsii.String("columnType"),
//
//   	// the properties below are optional
//   	ColumnIndex: jsii.Number(123),
//   	ColumnWidth: jsii.String("columnWidth"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-transposedtableoption.html
//
type CfnAnalysis_TransposedTableOptionProperty struct {
	// The column type of the column in a transposed table. Choose one of the following options:.
	//
	// - `ROW_HEADER_COLUMN` : Refers to the leftmost column of the row header in the transposed table.
	// - `VALUE_COLUMN` : Refers to all value columns in the transposed table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-transposedtableoption.html#cfn-quicksight-analysis-transposedtableoption-columntype
	//
	ColumnType *string `field:"required" json:"columnType" yaml:"columnType"`
	// The index of a columns in a transposed table.
	//
	// The index range is 0-9999.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-transposedtableoption.html#cfn-quicksight-analysis-transposedtableoption-columnindex
	//
	ColumnIndex *float64 `field:"optional" json:"columnIndex" yaml:"columnIndex"`
	// The width of a column in a transposed table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-transposedtableoption.html#cfn-quicksight-analysis-transposedtableoption-columnwidth
	//
	ColumnWidth *string `field:"optional" json:"columnWidth" yaml:"columnWidth"`
}

