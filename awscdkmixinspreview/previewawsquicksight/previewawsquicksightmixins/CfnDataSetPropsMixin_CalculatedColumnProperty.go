package previewawsquicksightmixins


// A calculated column for a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   calculatedColumnProperty := &CalculatedColumnProperty{
//   	ColumnId: jsii.String("columnId"),
//   	ColumnName: jsii.String("columnName"),
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-calculatedcolumn.html
//
type CfnDataSetPropsMixin_CalculatedColumnProperty struct {
	// A unique ID to identify a calculated column.
	//
	// During a dataset update, if the column ID of a calculated column matches that of an existing calculated column, Quick Sight preserves the existing calculated column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-calculatedcolumn.html#cfn-quicksight-dataset-calculatedcolumn-columnid
	//
	ColumnId *string `field:"optional" json:"columnId" yaml:"columnId"`
	// Column name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-calculatedcolumn.html#cfn-quicksight-dataset-calculatedcolumn-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// An expression that defines the calculated column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-calculatedcolumn.html#cfn-quicksight-dataset-calculatedcolumn-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

