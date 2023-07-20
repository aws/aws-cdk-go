package awspersonalize


// Describes the properties for hyperparameter optimization (HPO).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hpoConfigProperty := &HpoConfigProperty{
//   	AlgorithmHyperParameterRanges: &AlgorithmHyperParameterRangesProperty{
//   		CategoricalHyperParameterRanges: []interface{}{
//   			&CategoricalHyperParameterRangeProperty{
//   				Name: jsii.String("name"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		ContinuousHyperParameterRanges: []interface{}{
//   			&ContinuousHyperParameterRangeProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		IntegerHyperParameterRanges: []interface{}{
//   			&IntegerHyperParameterRangeProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	HpoObjective: &HpoObjectiveProperty{
//   		MetricName: jsii.String("metricName"),
//   		MetricRegex: jsii.String("metricRegex"),
//   		Type: jsii.String("type"),
//   	},
//   	HpoResourceConfig: &HpoResourceConfigProperty{
//   		MaxNumberOfTrainingJobs: jsii.String("maxNumberOfTrainingJobs"),
//   		MaxParallelTrainingJobs: jsii.String("maxParallelTrainingJobs"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoconfig.html
//
type CfnSolution_HpoConfigProperty struct {
	// The hyperparameters and their allowable ranges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoconfig.html#cfn-personalize-solution-hpoconfig-algorithmhyperparameterranges
	//
	AlgorithmHyperParameterRanges interface{} `field:"optional" json:"algorithmHyperParameterRanges" yaml:"algorithmHyperParameterRanges"`
	// The metric to optimize during HPO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoconfig.html#cfn-personalize-solution-hpoconfig-hpoobjective
	//
	HpoObjective interface{} `field:"optional" json:"hpoObjective" yaml:"hpoObjective"`
	// Describes the resource configuration for hyperparameter optimization (HPO).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoconfig.html#cfn-personalize-solution-hpoconfig-hporesourceconfig
	//
	HpoResourceConfig interface{} `field:"optional" json:"hpoResourceConfig" yaml:"hpoResourceConfig"`
}

