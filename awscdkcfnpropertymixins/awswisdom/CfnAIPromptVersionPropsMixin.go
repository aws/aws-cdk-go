package awswisdom

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Q in Connect AI Prompt version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAIPromptVersionPropsMixin := awscdkcfnpropertymixins.Aws_wisdom.NewCfnAIPromptVersionPropsMixin(&CfnAIPromptVersionMixinProps{
//   	AiPromptId: jsii.String("aiPromptId"),
//   	AssistantId: jsii.String("assistantId"),
//   	ModifiedTimeSeconds: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-aipromptversion.html
//
type CfnAIPromptVersionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAIPromptVersionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAIPromptVersionPropsMixin
type jsiiProxy_CfnAIPromptVersionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAIPromptVersionPropsMixin) Props() *CfnAIPromptVersionMixinProps {
	var returns *CfnAIPromptVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAIPromptVersionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::AIPromptVersion`.
func NewCfnAIPromptVersionPropsMixin(props *CfnAIPromptVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAIPromptVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAIPromptVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAIPromptVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::AIPromptVersion`.
func NewCfnAIPromptVersionPropsMixin_Override(c CfnAIPromptVersionPropsMixin, props *CfnAIPromptVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAIPromptVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAIPromptVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAIPromptVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wisdom.CfnAIPromptVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAIPromptVersionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAIPromptVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

