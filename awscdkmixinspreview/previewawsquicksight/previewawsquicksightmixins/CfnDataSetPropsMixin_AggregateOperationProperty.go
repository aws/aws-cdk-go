package previewawsquicksightmixins


// A transform operation that groups rows by specified columns and applies aggregation functions to calculate summary values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aggregateOperationProperty := &AggregateOperationProperty{
//   	Aggregations: []interface{}{
//   		&AggregationProperty{
//   			AggregationFunction: &DataPrepAggregationFunctionProperty{
//   				ListAggregation: &DataPrepListAggregationFunctionProperty{
//   					Distinct: jsii.Boolean(false),
//   					InputColumnName: jsii.String("inputColumnName"),
//   					Separator: jsii.String("separator"),
//   				},
//   				PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   					InputColumnName: jsii.String("inputColumnName"),
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   					FunctionType: jsii.String("functionType"),
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   			},
//   			NewColumnId: jsii.String("newColumnId"),
//   			NewColumnName: jsii.String("newColumnName"),
//   		},
//   	},
//   	Alias: jsii.String("alias"),
//   	GroupByColumnNames: []*string{
//   		jsii.String("groupByColumnNames"),
//   	},
//   	Source: &TransformOperationSourceProperty{
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   		TransformOperationId: jsii.String("transformOperationId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html
//
type CfnDataSetPropsMixin_AggregateOperationProperty struct {
	// The list of aggregation functions to apply to the grouped data, such as `SUM` , `COUNT` , or `AVERAGE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-aggregations
	//
	Aggregations interface{} `field:"optional" json:"aggregations" yaml:"aggregations"`
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The list of column names to group by when performing the aggregation.
	//
	// Rows with the same values in these columns will be grouped together.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-groupbycolumnnames
	//
	GroupByColumnNames *[]*string `field:"optional" json:"groupByColumnNames" yaml:"groupByColumnNames"`
	// The source transform operation that provides input data for the aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

