package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Attaches an elastic network interface (ENI) to an Amazon EC2 instance.
//
// You can use this resource type to attach additional network interfaces to an instance without interruption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkInterfaceAttachmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnNetworkInterfaceAttachmentPropsMixin(&CfnNetworkInterfaceAttachmentMixinProps{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	DeviceIndex: jsii.String("deviceIndex"),
//   	EnaQueueCount: jsii.Number(123),
//   	EnaSrdSpecification: &EnaSrdSpecificationProperty{
//   		EnaSrdEnabled: jsii.Boolean(false),
//   		EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   			EnaSrdUdpEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	InstanceId: jsii.String("instanceId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html
//
type CfnNetworkInterfaceAttachmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNetworkInterfaceAttachmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkInterfaceAttachmentPropsMixin
type jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin) Props() *CfnNetworkInterfaceAttachmentMixinProps {
	var returns *CfnNetworkInterfaceAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NetworkInterfaceAttachment`.
func NewCfnNetworkInterfaceAttachmentPropsMixin(props *CfnNetworkInterfaceAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNetworkInterfaceAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkInterfaceAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfaceAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NetworkInterfaceAttachment`.
func NewCfnNetworkInterfaceAttachmentPropsMixin_Override(c CfnNetworkInterfaceAttachmentPropsMixin, props *CfnNetworkInterfaceAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfaceAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNetworkInterfaceAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkInterfaceAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfaceAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkInterfaceAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkInterfaceAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNetworkInterfaceAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

