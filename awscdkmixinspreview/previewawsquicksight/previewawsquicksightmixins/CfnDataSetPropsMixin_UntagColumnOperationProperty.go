package previewawsquicksightmixins


// A transform operation that removes tags associated with a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   untagColumnOperationProperty := &UntagColumnOperationProperty{
//   	ColumnName: jsii.String("columnName"),
//   	TagNames: []*string{
//   		jsii.String("tagNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-untagcolumnoperation.html
//
type CfnDataSetPropsMixin_UntagColumnOperationProperty struct {
	// The column that this operation acts on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-untagcolumnoperation.html#cfn-quicksight-dataset-untagcolumnoperation-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// The column tags to remove from this column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-untagcolumnoperation.html#cfn-quicksight-dataset-untagcolumnoperation-tagnames
	//
	TagNames *[]*string `field:"optional" json:"tagNames" yaml:"tagNames"`
}

