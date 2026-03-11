package awsglobalaccelerator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::GlobalAccelerator::Listener` resource is a Global Accelerator resource type that contains information about how you create a listener to process inbound connections from clients to an accelerator.
//
// Connections arrive to assigned static IP addresses on a port, port range, or list of port ranges that you specify.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnListenerPropsMixin := awscdkcfnpropertymixins.Aws_globalaccelerator.NewCfnListenerPropsMixin(&CfnListenerMixinProps{
//   	AcceleratorArn: jsii.String("acceleratorArn"),
//   	ClientAffinity: jsii.String("clientAffinity"),
//   	PortRanges: []interface{}{
//   		&PortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Protocol: jsii.String("protocol"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html
//
type CfnListenerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnListenerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnListenerPropsMixin
type jsiiProxy_CfnListenerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnListenerPropsMixin) Props() *CfnListenerMixinProps {
	var returns *CfnListenerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GlobalAccelerator::Listener`.
func NewCfnListenerPropsMixin(props *CfnListenerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnListenerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnListenerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnListenerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_globalaccelerator.CfnListenerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GlobalAccelerator::Listener`.
func NewCfnListenerPropsMixin_Override(c CfnListenerPropsMixin, props *CfnListenerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_globalaccelerator.CfnListenerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnListenerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnListenerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_globalaccelerator.CfnListenerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListenerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_globalaccelerator.CfnListenerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnListenerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnListenerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

