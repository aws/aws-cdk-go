package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericalAggregationFunctionProperty := &NumericalAggregationFunctionProperty{
//   	PercentileAggregation: &PercentileAggregationProperty{
//   		PercentileValue: jsii.Number(123),
//   	},
//   	SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   }
//
type CfnTemplate_NumericalAggregationFunctionProperty struct {
	// `CfnTemplate.NumericalAggregationFunctionProperty.PercentileAggregation`.
	PercentileAggregation interface{} `field:"optional" json:"percentileAggregation" yaml:"percentileAggregation"`
	// `CfnTemplate.NumericalAggregationFunctionProperty.SimpleNumericalAggregation`.
	SimpleNumericalAggregation *string `field:"optional" json:"simpleNumericalAggregation" yaml:"simpleNumericalAggregation"`
}

