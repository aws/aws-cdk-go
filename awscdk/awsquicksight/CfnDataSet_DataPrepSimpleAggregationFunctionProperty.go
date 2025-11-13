package awsquicksight


// A simple aggregation function that performs standard statistical operations on a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPrepSimpleAggregationFunctionProperty := &DataPrepSimpleAggregationFunctionProperty{
//   	FunctionType: jsii.String("functionType"),
//
//   	// the properties below are optional
//   	InputColumnName: jsii.String("inputColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepsimpleaggregationfunction.html
//
type CfnDataSet_DataPrepSimpleAggregationFunctionProperty struct {
	// The type of aggregation function to perform, such as `COUNT` , `SUM` , `AVERAGE` , `MIN` , `MAX` , `MEDIAN` , `VARIANCE` , or `STANDARD_DEVIATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepsimpleaggregationfunction.html#cfn-quicksight-dataset-dataprepsimpleaggregationfunction-functiontype
	//
	FunctionType *string `field:"required" json:"functionType" yaml:"functionType"`
	// The name of the column on which to perform the aggregation function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepsimpleaggregationfunction.html#cfn-quicksight-dataset-dataprepsimpleaggregationfunction-inputcolumnname
	//
	InputColumnName *string `field:"optional" json:"inputColumnName" yaml:"inputColumnName"`
}

