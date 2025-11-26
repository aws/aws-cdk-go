package previewawsquicksightmixins


// A simple aggregation function that performs standard statistical operations on a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPrepSimpleAggregationFunctionProperty := &DataPrepSimpleAggregationFunctionProperty{
//   	FunctionType: jsii.String("functionType"),
//   	InputColumnName: jsii.String("inputColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepsimpleaggregationfunction.html
//
type CfnDataSetPropsMixin_DataPrepSimpleAggregationFunctionProperty struct {
	// The type of aggregation function to perform, such as `COUNT` , `SUM` , `AVERAGE` , `MIN` , `MAX` , `MEDIAN` , `VARIANCE` , or `STANDARD_DEVIATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepsimpleaggregationfunction.html#cfn-quicksight-dataset-dataprepsimpleaggregationfunction-functiontype
	//
	FunctionType *string `field:"optional" json:"functionType" yaml:"functionType"`
	// The name of the column on which to perform the aggregation function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-dataprepsimpleaggregationfunction.html#cfn-quicksight-dataset-dataprepsimpleaggregationfunction-inputcolumnname
	//
	InputColumnName *string `field:"optional" json:"inputColumnName" yaml:"inputColumnName"`
}

