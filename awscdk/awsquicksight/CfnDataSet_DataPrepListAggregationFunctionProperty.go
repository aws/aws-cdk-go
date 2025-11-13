package awsquicksight


// An aggregation function that concatenates values from multiple rows into a single string with a specified separator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPrepListAggregationFunctionProperty := &DataPrepListAggregationFunctionProperty{
//   	Distinct: jsii.Boolean(false),
//   	Separator: jsii.String("separator"),
//
//   	// the properties below are optional
//   	InputColumnName: jsii.String("inputColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html
//
type CfnDataSet_DataPrepListAggregationFunctionProperty struct {
	// Whether to include only distinct values in the concatenated result, removing duplicates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html#cfn-quicksight-dataset-datapreplistaggregationfunction-distinct
	//
	// Default: - false.
	//
	Distinct interface{} `field:"required" json:"distinct" yaml:"distinct"`
	// The string used to separate values in the concatenated result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html#cfn-quicksight-dataset-datapreplistaggregationfunction-separator
	//
	Separator *string `field:"required" json:"separator" yaml:"separator"`
	// The name of the column containing values to be concatenated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreplistaggregationfunction.html#cfn-quicksight-dataset-datapreplistaggregationfunction-inputcolumnname
	//
	InputColumnName *string `field:"optional" json:"inputColumnName" yaml:"inputColumnName"`
}

