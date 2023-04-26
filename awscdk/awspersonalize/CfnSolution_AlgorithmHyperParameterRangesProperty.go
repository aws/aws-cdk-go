package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   algorithmHyperParameterRangesProperty := &AlgorithmHyperParameterRangesProperty{
//   	CategoricalHyperParameterRanges: []interface{}{
//   		&CategoricalHyperParameterRangeProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	ContinuousHyperParameterRanges: []interface{}{
//   		&ContinuousHyperParameterRangeProperty{
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	IntegerHyperParameterRanges: []interface{}{
//   		&IntegerHyperParameterRangeProperty{
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnSolution_AlgorithmHyperParameterRangesProperty struct {
	// `CfnSolution.AlgorithmHyperParameterRangesProperty.CategoricalHyperParameterRanges`.
	CategoricalHyperParameterRanges interface{} `field:"optional" json:"categoricalHyperParameterRanges" yaml:"categoricalHyperParameterRanges"`
	// `CfnSolution.AlgorithmHyperParameterRangesProperty.ContinuousHyperParameterRanges`.
	ContinuousHyperParameterRanges interface{} `field:"optional" json:"continuousHyperParameterRanges" yaml:"continuousHyperParameterRanges"`
	// `CfnSolution.AlgorithmHyperParameterRangesProperty.IntegerHyperParameterRanges`.
	IntegerHyperParameterRanges interface{} `field:"optional" json:"integerHyperParameterRanges" yaml:"integerHyperParameterRanges"`
}

