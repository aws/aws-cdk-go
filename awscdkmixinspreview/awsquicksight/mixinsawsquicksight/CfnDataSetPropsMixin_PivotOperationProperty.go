package mixinsawsquicksight


// A transform operation that pivots data by converting row values into columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotOperationProperty := &PivotOperationProperty{
//   	Alias: jsii.String("alias"),
//   	GroupByColumnNames: []*string{
//   		jsii.String("groupByColumnNames"),
//   	},
//   	PivotConfiguration: &PivotConfigurationProperty{
//   		LabelColumnName: jsii.String("labelColumnName"),
//   		PivotedLabels: []interface{}{
//   			&PivotedLabelProperty{
//   				LabelName: jsii.String("labelName"),
//   				NewColumnId: jsii.String("newColumnId"),
//   				NewColumnName: jsii.String("newColumnName"),
//   			},
//   		},
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
//   	ValueColumnConfiguration: &ValueColumnConfigurationProperty{
//   		AggregationFunction: &DataPrepAggregationFunctionProperty{
//   			ListAggregation: &DataPrepListAggregationFunctionProperty{
//   				Distinct: jsii.Boolean(false),
//   				InputColumnName: jsii.String("inputColumnName"),
//   				Separator: jsii.String("separator"),
//   			},
//   			PercentileAggregation: &DataPrepPercentileAggregationFunctionProperty{
//   				InputColumnName: jsii.String("inputColumnName"),
//   				PercentileValue: jsii.Number(123),
//   			},
//   			SimpleAggregation: &DataPrepSimpleAggregationFunctionProperty{
//   				FunctionType: jsii.String("functionType"),
//   				InputColumnName: jsii.String("inputColumnName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html
//
type CfnDataSetPropsMixin_PivotOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The list of column names to group by when performing the pivot operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-groupbycolumnnames
	//
	GroupByColumnNames *[]*string `field:"optional" json:"groupByColumnNames" yaml:"groupByColumnNames"`
	// Configuration that specifies which labels to pivot and how to structure the resulting columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-pivotconfiguration
	//
	PivotConfiguration interface{} `field:"optional" json:"pivotConfiguration" yaml:"pivotConfiguration"`
	// The source transform operation that provides input data for pivoting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Configuration for how to aggregate values when multiple rows map to the same pivoted column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-pivotoperation.html#cfn-quicksight-dataset-pivotoperation-valuecolumnconfiguration
	//
	ValueColumnConfiguration interface{} `field:"optional" json:"valueColumnConfiguration" yaml:"valueColumnConfiguration"`
}

