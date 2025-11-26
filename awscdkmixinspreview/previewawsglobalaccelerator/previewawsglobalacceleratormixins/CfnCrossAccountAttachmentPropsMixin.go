package previewawsglobalacceleratormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglobalaccelerator/previewawsglobalacceleratormixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a cross-account attachment in AWS Global Accelerator .
//
// You create a cross-account attachment to specify the *principals* who have permission to work with *resources* in accelerators in their own account. You specify, in the same attachment, the resources that are shared.
//
// A principal can be an AWS account number or the Amazon Resource Name (ARN) for an accelerator. For account numbers that are listed as principals, to work with a resource listed in the attachment, you must sign in to an account specified as a principal. Then, you can work with resources that are listed, with any of your accelerators. If an accelerator ARN is listed in the cross-account attachment as a principal, anyone with permission to make updates to the accelerator can work with resources that are listed in the attachment.
//
// Specify each principal and resource separately. To specify two CIDR address pools, list them individually under `Resources` , and so on. For a command line operation, for example, you might use a statement like the following:
//
// `"Resources": [{"Cidr": "169.254.60.0/24"},{"Cidr": "169.254.59.0/24"}]`
//
// For more information, see [Working with cross-account attachments and resources in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html) in the *AWS Global Accelerator Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCrossAccountAttachmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnCrossAccountAttachmentPropsMixin(&CfnCrossAccountAttachmentMixinProps{
//   	Name: jsii.String("name"),
//   	Principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	Resources: []interface{}{
//   		&ResourceProperty{
//   			Cidr: jsii.String("cidr"),
//   			EndpointId: jsii.String("endpointId"),
//   			Region: jsii.String("region"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-crossaccountattachment.html
//
type CfnCrossAccountAttachmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCrossAccountAttachmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCrossAccountAttachmentPropsMixin
type jsiiProxy_CfnCrossAccountAttachmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCrossAccountAttachmentPropsMixin) Props() *CfnCrossAccountAttachmentMixinProps {
	var returns *CfnCrossAccountAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrossAccountAttachmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GlobalAccelerator::CrossAccountAttachment`.
func NewCfnCrossAccountAttachmentPropsMixin(props *CfnCrossAccountAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCrossAccountAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCrossAccountAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCrossAccountAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnCrossAccountAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GlobalAccelerator::CrossAccountAttachment`.
func NewCfnCrossAccountAttachmentPropsMixin_Override(c CfnCrossAccountAttachmentPropsMixin, props *CfnCrossAccountAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnCrossAccountAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCrossAccountAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCrossAccountAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnCrossAccountAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCrossAccountAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_globalaccelerator.mixins.CfnCrossAccountAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCrossAccountAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCrossAccountAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

