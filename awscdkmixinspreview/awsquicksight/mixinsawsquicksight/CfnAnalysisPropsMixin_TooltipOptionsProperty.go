package mixinsawsquicksight


// The display options for the visual tooltip.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tooltipOptionsProperty := &TooltipOptionsProperty{
//   	FieldBasedTooltip: &FieldBasedTooltipProperty{
//   		AggregationVisibility: jsii.String("aggregationVisibility"),
//   		TooltipFields: []interface{}{
//   			&TooltipItemProperty{
//   				ColumnTooltipItem: &ColumnTooltipItemProperty{
//   					Aggregation: &AggregationFunctionProperty{
//   						AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   							SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   							ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   						},
//   						CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   						DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   						NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   							PercentileAggregation: &PercentileAggregationProperty{
//   								PercentileValue: jsii.Number(123),
//   							},
//   							SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   						},
//   					},
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					Label: jsii.String("label"),
//   					TooltipTarget: jsii.String("tooltipTarget"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				FieldTooltipItem: &FieldTooltipItemProperty{
//   					FieldId: jsii.String("fieldId"),
//   					Label: jsii.String("label"),
//   					TooltipTarget: jsii.String("tooltipTarget"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   		},
//   		TooltipTitleType: jsii.String("tooltipTitleType"),
//   	},
//   	SelectedTooltipType: jsii.String("selectedTooltipType"),
//   	TooltipVisibility: jsii.String("tooltipVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html
//
type CfnAnalysisPropsMixin_TooltipOptionsProperty struct {
	// The setup for the detailed tooltip.
	//
	// The tooltip setup is always saved. The display type is decided based on the tooltip type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html#cfn-quicksight-analysis-tooltipoptions-fieldbasedtooltip
	//
	FieldBasedTooltip interface{} `field:"optional" json:"fieldBasedTooltip" yaml:"fieldBasedTooltip"`
	// The selected type for the tooltip. Choose one of the following options:.
	//
	// - `BASIC` : A basic tooltip.
	// - `DETAILED` : A detailed tooltip.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html#cfn-quicksight-analysis-tooltipoptions-selectedtooltiptype
	//
	SelectedTooltipType *string `field:"optional" json:"selectedTooltipType" yaml:"selectedTooltipType"`
	// Determines whether or not the tooltip is visible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html#cfn-quicksight-analysis-tooltipoptions-tooltipvisibility
	//
	TooltipVisibility *string `field:"optional" json:"tooltipVisibility" yaml:"tooltipVisibility"`
}

