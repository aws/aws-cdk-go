package mixinsawsquicksight


// Configuration for a pivot operation, specifying which column contains labels and how to pivot them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotConfigurationProperty := &PivotConfigurationProperty{
//   	LabelColumnName: jsii.String("labelColumnName"),
//   	PivotedLabels: []interface{}{
//   		&PivotedLabelProperty{
//   			LabelName: jsii.String("labelName"),
//   			NewColumnId: jsii.String("newColumnId"),
//   			NewColumnName: jsii.String("newColumnName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotconfiguration.html
//
type CfnDataSetPropsMixin_PivotConfigurationProperty struct {
	// The name of the column that contains the labels to be pivoted into separate columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotconfiguration.html#cfn-quicksight-dataset-pivotconfiguration-labelcolumnname
	//
	LabelColumnName *string `field:"optional" json:"labelColumnName" yaml:"labelColumnName"`
	// The list of specific label values to pivot into separate columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotconfiguration.html#cfn-quicksight-dataset-pivotconfiguration-pivotedlabels
	//
	PivotedLabels interface{} `field:"optional" json:"pivotedLabels" yaml:"pivotedLabels"`
}

