package awsservicediscovery

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A complex type that contains information about an instance that AWS Cloud Map creates when you submit a `RegisterInstance` request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceAttributes interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnInstancePropsMixin := awscdkcfnpropertymixins.Aws_servicediscovery.NewCfnInstancePropsMixin(&CfnInstanceMixinProps{
//   	InstanceAttributes: instanceAttributes,
//   	InstanceId: jsii.String("instanceId"),
//   	ServiceId: jsii.String("serviceId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html
//
type CfnInstancePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnInstanceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstancePropsMixin
type jsiiProxy_CfnInstancePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnInstancePropsMixin) Props() *CfnInstanceMixinProps {
	var returns *CfnInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstancePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceDiscovery::Instance`.
func NewCfnInstancePropsMixin(props *CfnInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicediscovery.CfnInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceDiscovery::Instance`.
func NewCfnInstancePropsMixin_Override(c CfnInstancePropsMixin, props *CfnInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_servicediscovery.CfnInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_servicediscovery.CfnInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_servicediscovery.CfnInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstancePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

