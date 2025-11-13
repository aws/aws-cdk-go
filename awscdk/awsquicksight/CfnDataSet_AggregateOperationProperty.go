package awsquicksight


// A transform operation that groups rows by specified columns and applies aggregation functions to calculate summary values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregateOperationProperty := &AggregateOperationProperty{
//   	Aggregations: []interface{}{
//   		&AggregationProperty{
//   			AggregationFunction: &DataPrepAggregationFunctionProperty{
//   				ListAggregation: &DataPrepListAggregationFunctionProperty{
//   					Distinct: jsii.Boolean(false),
//   					Separator: jsii.String("separator"),
//
//   					// the properties below are optional
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   				PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   					PercentileValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   				SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   					FunctionType: jsii.String("functionType"),
//
//   					// the properties below are optional
//   					InputColumnName: jsii.String("inputColumnName"),
//   				},
//   			},
//   			NewColumnId: jsii.String("newColumnId"),
//   			NewColumnName: jsii.String("newColumnName"),
//   		},
//   	},
//   	Alias: jsii.String("alias"),
//   	Source: &TransformOperationSourceProperty{
//   		TransformOperationId: jsii.String("transformOperationId"),
//
//   		// the properties below are optional
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	GroupByColumnNames: []*string{
//   		jsii.String("groupByColumnNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html
//
type CfnDataSet_AggregateOperationProperty struct {
	// The list of aggregation functions to apply to the grouped data, such as `SUM` , `COUNT` , or `AVERAGE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-aggregations
	//
	Aggregations interface{} `field:"required" json:"aggregations" yaml:"aggregations"`
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The source transform operation that provides input data for the aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// The list of column names to group by when performing the aggregation.
	//
	// Rows with the same values in these columns will be grouped together.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-aggregateoperation.html#cfn-quicksight-dataset-aggregateoperation-groupbycolumnnames
	//
	GroupByColumnNames *[]*string `field:"optional" json:"groupByColumnNames" yaml:"groupByColumnNames"`
}

