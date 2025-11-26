package previewawssmsvoicemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssmsvoice/previewawssmsvoicemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Attaches a resource-based policy to a End User Messaging  resource(phone number, sender Id, phone poll, or opt-out list) that is used for sharing the resource.
//
// A shared resource can be a Pool, Opt-out list, Sender Id, or Phone number. For more information about resource-based policies, see [Working with shared resources](https://docs.aws.amazon.com/sms-voice/latest/userguide/shared-resources.html) in the *End User Messaging  User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnResourcePolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnResourcePolicyPropsMixin(&CfnResourcePolicyMixinProps{
//   	PolicyDocument: policyDocument,
//   	ResourceArn: jsii.String("resourceArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-resourcepolicy.html
//
type CfnResourcePolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResourcePolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourcePolicyPropsMixin
type jsiiProxy_CfnResourcePolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResourcePolicyPropsMixin) Props() *CfnResourcePolicyMixinProps {
	var returns *CfnResourcePolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SMSVOICE::ResourcePolicy`.
func NewCfnResourcePolicyPropsMixin(props *CfnResourcePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResourcePolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourcePolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourcePolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnResourcePolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SMSVOICE::ResourcePolicy`.
func NewCfnResourcePolicyPropsMixin_Override(c CfnResourcePolicyPropsMixin, props *CfnResourcePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnResourcePolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResourcePolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourcePolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnResourcePolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourcePolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnResourcePolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

