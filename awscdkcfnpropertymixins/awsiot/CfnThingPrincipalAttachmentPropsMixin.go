package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::ThingPrincipalAttachment` resource to attach a principal (an X.509 certificate or another credential) to a thing.
//
// For more information about working with AWS IoT things and principals, see [Authorization](https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnThingPrincipalAttachmentPropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnThingPrincipalAttachmentPropsMixin(&CfnThingPrincipalAttachmentMixinProps{
//   	Principal: jsii.String("principal"),
//   	ThingName: jsii.String("thingName"),
//   	ThingPrincipalType: jsii.String("thingPrincipalType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html
//
type CfnThingPrincipalAttachmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnThingPrincipalAttachmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnThingPrincipalAttachmentPropsMixin
type jsiiProxy_CfnThingPrincipalAttachmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnThingPrincipalAttachmentPropsMixin) Props() *CfnThingPrincipalAttachmentMixinProps {
	var returns *CfnThingPrincipalAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThingPrincipalAttachmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::ThingPrincipalAttachment`.
func NewCfnThingPrincipalAttachmentPropsMixin(props *CfnThingPrincipalAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnThingPrincipalAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnThingPrincipalAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnThingPrincipalAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnThingPrincipalAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::ThingPrincipalAttachment`.
func NewCfnThingPrincipalAttachmentPropsMixin_Override(c CfnThingPrincipalAttachmentPropsMixin, props *CfnThingPrincipalAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnThingPrincipalAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnThingPrincipalAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnThingPrincipalAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnThingPrincipalAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnThingPrincipalAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnThingPrincipalAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnThingPrincipalAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnThingPrincipalAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

