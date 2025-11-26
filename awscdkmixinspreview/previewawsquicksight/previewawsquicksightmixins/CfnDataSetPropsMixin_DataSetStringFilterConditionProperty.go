package previewawsquicksightmixins


// A filter condition for string columns, supporting both comparison and list-based filtering.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetStringFilterConditionProperty := &DataSetStringFilterConditionProperty{
//   	ColumnName: jsii.String("columnName"),
//   	ComparisonFilterCondition: &DataSetStringComparisonFilterConditionProperty{
//   		Operator: jsii.String("operator"),
//   		Value: &DataSetStringFilterValueProperty{
//   			StaticValue: jsii.String("staticValue"),
//   		},
//   	},
//   	ListFilterCondition: &DataSetStringListFilterConditionProperty{
//   		Operator: jsii.String("operator"),
//   		Values: &DataSetStringListFilterValueProperty{
//   			StaticValues: []*string{
//   				jsii.String("staticValues"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringfiltercondition.html
//
type CfnDataSetPropsMixin_DataSetStringFilterConditionProperty struct {
	// The name of the string column to filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringfiltercondition.html#cfn-quicksight-dataset-datasetstringfiltercondition-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// A comparison-based filter condition for the string column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringfiltercondition.html#cfn-quicksight-dataset-datasetstringfiltercondition-comparisonfiltercondition
	//
	ComparisonFilterCondition interface{} `field:"optional" json:"comparisonFilterCondition" yaml:"comparisonFilterCondition"`
	// A list-based filter condition that includes or excludes values from a specified list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringfiltercondition.html#cfn-quicksight-dataset-datasetstringfiltercondition-listfiltercondition
	//
	ListFilterCondition interface{} `field:"optional" json:"listFilterCondition" yaml:"listFilterCondition"`
}

