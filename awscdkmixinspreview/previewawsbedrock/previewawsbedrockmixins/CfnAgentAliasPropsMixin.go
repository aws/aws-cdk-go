package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an agent alias as a resource in a top-level template. Minimally, you must specify the following properties:.
//
// - AgentAliasName – Specify a name for the alias.
//
// For more information about creating aliases for an agent in Amazon Bedrock , see [Deploy an Amazon Bedrock agent](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-deploy.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasPropsMixin := awscdkmixinspreview.Mixins.NewCfnAgentAliasPropsMixin(&CfnAgentAliasMixinProps{
//   	AgentAliasName: jsii.String("agentAliasName"),
//   	AgentId: jsii.String("agentId"),
//   	Description: jsii.String("description"),
//   	RoutingConfiguration: []interface{}{
//   		&AgentAliasRoutingConfigurationListItemProperty{
//   			AgentVersion: jsii.String("agentVersion"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agentalias.html
//
type CfnAgentAliasPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAgentAliasMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentAliasPropsMixin
type jsiiProxy_CfnAgentAliasPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAgentAliasPropsMixin) Props() *CfnAgentAliasMixinProps {
	var returns *CfnAgentAliasMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentAliasPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::AgentAlias`.
func NewCfnAgentAliasPropsMixin(props *CfnAgentAliasMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAgentAliasPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentAliasPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentAliasPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::AgentAlias`.
func NewCfnAgentAliasPropsMixin_Override(c CfnAgentAliasPropsMixin, props *CfnAgentAliasMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAgentAliasPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentAliasPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentAliasPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentAliasPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAgentAliasPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

