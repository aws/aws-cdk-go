package awsquicksight


// The sort configuration of a `RadarChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   radarChartSortConfigurationProperty := &RadarChartSortConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html
//
type CfnTemplate_RadarChartSortConfigurationProperty struct {
	// The category items limit for a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-categoryitemslimit
	//
	CategoryItemsLimit interface{} `field:"optional" json:"categoryItemsLimit" yaml:"categoryItemsLimit"`
	// The category sort options of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-categorysort
	//
	CategorySort interface{} `field:"optional" json:"categorySort" yaml:"categorySort"`
	// The color items limit of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-coloritemslimit
	//
	ColorItemsLimit interface{} `field:"optional" json:"colorItemsLimit" yaml:"colorItemsLimit"`
	// The color sort configuration of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartsortconfiguration.html#cfn-quicksight-template-radarchartsortconfiguration-colorsort
	//
	ColorSort interface{} `field:"optional" json:"colorSort" yaml:"colorSort"`
}

