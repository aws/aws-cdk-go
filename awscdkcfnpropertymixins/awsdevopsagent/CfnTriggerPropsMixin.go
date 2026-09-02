package awsdevopsagent

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdevopsagent/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::DevOpsAgent::Trigger.
//
// A trigger defines an automated action that fires on a schedule within an Agent Space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var action interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnTriggerPropsMixin := awscdkcfnpropertymixins.Aws_devopsagent.NewCfnTriggerPropsMixin(&CfnTriggerMixinProps{
//   	Action: action,
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	Condition: &ConditionProperty{
//   		Schedule: &ScheduleProperty{
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-trigger.html
//
type CfnTriggerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTriggerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTriggerPropsMixin
type jsiiProxy_CfnTriggerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTriggerPropsMixin) Props() *CfnTriggerMixinProps {
	var returns *CfnTriggerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTriggerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DevOpsAgent::Trigger`.
func NewCfnTriggerPropsMixin(props *CfnTriggerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTriggerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTriggerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTriggerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnTriggerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DevOpsAgent::Trigger`.
func NewCfnTriggerPropsMixin_Override(c CfnTriggerPropsMixin, props *CfnTriggerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnTriggerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTriggerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTriggerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnTriggerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTriggerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnTriggerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTriggerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTriggerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

