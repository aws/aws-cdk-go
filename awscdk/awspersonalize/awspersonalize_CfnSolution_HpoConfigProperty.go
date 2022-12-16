package awspersonalize


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hpoConfigProperty := &hpoConfigProperty{
//   	algorithmHyperParameterRanges: &algorithmHyperParameterRangesProperty{
//   		categoricalHyperParameterRanges: []interface{}{
//   			&categoricalHyperParameterRangeProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		continuousHyperParameterRanges: []interface{}{
//   			&continuousHyperParameterRangeProperty{
//   				maxValue: jsii.Number(123),
//   				minValue: jsii.Number(123),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		integerHyperParameterRanges: []interface{}{
//   			&integerHyperParameterRangeProperty{
//   				maxValue: jsii.Number(123),
//   				minValue: jsii.Number(123),
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	hpoObjective: &hpoObjectiveProperty{
//   		metricName: jsii.String("metricName"),
//   		metricRegex: jsii.String("metricRegex"),
//   		type: jsii.String("type"),
//   	},
//   	hpoResourceConfig: &hpoResourceConfigProperty{
//   		maxNumberOfTrainingJobs: jsii.String("maxNumberOfTrainingJobs"),
//   		maxParallelTrainingJobs: jsii.String("maxParallelTrainingJobs"),
//   	},
//   }
//
type CfnSolution_HpoConfigProperty struct {
	// `CfnSolution.HpoConfigProperty.AlgorithmHyperParameterRanges`.
	AlgorithmHyperParameterRanges interface{} `field:"optional" json:"algorithmHyperParameterRanges" yaml:"algorithmHyperParameterRanges"`
	// `CfnSolution.HpoConfigProperty.HpoObjective`.
	HpoObjective interface{} `field:"optional" json:"hpoObjective" yaml:"hpoObjective"`
	// `CfnSolution.HpoConfigProperty.HpoResourceConfig`.
	HpoResourceConfig interface{} `field:"optional" json:"hpoResourceConfig" yaml:"hpoResourceConfig"`
}

