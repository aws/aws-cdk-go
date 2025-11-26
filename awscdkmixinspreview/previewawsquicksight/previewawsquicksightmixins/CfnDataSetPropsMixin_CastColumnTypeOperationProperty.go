package previewawsquicksightmixins


// A transform operation that casts a column to a different type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   castColumnTypeOperationProperty := &CastColumnTypeOperationProperty{
//   	ColumnName: jsii.String("columnName"),
//   	Format: jsii.String("format"),
//   	NewColumnType: jsii.String("newColumnType"),
//   	SubType: jsii.String("subType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html
//
type CfnDataSetPropsMixin_CastColumnTypeOperationProperty struct {
	// Column name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html#cfn-quicksight-dataset-castcolumntypeoperation-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// When casting a column from string to datetime type, you can supply a string in a format supported by Quick Sight to denote the source data format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html#cfn-quicksight-dataset-castcolumntypeoperation-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// New column data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html#cfn-quicksight-dataset-castcolumntypeoperation-newcolumntype
	//
	NewColumnType *string `field:"optional" json:"newColumnType" yaml:"newColumnType"`
	// The sub data type of the new column.
	//
	// Sub types are only available for decimal columns that are part of a SPICE dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypeoperation.html#cfn-quicksight-dataset-castcolumntypeoperation-subtype
	//
	SubType *string `field:"optional" json:"subType" yaml:"subType"`
}

