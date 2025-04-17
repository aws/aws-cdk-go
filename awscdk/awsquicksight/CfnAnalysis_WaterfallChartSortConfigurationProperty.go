package awsquicksight


// The sort configuration of a waterfall visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waterfallChartSortConfigurationProperty := &WaterfallChartSortConfigurationProperty{
//   	BreakdownItemsLimit: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	CategorySort: []interface{}{
//   		&FieldSortOptionsProperty{
//   			ColumnSort: &ColumnSortProperty{
//   				Direction: jsii.String("direction"),
//   				SortBy: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				AggregationFunction: &AggregationFunctionProperty{
//   					AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   						SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   						ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   					},
//   					CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   					DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   					NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   				},
//   			},
//   			FieldSort: &FieldSortProperty{
//   				Direction: jsii.String("direction"),
//   				FieldId: jsii.String("fieldId"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartsortconfiguration.html
//
type CfnAnalysis_WaterfallChartSortConfigurationProperty struct {
	// The limit on the number of bar groups that are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartsortconfiguration.html#cfn-quicksight-analysis-waterfallchartsortconfiguration-breakdownitemslimit
	//
	BreakdownItemsLimit interface{} `field:"optional" json:"breakdownItemsLimit" yaml:"breakdownItemsLimit"`
	// The sort configuration of the category fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-waterfallchartsortconfiguration.html#cfn-quicksight-analysis-waterfallchartsortconfiguration-categorysort
	//
	CategorySort interface{} `field:"optional" json:"categorySort" yaml:"categorySort"`
}

