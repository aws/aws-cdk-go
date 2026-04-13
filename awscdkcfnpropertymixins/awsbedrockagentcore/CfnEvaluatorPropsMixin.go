package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::Evaluator - Creates a custom evaluator for agent quality assessment using LLM-as-a-Judge configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnEvaluatorPropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnEvaluatorPropsMixin(&CfnEvaluatorMixinProps{
//   	Description: jsii.String("description"),
//   	EvaluatorConfig: &EvaluatorConfigProperty{
//   		CodeBased: &CodeBasedEvaluatorConfigProperty{
//   			LambdaConfig: &LambdaEvaluatorConfigProperty{
//   				LambdaArn: jsii.String("lambdaArn"),
//   				LambdaTimeoutInSeconds: jsii.Number(123),
//   			},
//   		},
//   		LlmAsAJudge: &LlmAsAJudgeEvaluatorConfigProperty{
//   			Instructions: jsii.String("instructions"),
//   			ModelConfig: &EvaluatorModelConfigProperty{
//   				BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   					AdditionalModelRequestFields: additionalModelRequestFields,
//   					InferenceConfig: &InferenceConfigurationProperty{
//   						MaxTokens: jsii.Number(123),
//   						Temperature: jsii.Number(123),
//   						TopP: jsii.Number(123),
//   					},
//   					ModelId: jsii.String("modelId"),
//   				},
//   			},
//   			RatingScale: &RatingScaleProperty{
//   				Categorical: []interface{}{
//   					&CategoricalScaleDefinitionProperty{
//   						Definition: jsii.String("definition"),
//   						Label: jsii.String("label"),
//   					},
//   				},
//   				Numerical: []interface{}{
//   					&NumericalScaleDefinitionProperty{
//   						Definition: jsii.String("definition"),
//   						Label: jsii.String("label"),
//   						Value: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	EvaluatorName: jsii.String("evaluatorName"),
//   	Level: jsii.String("level"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-evaluator.html
//
type CfnEvaluatorPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEvaluatorMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEvaluatorPropsMixin
type jsiiProxy_CfnEvaluatorPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEvaluatorPropsMixin) Props() *CfnEvaluatorMixinProps {
	var returns *CfnEvaluatorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEvaluatorPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::Evaluator`.
func NewCfnEvaluatorPropsMixin(props *CfnEvaluatorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEvaluatorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEvaluatorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEvaluatorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::Evaluator`.
func NewCfnEvaluatorPropsMixin_Override(c CfnEvaluatorPropsMixin, props *CfnEvaluatorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEvaluatorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEvaluatorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEvaluatorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnEvaluatorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEvaluatorPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEvaluatorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

