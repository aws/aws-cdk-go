package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::QuickSight::Agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAgentPropsMixin := awscdkcfnpropertymixins.Aws_quicksight.NewCfnAgentPropsMixin(&CfnAgentMixinProps{
//   	ActionConnectors: []*string{
//   		jsii.String("actionConnectors"),
//   	},
//   	AgentId: jsii.String("agentId"),
//   	AgentLifecycle: jsii.String("agentLifecycle"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	CustomPromptInput: &CustomPromptInputProperty{
//   		ExistingPrompt: &CustomPromptProfileProperty{
//   			ModelProfileId: jsii.String("modelProfileId"),
//   			QbsAwsAccountId: jsii.String("qbsAwsAccountId"),
//   			SubscriptionId: jsii.String("subscriptionId"),
//   		},
//   		NewPrompt: &CustomPromptInputParametersProperty{
//   			CustomInstructions: jsii.String("customInstructions"),
//   			Identity: jsii.String("identity"),
//   			OutputStyle: jsii.String("outputStyle"),
//   			ResponseLength: jsii.String("responseLength"),
//   			Tone: jsii.String("tone"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IconId: jsii.String("iconId"),
//   	Name: jsii.String("name"),
//   	Spaces: []*string{
//   		jsii.String("spaces"),
//   	},
//   	StarterPrompts: []*string{
//   		jsii.String("starterPrompts"),
//   	},
//   	Tags: []AgentTagProperty{
//   		&AgentTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WelcomeMessage: jsii.String("welcomeMessage"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html
//
type CfnAgentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAgentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentPropsMixin
type jsiiProxy_CfnAgentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAgentPropsMixin) Props() *CfnAgentMixinProps {
	var returns *CfnAgentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::Agent`.
func NewCfnAgentPropsMixin(props *CfnAgentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAgentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAgentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::Agent`.
func NewCfnAgentPropsMixin_Override(c CfnAgentPropsMixin, props *CfnAgentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAgentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAgentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAgentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnAgentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAgentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

