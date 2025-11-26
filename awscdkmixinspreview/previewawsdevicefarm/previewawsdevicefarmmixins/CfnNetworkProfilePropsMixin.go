package previewawsdevicefarmmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdevicefarm/previewawsdevicefarmmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a network profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnNetworkProfilePropsMixin(&CfnNetworkProfileMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-networkprofile.html
//
type CfnNetworkProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNetworkProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkProfilePropsMixin
type jsiiProxy_CfnNetworkProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnNetworkProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DeviceFarm::NetworkProfile`.
func NewCfnNetworkProfilePropsMixin(props *CfnNetworkProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNetworkProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devicefarm.mixins.CfnNetworkProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DeviceFarm::NetworkProfile`.
func NewCfnNetworkProfilePropsMixin_Override(c CfnNetworkProfilePropsMixin, props *CfnNetworkProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devicefarm.mixins.CfnNetworkProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNetworkProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devicefarm.mixins.CfnNetworkProfilePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_devicefarm.mixins.CfnNetworkProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

