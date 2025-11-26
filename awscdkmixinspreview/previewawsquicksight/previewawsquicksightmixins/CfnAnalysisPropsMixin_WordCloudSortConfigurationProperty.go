package previewawsquicksightmixins


// The sort configuration of a word cloud visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   wordCloudSortConfigurationProperty := &WordCloudSortConfigurationProperty{
//   	CategoryItemsLimit: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	CategorySort: []interface{}{
//   		&FieldSortOptionsProperty{
//   			ColumnSort: &ColumnSortProperty{
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
//   				Direction: jsii.String("direction"),
//   				SortBy: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-wordcloudsortconfiguration.html
//
type CfnAnalysisPropsMixin_WordCloudSortConfigurationProperty struct {
	// The limit on the number of groups that are displayed in a word cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-wordcloudsortconfiguration.html#cfn-quicksight-analysis-wordcloudsortconfiguration-categoryitemslimit
	//
	CategoryItemsLimit interface{} `field:"optional" json:"categoryItemsLimit" yaml:"categoryItemsLimit"`
	// The sort configuration of group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-wordcloudsortconfiguration.html#cfn-quicksight-analysis-wordcloudsortconfiguration-categorysort
	//
	CategorySort interface{} `field:"optional" json:"categorySort" yaml:"categorySort"`
}

