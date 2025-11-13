package awsquicksight


// Defines the type of aggregation function to apply to data during data preparation, supporting simple and list aggregations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPrepAggregationFunctionProperty := &DataPrepAggregationFunctionProperty{
//   	ListAggregation: &DataPrepListAggregationFunctionProperty{
//   		Distinct: jsii.Boolean(false),
//   		Separator: jsii.String("separator"),
//
//   		// the properties below are optional
//   		InputColumnName: jsii.String("inputColumnName"),
//   	},
//   	PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   		PercentileValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		InputColumnName: jsii.String("inputColumnName"),
//   	},
//   	SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   		FunctionType: jsii.String("functionType"),
//
//   		// the properties below are optional
//   		InputColumnName: jsii.String("inputColumnName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepaggregationfunction.html
//
type CfnDataSet_DataPrepAggregationFunctionProperty struct {
	// A list aggregation function that concatenates values from multiple rows into a single delimited string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepaggregationfunction.html#cfn-quicksight-dataset-dataprepaggregationfunction-listaggregation
	//
	ListAggregation interface{} `field:"optional" json:"listAggregation" yaml:"listAggregation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepaggregationfunction.html#cfn-quicksight-dataset-dataprepaggregationfunction-percentileaggregation
	//
	PercentileAggregation interface{} `field:"optional" json:"percentileAggregation" yaml:"percentileAggregation"`
	// A simple aggregation function such as `SUM` , `COUNT` , `AVERAGE` , `MIN` , `MAX` , `MEDIAN` , `VARIANCE` , or `STANDARD_DEVIATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepaggregationfunction.html#cfn-quicksight-dataset-dataprepaggregationfunction-simpleaggregation
	//
	SimpleAggregation interface{} `field:"optional" json:"simpleAggregation" yaml:"simpleAggregation"`
}

