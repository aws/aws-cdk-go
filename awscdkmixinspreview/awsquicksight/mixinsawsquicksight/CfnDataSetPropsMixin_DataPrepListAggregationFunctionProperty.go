package mixinsawsquicksight


// An aggregation function that concatenates values from multiple rows into a single string with a specified separator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPrepListAggregationFunctionProperty := &DataPrepListAggregationFunctionProperty{
//   	Distinct: jsii.Boolean(false),
//   	InputColumnName: jsii.String("inputColumnName"),
//   	Separator: jsii.String("separator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html
//
type CfnDataSetPropsMixin_DataPrepListAggregationFunctionProperty struct {
	// Whether to include only distinct values in the concatenated result, removing duplicates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html#cfn-quicksight-dataset-datapreplistaggregationfunction-distinct
	//
	// Default: - false.
	//
	Distinct interface{} `field:"optional" json:"distinct" yaml:"distinct"`
	// The name of the column containing values to be concatenated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html#cfn-quicksight-dataset-datapreplistaggregationfunction-inputcolumnname
	//
	InputColumnName *string `field:"optional" json:"inputColumnName" yaml:"inputColumnName"`
	// The string used to separate values in the concatenated result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html#cfn-quicksight-dataset-datapreplistaggregationfunction-separator
	//
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

