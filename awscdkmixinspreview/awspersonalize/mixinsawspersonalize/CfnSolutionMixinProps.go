package mixinsawspersonalize


// Properties for CfnSolutionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var autoMlConfig interface{}
//   var hpoConfig interface{}
//
//   cfnSolutionMixinProps := &CfnSolutionMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	EventType: jsii.String("eventType"),
//   	Name: jsii.String("name"),
//   	PerformAutoMl: jsii.Boolean(false),
//   	PerformHpo: jsii.Boolean(false),
//   	RecipeArn: jsii.String("recipeArn"),
//   	SolutionConfig: &SolutionConfigProperty{
//   		AlgorithmHyperParameters: map[string]*string{
//   			"algorithmHyperParametersKey": jsii.String("algorithmHyperParameters"),
//   		},
//   		AutoMlConfig: autoMlConfig,
//   		EventValueThreshold: jsii.String("eventValueThreshold"),
//   		FeatureTransformationParameters: map[string]*string{
//   			"featureTransformationParametersKey": jsii.String("featureTransformationParameters"),
//   		},
//   		HpoConfig: hpoConfig,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html
//
type CfnSolutionMixinProps struct {
	// The Amazon Resource Name (ARN) of the dataset group that provides the training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html#cfn-personalize-solution-datasetgrouparn
	//
	DatasetGroupArn *string `field:"optional" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The event type (for example, 'click' or 'like') that is used for training the model.
	//
	// If no `eventType` is provided, Amazon Personalize uses all interactions for training with equal weight regardless of type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html#cfn-personalize-solution-eventtype
	//
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The name of the solution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html#cfn-personalize-solution-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// > We don't recommend enabling automated machine learning.
	//
	// Instead, match your use case to the available Amazon Personalize recipes. For more information, see [Determining your use case.](https://docs.aws.amazon.com/personalize/latest/dg/determining-use-case.html)
	//
	// When true, Amazon Personalize performs a search for the best USER_PERSONALIZATION recipe from the list specified in the solution configuration ( `recipeArn` must not be specified). When false (the default), Amazon Personalize uses `recipeArn` for training.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html#cfn-personalize-solution-performautoml
	//
	PerformAutoMl interface{} `field:"optional" json:"performAutoMl" yaml:"performAutoMl"`
	// Whether to perform hyperparameter optimization (HPO) on the chosen recipe.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html#cfn-personalize-solution-performhpo
	//
	PerformHpo interface{} `field:"optional" json:"performHpo" yaml:"performHpo"`
	// The ARN of the recipe used to create the solution.
	//
	// This is required when `performAutoML` is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html#cfn-personalize-solution-recipearn
	//
	RecipeArn *string `field:"optional" json:"recipeArn" yaml:"recipeArn"`
	// Describes the configuration properties for the solution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html#cfn-personalize-solution-solutionconfig
	//
	SolutionConfig interface{} `field:"optional" json:"solutionConfig" yaml:"solutionConfig"`
}

