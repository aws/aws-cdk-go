package mixinsawsquicksight


// Configuration for how to handle value columns in pivot operations, including aggregation settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   valueColumnConfigurationProperty := &ValueColumnConfigurationProperty{
//   	AggregationFunction: &DataPrepAggregationFunctionProperty{
//   		ListAggregation: &DataPrepListAggregationFunctionProperty{
//   			Distinct: jsii.Boolean(false),
//   			InputColumnName: jsii.String("inputColumnName"),
//   			Separator: jsii.String("separator"),
//   		},
//   		PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   			InputColumnName: jsii.String("inputColumnName"),
//   			PercentileValue: jsii.Number(123),
//   		},
//   		SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   			FunctionType: jsii.String("functionType"),
//   			InputColumnName: jsii.String("inputColumnName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-valuecolumnconfiguration.html
//
type CfnDataSetPropsMixin_ValueColumnConfigurationProperty struct {
	// The aggregation function to apply when multiple values map to the same pivoted cell.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-valuecolumnconfiguration.html#cfn-quicksight-dataset-valuecolumnconfiguration-aggregationfunction
	//
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
}

