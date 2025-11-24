package mixinsawsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiotwireless/mixinsawsiotwireless/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// A FUOTA task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFuotaTaskPropsMixin := awscdkmixinspreview.Mixins.NewCfnFuotaTaskPropsMixin(&CfnFuotaTaskMixinProps{
//   	AssociateMulticastGroup: jsii.String("associateMulticastGroup"),
//   	AssociateWirelessDevice: jsii.String("associateWirelessDevice"),
//   	Description: jsii.String("description"),
//   	DisassociateMulticastGroup: jsii.String("disassociateMulticastGroup"),
//   	DisassociateWirelessDevice: jsii.String("disassociateWirelessDevice"),
//   	FirmwareUpdateImage: jsii.String("firmwareUpdateImage"),
//   	FirmwareUpdateRole: jsii.String("firmwareUpdateRole"),
//   	LoRaWan: &LoRaWANProperty{
//   		RfRegion: jsii.String("rfRegion"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html
//
type CfnFuotaTaskPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFuotaTaskMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFuotaTaskPropsMixin
type jsiiProxy_CfnFuotaTaskPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFuotaTaskPropsMixin) Props() *CfnFuotaTaskMixinProps {
	var returns *CfnFuotaTaskMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFuotaTaskPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::FuotaTask`.
func NewCfnFuotaTaskPropsMixin(props *CfnFuotaTaskMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFuotaTaskPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFuotaTaskPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFuotaTaskPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnFuotaTaskPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::FuotaTask`.
func NewCfnFuotaTaskPropsMixin_Override(c CfnFuotaTaskPropsMixin, props *CfnFuotaTaskMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnFuotaTaskPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFuotaTaskPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFuotaTaskPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnFuotaTaskPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFuotaTaskPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnFuotaTaskPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFuotaTaskPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFuotaTaskPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

