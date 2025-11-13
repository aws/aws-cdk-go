package awsquicksight


// A transform operation that pivots data by converting row values into columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotOperationProperty := &PivotOperationProperty{
//   	Alias: jsii.String("alias"),
//   	PivotConfiguration: &PivotConfigurationProperty{
//   		PivotedLabels: []interface{}{
//   			&PivotedLabelProperty{
//   				LabelName: jsii.String("labelName"),
//   				NewColumnId: jsii.String("newColumnId"),
//   				NewColumnName: jsii.String("newColumnName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		LabelColumnName: jsii.String("labelColumnName"),
//   	},
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
//   	ValueColumnConfiguration: &ValueColumnConfigurationProperty{
//   		AggregationFunction: &DataPrepAggregationFunctionProperty{
//   			ListAggregation: &DataPrepListAggregationFunctionProperty{
//   				Distinct: jsii.Boolean(false),
//   				Separator: jsii.String("separator"),
//
//   				// the properties below are optional
//   				InputColumnName: jsii.String("inputColumnName"),
//   			},
//   			PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   				PercentileValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				InputColumnName: jsii.String("inputColumnName"),
//   			},
//   			SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   				FunctionType: jsii.String("functionType"),
//
//   				// the properties below are optional
//   				InputColumnName: jsii.String("inputColumnName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html
//
type CfnDataSet_PivotOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Configuration that specifies which labels to pivot and how to structure the resulting columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-pivotconfiguration
	//
	PivotConfiguration interface{} `field:"required" json:"pivotConfiguration" yaml:"pivotConfiguration"`
	// The source transform operation that provides input data for pivoting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Configuration for how to aggregate values when multiple rows map to the same pivoted column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-valuecolumnconfiguration
	//
	ValueColumnConfiguration interface{} `field:"required" json:"valueColumnConfiguration" yaml:"valueColumnConfiguration"`
	// The list of column names to group by when performing the pivot operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-groupbycolumnnames
	//
	GroupByColumnNames *[]*string `field:"optional" json:"groupByColumnNames" yaml:"groupByColumnNames"`
}

