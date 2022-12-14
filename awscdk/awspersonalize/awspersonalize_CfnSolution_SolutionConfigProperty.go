package awspersonalize


// Describes the configuration properties for the solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   solutionConfigProperty := &solutionConfigProperty{
//   	algorithmHyperParameters: map[string]*string{
//   		"algorithmHyperParametersKey": jsii.String("algorithmHyperParameters"),
//   	},
//   	autoMlConfig: &autoMLConfigProperty{
//   		metricName: jsii.String("metricName"),
//   		recipeList: []*string{
//   			jsii.String("recipeList"),
//   		},
//   	},
//   	eventValueThreshold: jsii.String("eventValueThreshold"),
//   	featureTransformationParameters: map[string]*string{
//   		"featureTransformationParametersKey": jsii.String("featureTransformationParameters"),
//   	},
//   	hpoConfig: &hpoConfigProperty{
//   		algorithmHyperParameterRanges: &algorithmHyperParameterRangesProperty{
//   			categoricalHyperParameterRanges: []interface{}{
//   				&categoricalHyperParameterRangeProperty{
//   					name: jsii.String("name"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			continuousHyperParameterRanges: []interface{}{
//   				&continuousHyperParameterRangeProperty{
//   					maxValue: jsii.Number(123),
//   					minValue: jsii.Number(123),
//   					name: jsii.String("name"),
//   				},
//   			},
//   			integerHyperParameterRanges: []interface{}{
//   				&integerHyperParameterRangeProperty{
//   					maxValue: jsii.Number(123),
//   					minValue: jsii.Number(123),
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		hpoObjective: &hpoObjectiveProperty{
//   			metricName: jsii.String("metricName"),
//   			metricRegex: jsii.String("metricRegex"),
//   			type: jsii.String("type"),
//   		},
//   		hpoResourceConfig: &hpoResourceConfigProperty{
//   			maxNumberOfTrainingJobs: jsii.String("maxNumberOfTrainingJobs"),
//   			maxParallelTrainingJobs: jsii.String("maxParallelTrainingJobs"),
//   		},
//   	},
//   }
//
type CfnSolution_SolutionConfigProperty struct {
	// Lists the hyperparameter names and ranges.
	AlgorithmHyperParameters interface{} `field:"optional" json:"algorithmHyperParameters" yaml:"algorithmHyperParameters"`
	// The [AutoMLConfig](https://docs.aws.amazon.com/personalize/latest/dg/API_AutoMLConfig.html) object containing a list of recipes to search when AutoML is performed.
	AutoMlConfig interface{} `field:"optional" json:"autoMlConfig" yaml:"autoMlConfig"`
	// Only events with a value greater than or equal to this threshold are used for training a model.
	EventValueThreshold *string `field:"optional" json:"eventValueThreshold" yaml:"eventValueThreshold"`
	// Lists the feature transformation parameters.
	FeatureTransformationParameters interface{} `field:"optional" json:"featureTransformationParameters" yaml:"featureTransformationParameters"`
	// Describes the properties for hyperparameter optimization (HPO).
	HpoConfig interface{} `field:"optional" json:"hpoConfig" yaml:"hpoConfig"`
}

