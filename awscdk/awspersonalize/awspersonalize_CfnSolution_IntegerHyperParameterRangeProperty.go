package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerHyperParameterRangeProperty := &integerHyperParameterRangeProperty{
//   	maxValue: jsii.Number(123),
//   	minValue: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnSolution_IntegerHyperParameterRangeProperty struct {
	// `CfnSolution.IntegerHyperParameterRangeProperty.MaxValue`.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// `CfnSolution.IntegerHyperParameterRangeProperty.MinValue`.
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// `CfnSolution.IntegerHyperParameterRangeProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

