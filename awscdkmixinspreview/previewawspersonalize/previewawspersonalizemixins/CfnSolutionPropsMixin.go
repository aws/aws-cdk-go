package previewawspersonalizemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspersonalize/previewawspersonalizemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > By default, all new solutions use automatic training.
//
// With automatic training, you incur training costs while your solution is active. To avoid unnecessary costs, when you are finished you can [update the solution](https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateSolution.html) to turn off automatic training. For information about training costs, see [Amazon Personalize pricing](https://docs.aws.amazon.com/https://aws.amazon.com/personalize/pricing/) .
//
// An object that provides information about a solution. A solution includes the custom recipe, customized parameters, and trained models (Solution Versions) that Amazon Personalize uses to generate recommendations.
//
// After you create a solution, you can’t change its configuration. If you need to make changes, you can [clone the solution](https://docs.aws.amazon.com/personalize/latest/dg/cloning-solution.html) with the Amazon Personalize console or create a new one.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var autoMlConfig interface{}
//   var hpoConfig interface{}
//
//   cfnSolutionPropsMixin := awscdkmixinspreview.Mixins.NewCfnSolutionPropsMixin(&CfnSolutionMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-solution.html
//
type CfnSolutionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSolutionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSolutionPropsMixin
type jsiiProxy_CfnSolutionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSolutionPropsMixin) Props() *CfnSolutionMixinProps {
	var returns *CfnSolutionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSolutionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Personalize::Solution`.
func NewCfnSolutionPropsMixin(props *CfnSolutionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSolutionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSolutionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSolutionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_personalize.mixins.CfnSolutionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Personalize::Solution`.
func NewCfnSolutionPropsMixin_Override(c CfnSolutionPropsMixin, props *CfnSolutionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_personalize.mixins.CfnSolutionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSolutionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSolutionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_personalize.mixins.CfnSolutionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSolutionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_personalize.mixins.CfnSolutionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSolutionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSolutionPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

