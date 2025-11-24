package mixinsawsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsstepfunctions/mixinsawsstepfunctions/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a state machine [alias](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html) . An alias routes traffic to one or two versions of the same state machine.
//
// You can create up to 100 aliases for each state machine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineAliasPropsMixin := awscdkmixinspreview.Mixins.NewCfnStateMachineAliasPropsMixin(&CfnStateMachineAliasMixinProps{
//   	DeploymentPreference: &DeploymentPreferenceProperty{
//   		Alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		Interval: jsii.Number(123),
//   		Percentage: jsii.Number(123),
//   		StateMachineVersionArn: jsii.String("stateMachineVersionArn"),
//   		Type: jsii.String("type"),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RoutingConfiguration: []interface{}{
//   		&RoutingConfigurationVersionProperty{
//   			StateMachineVersionArn: jsii.String("stateMachineVersionArn"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html
//
type CfnStateMachineAliasPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStateMachineAliasMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStateMachineAliasPropsMixin
type jsiiProxy_CfnStateMachineAliasPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStateMachineAliasPropsMixin) Props() *CfnStateMachineAliasMixinProps {
	var returns *CfnStateMachineAliasMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachineAliasPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::StepFunctions::StateMachineAlias`.
func NewCfnStateMachineAliasPropsMixin(props *CfnStateMachineAliasMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStateMachineAliasPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStateMachineAliasPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStateMachineAliasPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::StepFunctions::StateMachineAlias`.
func NewCfnStateMachineAliasPropsMixin_Override(c CfnStateMachineAliasPropsMixin, props *CfnStateMachineAliasMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStateMachineAliasPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStateMachineAliasPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStateMachineAliasPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineAliasPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStateMachineAliasPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStateMachineAliasPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

