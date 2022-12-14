package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousHyperParameterRangeProperty := &continuousHyperParameterRangeProperty{
//   	maxValue: jsii.Number(123),
//   	minValue: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnSolution_ContinuousHyperParameterRangeProperty struct {
	// `CfnSolution.ContinuousHyperParameterRangeProperty.MaxValue`.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// `CfnSolution.ContinuousHyperParameterRangeProperty.MinValue`.
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// `CfnSolution.ContinuousHyperParameterRangeProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

