package previewawsssomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssso/previewawsssomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an OAuth 2.0 customer managed application in IAM Identity Center for the given application provider.
//
// > This API does not support creating SAML 2.0 customer managed applications or AWS managed applications. To learn how to create an AWS managed application, see the application user guide. You can create a SAML 2.0 customer managed application in the the console only. See [Setting up customer managed SAML 2.0 applications](https://docs.aws.amazon.com/singlesignon/latest/userguide/customermanagedapps-saml2-setup.html) . For more information on these application types, see [AWS managed applications](https://docs.aws.amazon.com/singlesignon/latest/userguide/awsapps.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationPropsMixin(&CfnApplicationMixinProps{
//   	ApplicationProviderArn: jsii.String("applicationProviderArn"),
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	PortalOptions: &PortalOptionsConfigurationProperty{
//   		SignInOptions: &SignInOptionsProperty{
//   			ApplicationUrl: jsii.String("applicationUrl"),
//   			Origin: jsii.String("origin"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Status: jsii.String("status"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html
//
type CfnApplicationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationPropsMixin
type jsiiProxy_CfnApplicationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Props() *CfnApplicationMixinProps {
	var returns *CfnApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSO::Application`.
func NewCfnApplicationPropsMixin(props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSO::Application`.
func NewCfnApplicationPropsMixin_Override(c CfnApplicationPropsMixin, props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sso.mixins.CfnApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

