package awspersonalize


// Describes the configuration properties for the solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoMlConfig interface{}
//   var hpoConfig interface{}
//
//   solutionConfigProperty := &SolutionConfigProperty{
//   	AlgorithmHyperParameters: map[string]*string{
//   		"algorithmHyperParametersKey": jsii.String("algorithmHyperParameters"),
//   	},
//   	AutoMlConfig: autoMlConfig,
//   	EventValueThreshold: jsii.String("eventValueThreshold"),
//   	FeatureTransformationParameters: map[string]*string{
//   		"featureTransformationParametersKey": jsii.String("featureTransformationParameters"),
//   	},
//   	HpoConfig: hpoConfig,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-solutionconfig.html
//
type CfnSolution_SolutionConfigProperty struct {
	// Lists the algorithm hyperparameters and their values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-solutionconfig.html#cfn-personalize-solution-solutionconfig-algorithmhyperparameters
	//
	AlgorithmHyperParameters interface{} `field:"optional" json:"algorithmHyperParameters" yaml:"algorithmHyperParameters"`
	// The [AutoMLConfig](https://docs.aws.amazon.com/personalize/latest/dg/API_AutoMLConfig.html) object containing a list of recipes to search when AutoML is performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-solutionconfig.html#cfn-personalize-solution-solutionconfig-automlconfig
	//
	AutoMlConfig interface{} `field:"optional" json:"autoMlConfig" yaml:"autoMlConfig"`
	// Only events with a value greater than or equal to this threshold are used for training a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-solutionconfig.html#cfn-personalize-solution-solutionconfig-eventvaluethreshold
	//
	EventValueThreshold *string `field:"optional" json:"eventValueThreshold" yaml:"eventValueThreshold"`
	// Lists the feature transformation parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-solutionconfig.html#cfn-personalize-solution-solutionconfig-featuretransformationparameters
	//
	FeatureTransformationParameters interface{} `field:"optional" json:"featureTransformationParameters" yaml:"featureTransformationParameters"`
	// Describes the properties for hyperparameter optimization (HPO).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-solutionconfig.html#cfn-personalize-solution-solutionconfig-hpoconfig
	//
	HpoConfig interface{} `field:"optional" json:"hpoConfig" yaml:"hpoConfig"`
}

