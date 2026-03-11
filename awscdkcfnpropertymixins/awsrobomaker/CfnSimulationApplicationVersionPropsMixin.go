package awsrobomaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RoboMaker::SimulationApplicationVersion` resource creates a version of an AWS RoboMaker simulation application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSimulationApplicationVersionPropsMixin := awscdkcfnpropertymixins.Aws_robomaker.NewCfnSimulationApplicationVersionPropsMixin(&CfnSimulationApplicationVersionMixinProps{
//   	Application: jsii.String("application"),
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-simulationapplicationversion.html
//
type CfnSimulationApplicationVersionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSimulationApplicationVersionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSimulationApplicationVersionPropsMixin
type jsiiProxy_CfnSimulationApplicationVersionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSimulationApplicationVersionPropsMixin) Props() *CfnSimulationApplicationVersionMixinProps {
	var returns *CfnSimulationApplicationVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimulationApplicationVersionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RoboMaker::SimulationApplicationVersion`.
func NewCfnSimulationApplicationVersionPropsMixin(props *CfnSimulationApplicationVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSimulationApplicationVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSimulationApplicationVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSimulationApplicationVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_robomaker.CfnSimulationApplicationVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RoboMaker::SimulationApplicationVersion`.
func NewCfnSimulationApplicationVersionPropsMixin_Override(c CfnSimulationApplicationVersionPropsMixin, props *CfnSimulationApplicationVersionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_robomaker.CfnSimulationApplicationVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSimulationApplicationVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSimulationApplicationVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_robomaker.CfnSimulationApplicationVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSimulationApplicationVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_robomaker.CfnSimulationApplicationVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSimulationApplicationVersionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSimulationApplicationVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

