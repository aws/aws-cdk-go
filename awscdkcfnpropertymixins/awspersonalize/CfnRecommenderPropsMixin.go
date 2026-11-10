package awspersonalize

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A recommender for a Domain dataset group that generates recommendations using a domain use case recipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRecommenderPropsMixin := awscdkcfnpropertymixins.Aws_personalize.NewCfnRecommenderPropsMixin(&CfnRecommenderMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	Name: jsii.String("name"),
//   	RecipeArn: jsii.String("recipeArn"),
//   	RecommenderConfig: &RecommenderConfigProperty{
//   		EnableMetadataWithRecommendations: jsii.Boolean(false),
//   		ItemExplorationConfig: map[string]*string{
//   			"itemExplorationConfigKey": jsii.String("itemExplorationConfig"),
//   		},
//   		MinRecommendationRequestsPerSecond: jsii.Number(123),
//   		TrainingDataConfig: &TrainingDataConfigProperty{
//   			ExcludedDatasetColumns: map[string][]*string{
//   				"excludedDatasetColumnsKey": []*string{
//   					jsii.String("excludedDatasetColumns"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-recommender.html
//
type CfnRecommenderPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRecommenderMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRecommenderPropsMixin
type jsiiProxy_CfnRecommenderPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRecommenderPropsMixin) Props() *CfnRecommenderMixinProps {
	var returns *CfnRecommenderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecommenderPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Personalize::Recommender`.
func NewCfnRecommenderPropsMixin(props *CfnRecommenderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRecommenderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRecommenderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecommenderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnRecommenderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Personalize::Recommender`.
func NewCfnRecommenderPropsMixin_Override(c CfnRecommenderPropsMixin, props *CfnRecommenderMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnRecommenderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRecommenderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecommenderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnRecommenderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecommenderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnRecommenderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecommenderPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRecommenderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

