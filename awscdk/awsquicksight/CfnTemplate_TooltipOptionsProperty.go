package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tooltipOptionsProperty := &TooltipOptionsProperty{
//   	FieldBasedTooltip: &FieldBasedTooltipProperty{
//   		AggregationVisibility: jsii.String("aggregationVisibility"),
//   		TooltipFields: []interface{}{
//   			&TooltipItemProperty{
//   				ColumnTooltipItem: &ColumnTooltipItemProperty{
//   					Column: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//
//   					// the properties below are optional
//   					Aggregation: &AggregationFunctionProperty{
//   						CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   						DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   						NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   							PercentileAggregation: &PercentileAggregationProperty{
//   								PercentileValue: jsii.Number(123),
//   							},
//   							SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   						},
//   					},
//   					Label: jsii.String("label"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				FieldTooltipItem: &FieldTooltipItemProperty{
//   					FieldId: jsii.String("fieldId"),
//
//   					// the properties below are optional
//   					Label: jsii.String("label"),
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
type CfnTemplate_TooltipOptionsProperty struct {
	// `CfnTemplate.TooltipOptionsProperty.FieldBasedTooltip`.
	FieldBasedTooltip interface{} `field:"optional" json:"fieldBasedTooltip" yaml:"fieldBasedTooltip"`
	// `CfnTemplate.TooltipOptionsProperty.SelectedTooltipType`.
	SelectedTooltipType *string `field:"optional" json:"selectedTooltipType" yaml:"selectedTooltipType"`
	// `CfnTemplate.TooltipOptionsProperty.TooltipVisibility`.
	TooltipVisibility *string `field:"optional" json:"tooltipVisibility" yaml:"tooltipVisibility"`
}

