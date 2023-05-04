package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerHyperParameterRangeProperty := &IntegerHyperParameterRangeProperty{
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//   	Name: jsii.String("name"),
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

