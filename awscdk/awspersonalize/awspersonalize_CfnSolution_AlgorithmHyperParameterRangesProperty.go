package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   algorithmHyperParameterRangesProperty := &algorithmHyperParameterRangesProperty{
//   	categoricalHyperParameterRanges: []interface{}{
//   		&categoricalHyperParameterRangeProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	continuousHyperParameterRanges: []interface{}{
//   		&continuousHyperParameterRangeProperty{
//   			maxValue: jsii.Number(123),
//   			minValue: jsii.Number(123),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	integerHyperParameterRanges: []interface{}{
//   		&integerHyperParameterRangeProperty{
//   			maxValue: jsii.Number(123),
//   			minValue: jsii.Number(123),
//   			name: jsii.String("name"),
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

