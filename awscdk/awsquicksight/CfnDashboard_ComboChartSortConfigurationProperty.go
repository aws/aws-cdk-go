package awsquicksight


// The sort configuration of a `ComboChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   comboChartSortConfigurationProperty := &ComboChartSortConfigurationProperty{
//   	CategoryItemsLimit: &ItemsLimitConfigurationProperty{
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
//   	ColorItemsLimit: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	ColorSort: []interface{}{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartsortconfiguration.html
//
type CfnDashboard_ComboChartSortConfigurationProperty struct {
	// The item limit configuration for the category field well of a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartsortconfiguration.html#cfn-quicksight-dashboard-combochartsortconfiguration-categoryitemslimit
	//
	CategoryItemsLimit interface{} `field:"optional" json:"categoryItemsLimit" yaml:"categoryItemsLimit"`
	// The sort configuration of the category field well in a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartsortconfiguration.html#cfn-quicksight-dashboard-combochartsortconfiguration-categorysort
	//
	CategorySort interface{} `field:"optional" json:"categorySort" yaml:"categorySort"`
	// The item limit configuration of the color field well in a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartsortconfiguration.html#cfn-quicksight-dashboard-combochartsortconfiguration-coloritemslimit
	//
	ColorItemsLimit interface{} `field:"optional" json:"colorItemsLimit" yaml:"colorItemsLimit"`
	// The sort configuration of the color field well in a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartsortconfiguration.html#cfn-quicksight-dashboard-combochartsortconfiguration-colorsort
	//
	ColorSort interface{} `field:"optional" json:"colorSort" yaml:"colorSort"`
}

