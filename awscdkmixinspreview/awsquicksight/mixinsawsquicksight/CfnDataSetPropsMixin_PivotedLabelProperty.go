package mixinsawsquicksight


// Specifies a label value to be pivoted into a separate column, including the new column name and identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotedLabelProperty := &PivotedLabelProperty{
//   	LabelName: jsii.String("labelName"),
//   	NewColumnId: jsii.String("newColumnId"),
//   	NewColumnName: jsii.String("newColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html
//
type CfnDataSetPropsMixin_PivotedLabelProperty struct {
	// The label value from the source data to be pivoted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html#cfn-quicksight-dataset-pivotedlabel-labelname
	//
	LabelName *string `field:"optional" json:"labelName" yaml:"labelName"`
	// A unique identifier for the new column created from this pivoted label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html#cfn-quicksight-dataset-pivotedlabel-newcolumnid
	//
	NewColumnId *string `field:"optional" json:"newColumnId" yaml:"newColumnId"`
	// The name for the new column created from this pivoted label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html#cfn-quicksight-dataset-pivotedlabel-newcolumnname
	//
	NewColumnName *string `field:"optional" json:"newColumnName" yaml:"newColumnName"`
}

