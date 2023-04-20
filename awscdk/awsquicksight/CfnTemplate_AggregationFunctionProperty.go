package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationFunctionProperty := &AggregationFunctionProperty{
//   	CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   	DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   	NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   		PercentileAggregation: &PercentileAggregationProperty{
//   			PercentileValue: jsii.Number(123),
//   		},
//   		SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   	},
//   }
//
type CfnTemplate_AggregationFunctionProperty struct {
	// `CfnTemplate.AggregationFunctionProperty.CategoricalAggregationFunction`.
	CategoricalAggregationFunction *string `field:"optional" json:"categoricalAggregationFunction" yaml:"categoricalAggregationFunction"`
	// `CfnTemplate.AggregationFunctionProperty.DateAggregationFunction`.
	DateAggregationFunction *string `field:"optional" json:"dateAggregationFunction" yaml:"dateAggregationFunction"`
	// `CfnTemplate.AggregationFunctionProperty.NumericalAggregationFunction`.
	NumericalAggregationFunction interface{} `field:"optional" json:"numericalAggregationFunction" yaml:"numericalAggregationFunction"`
}

