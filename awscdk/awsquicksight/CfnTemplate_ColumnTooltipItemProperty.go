package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnTooltipItemProperty := &ColumnTooltipItemProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	Aggregation: &AggregationFunctionProperty{
//   		CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   		DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   		NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   			PercentileAggregation: &PercentileAggregationProperty{
//   				PercentileValue: jsii.Number(123),
//   			},
//   			SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   		},
//   	},
//   	Label: jsii.String("label"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnTemplate_ColumnTooltipItemProperty struct {
	// `CfnTemplate.ColumnTooltipItemProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnTemplate.ColumnTooltipItemProperty.Aggregation`.
	Aggregation interface{} `field:"optional" json:"aggregation" yaml:"aggregation"`
	// `CfnTemplate.ColumnTooltipItemProperty.Label`.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// `CfnTemplate.ColumnTooltipItemProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

