package awswellarchitected

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::WellArchitected::AgentContext.
//
// An Agent Context provides supplemental application context (free-form content) within an Agent Profile, used to improve the quality of Well-Architected Agent recommendations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var content interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnAgentContextPropsMixin := awscdkcfnpropertymixins.Aws_wellarchitected.NewCfnAgentContextPropsMixin(&CfnAgentContextMixinProps{
//   	Content: content,
//   	ContextType: jsii.String("contextType"),
//   	ProfileArn: jsii.String("profileArn"),
//   	Title: jsii.String("title"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentcontext.html
//
type CfnAgentContextPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAgentContextMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentContextPropsMixin
type jsiiProxy_CfnAgentContextPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAgentContextPropsMixin) Props() *CfnAgentContextMixinProps {
	var returns *CfnAgentContextMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentContextPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WellArchitected::AgentContext`.
func NewCfnAgentContextPropsMixin(props *CfnAgentContextMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAgentContextPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentContextPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentContextPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentContextPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WellArchitected::AgentContext`.
func NewCfnAgentContextPropsMixin_Override(c CfnAgentContextPropsMixin, props *CfnAgentContextMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentContextPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAgentContextPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentContextPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentContextPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentContextPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentContextPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentContextPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAgentContextPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

