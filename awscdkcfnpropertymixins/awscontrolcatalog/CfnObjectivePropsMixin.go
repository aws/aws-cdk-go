package awscontrolcatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscontrolcatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Returns information about an objective in the AWS Control Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnObjectivePropsMixin := awscdkcfnpropertymixins.Aws_controlcatalog.NewCfnObjectivePropsMixin(&CfnObjectiveMixinProps{
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-controlcatalog-objective.html
//
type CfnObjectivePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnObjectiveMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnObjectivePropsMixin
type jsiiProxy_CfnObjectivePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnObjectivePropsMixin) Props() *CfnObjectiveMixinProps {
	var returns *CfnObjectiveMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObjectivePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ControlCatalog::Objective`.
func NewCfnObjectivePropsMixin(props *CfnObjectiveMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnObjectivePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnObjectivePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnObjectivePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_controlcatalog.CfnObjectivePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ControlCatalog::Objective`.
func NewCfnObjectivePropsMixin_Override(c CfnObjectivePropsMixin, props *CfnObjectiveMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_controlcatalog.CfnObjectivePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnObjectivePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnObjectivePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_controlcatalog.CfnObjectivePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnObjectivePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_controlcatalog.CfnObjectivePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnObjectivePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnObjectivePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

