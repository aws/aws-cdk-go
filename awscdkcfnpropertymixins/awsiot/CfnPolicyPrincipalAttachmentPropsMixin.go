package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::PolicyPrincipalAttachment` resource to attach an AWS IoT policy to a principal (an X.509 certificate or other credential).
//
// For information about working with AWS IoT policies and principals, see [Authorization](https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPolicyPrincipalAttachmentPropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnPolicyPrincipalAttachmentPropsMixin(&CfnPolicyPrincipalAttachmentMixinProps{
//   	PolicyName: jsii.String("policyName"),
//   	Principal: jsii.String("principal"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html
//
type CfnPolicyPrincipalAttachmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPolicyPrincipalAttachmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyPrincipalAttachmentPropsMixin
type jsiiProxy_CfnPolicyPrincipalAttachmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachmentPropsMixin) Props() *CfnPolicyPrincipalAttachmentMixinProps {
	var returns *CfnPolicyPrincipalAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPrincipalAttachmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::PolicyPrincipalAttachment`.
func NewCfnPolicyPrincipalAttachmentPropsMixin(props *CfnPolicyPrincipalAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPolicyPrincipalAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyPrincipalAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyPrincipalAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnPolicyPrincipalAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::PolicyPrincipalAttachment`.
func NewCfnPolicyPrincipalAttachmentPropsMixin_Override(c CfnPolicyPrincipalAttachmentPropsMixin, props *CfnPolicyPrincipalAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnPolicyPrincipalAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPolicyPrincipalAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyPrincipalAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnPolicyPrincipalAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyPrincipalAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnPolicyPrincipalAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPolicyPrincipalAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

