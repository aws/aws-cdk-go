package awspersonalize


// Specifies the hyperparameters and their ranges.
//
// Hyperparameters can be categorical, continuous, or integer-valued.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-algorithmhyperparameterranges.html
//
type CfnSolution_AlgorithmHyperParameterRangesProperty struct {
	// Provides the name and range of a categorical hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-algorithmhyperparameterranges.html#cfn-personalize-solution-algorithmhyperparameterranges-categoricalhyperparameterranges
	//
	CategoricalHyperParameterRanges interface{} `field:"optional" json:"categoricalHyperParameterRanges" yaml:"categoricalHyperParameterRanges"`
	// Provides the name and range of a continuous hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-algorithmhyperparameterranges.html#cfn-personalize-solution-algorithmhyperparameterranges-continuoushyperparameterranges
	//
	ContinuousHyperParameterRanges interface{} `field:"optional" json:"continuousHyperParameterRanges" yaml:"continuousHyperParameterRanges"`
	// Provides the name and range of an integer-valued hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-algorithmhyperparameterranges.html#cfn-personalize-solution-algorithmhyperparameterranges-integerhyperparameterranges
	//
	IntegerHyperParameterRanges interface{} `field:"optional" json:"integerHyperParameterRanges" yaml:"integerHyperParameterRanges"`
}

