package awsdevicefarm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a network profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNetworkProfilePropsMixin := awscdkcfnpropertymixins.Aws_devicefarm.NewCfnNetworkProfilePropsMixin(&CfnNetworkProfileMixinProps{
//   	Description: jsii.String("description"),
//   	DownlinkBandwidthBits: jsii.Number(123),
//   	DownlinkDelayMs: jsii.Number(123),
//   	DownlinkJitterMs: jsii.Number(123),
//   	DownlinkLossPercent: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	ProjectArn: jsii.String("projectArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UplinkBandwidthBits: jsii.Number(123),
//   	UplinkDelayMs: jsii.Number(123),
//   	UplinkJitterMs: jsii.Number(123),
//   	UplinkLossPercent: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-networkprofile.html
//
type CfnNetworkProfilePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNetworkProfileMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkProfilePropsMixin
type jsiiProxy_CfnNetworkProfilePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNetworkProfilePropsMixin) Props() *CfnNetworkProfileMixinProps {
	var returns *CfnNetworkProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkProfilePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DeviceFarm::NetworkProfile`.
func NewCfnNetworkProfilePropsMixin(props *CfnNetworkProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNetworkProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnNetworkProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DeviceFarm::NetworkProfile`.
func NewCfnNetworkProfilePropsMixin_Override(c CfnNetworkProfilePropsMixin, props *CfnNetworkProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnNetworkProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNetworkProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnNetworkProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_devicefarm.CfnNetworkProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNetworkProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

