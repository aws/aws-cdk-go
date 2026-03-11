package awsdevicefarm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a request to the create device pool operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDevicePoolPropsMixin := awscdkcfnpropertymixins.Aws_devicefarm.NewCfnDevicePoolPropsMixin(&CfnDevicePoolMixinProps{
//   	Description: jsii.String("description"),
//   	MaxDevices: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	ProjectArn: jsii.String("projectArn"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Attribute: jsii.String("attribute"),
//   			Operator: jsii.String("operator"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-devicepool.html
//
type CfnDevicePoolPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDevicePoolMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDevicePoolPropsMixin
type jsiiProxy_CfnDevicePoolPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDevicePoolPropsMixin) Props() *CfnDevicePoolMixinProps {
	var returns *CfnDevicePoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevicePoolPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DeviceFarm::DevicePool`.
func NewCfnDevicePoolPropsMixin(props *CfnDevicePoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDevicePoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDevicePoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDevicePoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnDevicePoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DeviceFarm::DevicePool`.
func NewCfnDevicePoolPropsMixin_Override(c CfnDevicePoolPropsMixin, props *CfnDevicePoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnDevicePoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDevicePoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDevicePoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnDevicePoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDevicePoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnDevicePoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDevicePoolPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDevicePoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

