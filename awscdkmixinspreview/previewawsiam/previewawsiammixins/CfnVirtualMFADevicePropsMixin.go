package previewawsiammixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiam/previewawsiammixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new virtual MFA device for the AWS account .
//
// After creating the virtual MFA, use [EnableMFADevice](https://docs.aws.amazon.com/IAM/latest/APIReference/API_EnableMFADevice.html) to attach the MFA device to an IAM user. For more information about creating and working with virtual MFA devices, see [Using a virtual MFA device](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html) in the *IAM User Guide* .
//
// For information about the maximum number of MFA devices you can create, see [IAM and AWS STS quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide* .
//
// > The seed information contained in the QR code and the Base32 string should be treated like any other secret access information. In other words, protect the seed information as you would your AWS access keys or your passwords. After you provision your virtual device, you should ensure that the information is destroyed following secure procedures.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVirtualMFADevicePropsMixin := awscdkmixinspreview.Mixins.NewCfnVirtualMFADevicePropsMixin(&CfnVirtualMFADeviceMixinProps{
//   	Path: jsii.String("path"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Users: []*string{
//   		jsii.String("users"),
//   	},
//   	VirtualMfaDeviceName: jsii.String("virtualMfaDeviceName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-virtualmfadevice.html
//
type CfnVirtualMFADevicePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVirtualMFADeviceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVirtualMFADevicePropsMixin
type jsiiProxy_CfnVirtualMFADevicePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVirtualMFADevicePropsMixin) Props() *CfnVirtualMFADeviceMixinProps {
	var returns *CfnVirtualMFADeviceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVirtualMFADevicePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IAM::VirtualMFADevice`.
func NewCfnVirtualMFADevicePropsMixin(props *CfnVirtualMFADeviceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVirtualMFADevicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVirtualMFADevicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVirtualMFADevicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnVirtualMFADevicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IAM::VirtualMFADevice`.
func NewCfnVirtualMFADevicePropsMixin_Override(c CfnVirtualMFADevicePropsMixin, props *CfnVirtualMFADeviceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnVirtualMFADevicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVirtualMFADevicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVirtualMFADevicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnVirtualMFADevicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVirtualMFADevicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnVirtualMFADevicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVirtualMFADevicePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVirtualMFADevicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

