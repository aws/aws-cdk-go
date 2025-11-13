package awsquicksight


// Specifies a label value to be pivoted into a separate column, including the new column name and identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotedLabelProperty := &PivotedLabelProperty{
//   	LabelName: jsii.String("labelName"),
//   	NewColumnId: jsii.String("newColumnId"),
//   	NewColumnName: jsii.String("newColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html
//
type CfnDataSet_PivotedLabelProperty struct {
	// The label value from the source data to be pivoted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html#cfn-quicksight-dataset-pivotedlabel-labelname
	//
	LabelName *string `field:"required" json:"labelName" yaml:"labelName"`
	// A unique identifier for the new column created from this pivoted label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html#cfn-quicksight-dataset-pivotedlabel-newcolumnid
	//
	NewColumnId *string `field:"required" json:"newColumnId" yaml:"newColumnId"`
	// The name for the new column created from this pivoted label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotedlabel.html#cfn-quicksight-dataset-pivotedlabel-newcolumnname
	//
	NewColumnName *string `field:"required" json:"newColumnName" yaml:"newColumnName"`
}

