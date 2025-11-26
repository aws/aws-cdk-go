package previewawsquicksightmixins


// The sort configuration of a sankey diagram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sankeyDiagramSortConfigurationProperty := &SankeyDiagramSortConfigurationProperty{
//   	DestinationItemsLimit: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	SourceItemsLimit: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	WeightSort: []interface{}{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramsortconfiguration.html
//
type CfnAnalysisPropsMixin_SankeyDiagramSortConfigurationProperty struct {
	// The limit on the number of destination nodes that are displayed in a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramsortconfiguration.html#cfn-quicksight-analysis-sankeydiagramsortconfiguration-destinationitemslimit
	//
	DestinationItemsLimit interface{} `field:"optional" json:"destinationItemsLimit" yaml:"destinationItemsLimit"`
	// The limit on the number of source nodes that are displayed in a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramsortconfiguration.html#cfn-quicksight-analysis-sankeydiagramsortconfiguration-sourceitemslimit
	//
	SourceItemsLimit interface{} `field:"optional" json:"sourceItemsLimit" yaml:"sourceItemsLimit"`
	// The sort configuration of the weight fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sankeydiagramsortconfiguration.html#cfn-quicksight-analysis-sankeydiagramsortconfiguration-weightsort
	//
	WeightSort interface{} `field:"optional" json:"weightSort" yaml:"weightSort"`
}

