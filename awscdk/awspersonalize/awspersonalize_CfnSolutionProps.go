package awspersonalize


// Properties for defining a `CfnSolution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSolutionProps := &cfnSolutionProps{
//   	datasetGroupArn: jsii.String("datasetGroupArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	eventType: jsii.String("eventType"),
//   	performAutoMl: jsii.Boolean(false),
//   	performHpo: jsii.Boolean(false),
//   	recipeArn: jsii.String("recipeArn"),
//   	solutionConfig: &solutionConfigProperty{
//   		algorithmHyperParameters: map[string]*string{
//   			"algorithmHyperParametersKey": jsii.String("algorithmHyperParameters"),
//   		},
//   		autoMlConfig: &autoMLConfigProperty{
//   			metricName: jsii.String("metricName"),
//   			recipeList: []*string{
//   				jsii.String("recipeList"),
//   			},
//   		},
//   		eventValueThreshold: jsii.String("eventValueThreshold"),
//   		featureTransformationParameters: map[string]*string{
//   			"featureTransformationParametersKey": jsii.String("featureTransformationParameters"),
//   		},
//   		hpoConfig: &hpoConfigProperty{
//   			algorithmHyperParameterRanges: &algorithmHyperParameterRangesProperty{
//   				categoricalHyperParameterRanges: []interface{}{
//   					&categoricalHyperParameterRangeProperty{
//   						name: jsii.String("name"),
//   						values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				continuousHyperParameterRanges: []interface{}{
//   					&continuousHyperParameterRangeProperty{
//   						maxValue: jsii.Number(123),
//   						minValue: jsii.Number(123),
//   						name: jsii.String("name"),
//   					},
//   				},
//   				integerHyperParameterRanges: []interface{}{
//   					&integerHyperParameterRangeProperty{
//   						maxValue: jsii.Number(123),
//   						minValue: jsii.Number(123),
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			hpoObjective: &hpoObjectiveProperty{
//   				metricName: jsii.String("metricName"),
//   				metricRegex: jsii.String("metricRegex"),
//   				type: jsii.String("type"),
//   			},
//   			hpoResourceConfig: &hpoResourceConfigProperty{
//   				maxNumberOfTrainingJobs: jsii.String("maxNumberOfTrainingJobs"),
//   				maxParallelTrainingJobs: jsii.String("maxParallelTrainingJobs"),
//   			},
//   		},
//   	},
//   }
//
type CfnSolutionProps struct {
	// The Amazon Resource Name (ARN) of the dataset group that provides the training data.
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The name of the solution.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The event type (for example, 'click' or 'like') that is used for training the model.
	//
	// If no `eventType` is provided, Amazon Personalize uses all interactions for training with equal weight regardless of type.
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// When true, Amazon Personalize performs a search for the best USER_PERSONALIZATION recipe from the list specified in the solution configuration ( `recipeArn` must not be specified).
	//
	// When false (the default), Amazon Personalize uses `recipeArn` for training.
	PerformAutoMl interface{} `field:"optional" json:"performAutoMl" yaml:"performAutoMl"`
	// Whether to perform hyperparameter optimization (HPO) on the chosen recipe.
	//
	// The default is `false` .
	PerformHpo interface{} `field:"optional" json:"performHpo" yaml:"performHpo"`
	// The ARN of the recipe used to create the solution.
	RecipeArn *string `field:"optional" json:"recipeArn" yaml:"recipeArn"`
	// Describes the configuration properties for the solution.
	SolutionConfig interface{} `field:"optional" json:"solutionConfig" yaml:"solutionConfig"`
}

