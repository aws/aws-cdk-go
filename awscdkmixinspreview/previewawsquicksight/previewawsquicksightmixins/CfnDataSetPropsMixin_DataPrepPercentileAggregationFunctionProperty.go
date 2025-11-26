package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPrepPercentileAggregationFunctionProperty := &DataPrepPercentileAggregationFunctionProperty{
//   	InputColumnName: jsii.String("inputColumnName"),
//   	PercentileValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreppercentileaggregationfunction.html
//
type CfnDataSetPropsMixin_DataPrepPercentileAggregationFunctionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreppercentileaggregationfunction.html#cfn-quicksight-dataset-datapreppercentileaggregationfunction-inputcolumnname
	//
	InputColumnName *string `field:"optional" json:"inputColumnName" yaml:"inputColumnName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datapreppercentileaggregationfunction.html#cfn-quicksight-dataset-datapreppercentileaggregationfunction-percentilevalue
	//
	// Default: - 0.
	//
	PercentileValue *float64 `field:"optional" json:"percentileValue" yaml:"percentileValue"`
}

