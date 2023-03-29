package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tooltipItemProperty := &TooltipItemProperty{
//   	ColumnTooltipItem: &ColumnTooltipItemProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//
//   		// the properties below are optional
//   		Aggregation: &AggregationFunctionProperty{
//   			CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   			DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   			NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   				PercentileAggregation: &PercentileAggregationProperty{
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   			},
//   		},
//   		Label: jsii.String("label"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	FieldTooltipItem: &FieldTooltipItemProperty{
//   		FieldId: jsii.String("fieldId"),
//
//   		// the properties below are optional
//   		Label: jsii.String("label"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnTemplate_TooltipItemProperty struct {
	// `CfnTemplate.TooltipItemProperty.ColumnTooltipItem`.
	ColumnTooltipItem interface{} `field:"optional" json:"columnTooltipItem" yaml:"columnTooltipItem"`
	// `CfnTemplate.TooltipItemProperty.FieldTooltipItem`.
	FieldTooltipItem interface{} `field:"optional" json:"fieldTooltipItem" yaml:"fieldTooltipItem"`
}

