package previewawsquicksightmixins


// The column option of the transposed table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transposedTableOptionProperty := &TransposedTableOptionProperty{
//   	ColumnIndex: jsii.Number(123),
//   	ColumnType: jsii.String("columnType"),
//   	ColumnWidth: jsii.String("columnWidth"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-transposedtableoption.html
//
type CfnTemplatePropsMixin_TransposedTableOptionProperty struct {
	// The index of a columns in a transposed table.
	//
	// The index range is 0-9999.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-transposedtableoption.html#cfn-quicksight-template-transposedtableoption-columnindex
	//
	ColumnIndex *float64 `field:"optional" json:"columnIndex" yaml:"columnIndex"`
	// The column type of the column in a transposed table. Choose one of the following options:.
	//
	// - `ROW_HEADER_COLUMN` : Refers to the leftmost column of the row header in the transposed table.
	// - `VALUE_COLUMN` : Refers to all value columns in the transposed table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-transposedtableoption.html#cfn-quicksight-template-transposedtableoption-columntype
	//
	ColumnType *string `field:"optional" json:"columnType" yaml:"columnType"`
	// The width of a column in a transposed table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-transposedtableoption.html#cfn-quicksight-template-transposedtableoption-columnwidth
	//
	ColumnWidth *string `field:"optional" json:"columnWidth" yaml:"columnWidth"`
}

