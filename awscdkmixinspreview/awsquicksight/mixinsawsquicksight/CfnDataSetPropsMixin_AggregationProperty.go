package mixinsawsquicksight


// Defines an aggregation function to be applied to grouped data, creating a new column with the calculated result.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aggregationProperty := &AggregationProperty{
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
//   	NewColumnId: jsii.String("newColumnId"),
//   	NewColumnName: jsii.String("newColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregation.html
//
type CfnDataSetPropsMixin_AggregationProperty struct {
	// The aggregation function to apply, such as `SUM` , `COUNT` , `AVERAGE` , `MIN` , `MAX`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregation.html#cfn-quicksight-dataset-aggregation-aggregationfunction
	//
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// A unique identifier for the new column that will contain the aggregated values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregation.html#cfn-quicksight-dataset-aggregation-newcolumnid
	//
	NewColumnId *string `field:"optional" json:"newColumnId" yaml:"newColumnId"`
	// The name for the new column that will contain the aggregated values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregation.html#cfn-quicksight-dataset-aggregation-newcolumnname
	//
	NewColumnName *string `field:"optional" json:"newColumnName" yaml:"newColumnName"`
}

