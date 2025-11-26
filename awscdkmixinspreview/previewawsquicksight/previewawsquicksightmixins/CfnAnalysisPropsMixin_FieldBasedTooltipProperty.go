package previewawsquicksightmixins


// The setup for the detailed tooltip.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldBasedTooltipProperty := &FieldBasedTooltipProperty{
//   	AggregationVisibility: jsii.String("aggregationVisibility"),
//   	TooltipFields: []interface{}{
//   		&TooltipItemProperty{
//   			ColumnTooltipItem: &ColumnTooltipItemProperty{
//   				Aggregation: &AggregationFunctionProperty{
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
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				Label: jsii.String("label"),
//   				TooltipTarget: jsii.String("tooltipTarget"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			FieldTooltipItem: &FieldTooltipItemProperty{
//   				FieldId: jsii.String("fieldId"),
//   				Label: jsii.String("label"),
//   				TooltipTarget: jsii.String("tooltipTarget"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	TooltipTitleType: jsii.String("tooltipTitleType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html
//
type CfnAnalysisPropsMixin_FieldBasedTooltipProperty struct {
	// The visibility of `Show aggregations` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html#cfn-quicksight-analysis-fieldbasedtooltip-aggregationvisibility
	//
	AggregationVisibility *string `field:"optional" json:"aggregationVisibility" yaml:"aggregationVisibility"`
	// The fields configuration in the tooltip.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html#cfn-quicksight-analysis-fieldbasedtooltip-tooltipfields
	//
	TooltipFields interface{} `field:"optional" json:"tooltipFields" yaml:"tooltipFields"`
	// The type for the >tooltip title. Choose one of the following options:.
	//
	// - `NONE` : Doesn't use the primary value as the title.
	// - `PRIMARY_VALUE` : Uses primary value as the title.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fieldbasedtooltip.html#cfn-quicksight-analysis-fieldbasedtooltip-tooltiptitletype
	//
	TooltipTitleType *string `field:"optional" json:"tooltipTitleType" yaml:"tooltipTitleType"`
}

