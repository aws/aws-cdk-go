package awsquicksight


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
type CfnDashboard_FieldBasedTooltipProperty struct {
	// `CfnDashboard.FieldBasedTooltipProperty.AggregationVisibility`.
	AggregationVisibility *string `field:"optional" json:"aggregationVisibility" yaml:"aggregationVisibility"`
	// `CfnDashboard.FieldBasedTooltipProperty.TooltipFields`.
	TooltipFields interface{} `field:"optional" json:"tooltipFields" yaml:"tooltipFields"`
	// `CfnDashboard.FieldBasedTooltipProperty.TooltipTitleType`.
	TooltipTitleType *string `field:"optional" json:"tooltipTitleType" yaml:"tooltipTitleType"`
}

