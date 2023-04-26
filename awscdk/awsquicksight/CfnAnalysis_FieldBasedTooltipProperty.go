package awsquicksight


// The setup for the detailed tooltip.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldBasedTooltipProperty := &FieldBasedTooltipProperty{
//   	AggregationVisibility: jsii.String("aggregationVisibility"),
//   	TooltipFields: []interface{}{
//   		&TooltipItemProperty{
//   			ColumnTooltipItem: &ColumnTooltipItemProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				Aggregation: &AggregationFunctionProperty{
//   					CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   					DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   					NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   				},
//   				Label: jsii.String("label"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			FieldTooltipItem: &FieldTooltipItemProperty{
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				Label: jsii.String("label"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	TooltipTitleType: jsii.String("tooltipTitleType"),
//   }
//
type CfnAnalysis_FieldBasedTooltipProperty struct {
	// The visibility of `Show aggregations` .
	AggregationVisibility *string `field:"optional" json:"aggregationVisibility" yaml:"aggregationVisibility"`
	// The fields configuration in the tooltip.
	TooltipFields interface{} `field:"optional" json:"tooltipFields" yaml:"tooltipFields"`
	// The type for the >tooltip title. Choose one of the following options:.
	//
	// - `NONE` : Doesn't use the primary value as the title.
	// - `PRIMARY_VALUE` : Uses primary value as the title.
	TooltipTitleType *string `field:"optional" json:"tooltipTitleType" yaml:"tooltipTitleType"`
}

