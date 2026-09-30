package awswellarchitected

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::WellArchitected::AgentGoal.
//
// An Agent Goal expresses a Well-Architected objective (a title, optional description, and the pillars it targets) within an Agent Profile, used to focus Well-Architected Agent recommendations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAgentGoalPropsMixin := awscdkcfnpropertymixins.Aws_wellarchitected.NewCfnAgentGoalPropsMixin(&CfnAgentGoalMixinProps{
//   	Description: jsii.String("description"),
//   	Pillars: []*string{
//   		jsii.String("pillars"),
//   	},
//   	ProfileArn: jsii.String("profileArn"),
//   	Title: jsii.String("title"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html
//
type CfnAgentGoalPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAgentGoalMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentGoalPropsMixin
type jsiiProxy_CfnAgentGoalPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAgentGoalPropsMixin) Props() *CfnAgentGoalMixinProps {
	var returns *CfnAgentGoalMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentGoalPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WellArchitected::AgentGoal`.
func NewCfnAgentGoalPropsMixin(props *CfnAgentGoalMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAgentGoalPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentGoalPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentGoalPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentGoalPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WellArchitected::AgentGoal`.
func NewCfnAgentGoalPropsMixin_Override(c CfnAgentGoalPropsMixin, props *CfnAgentGoalMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentGoalPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAgentGoalPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentGoalPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentGoalPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentGoalPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnAgentGoalPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentGoalPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAgentGoalPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

