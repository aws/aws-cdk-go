package mixinsawsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiam/mixinsawsiam/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates an IAM entity to describe an identity provider (IdP) that supports [OpenID Connect (OIDC)](https://docs.aws.amazon.com/http://openid.net/connect/) .
//
// The OIDC provider that you create with this operation can be used as a principal in a role's trust policy. Such a policy establishes a trust relationship between AWS and the OIDC provider.
//
// When you create the IAM OIDC provider, you specify the following:
//
// - The URL of the OIDC identity provider (IdP) to trust
// - A list of client IDs (also known as audiences) that identify the application or applications that are allowed to authenticate using the OIDC provider
// - A list of tags that are attached to the specified IAM OIDC provider
// - A list of thumbprints of one or more server certificates that the IdP uses
//
// You get all of this information from the OIDC IdP that you want to use to access AWS .
//
// When you update the IAM OIDC provider, you specify the following:
//
// - The URL of the OIDC identity provider (IdP) to trust
// - A list of client IDs (also known as audiences) that replaces the existing list of client IDs associated with the OIDC IdP
// - A list of tags that replaces the existing list of tags attached to the specified IAM OIDC provider
// - A list of thumbprints that replaces the existing list of server certificates thumbprints that the IdP uses
//
// > The trust for the OIDC provider is derived from the IAM provider that this operation creates. Therefore, it is best to limit access to the [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) operation to highly privileged users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOIDCProviderPropsMixin := awscdkmixinspreview.Mixins.NewCfnOIDCProviderPropsMixin(&CfnOIDCProviderMixinProps{
//   	ClientIdList: []*string{
//   		jsii.String("clientIdList"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThumbprintList: []*string{
//   		jsii.String("thumbprintList"),
//   	},
//   	Url: jsii.String("url"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-oidcprovider.html
//
type CfnOIDCProviderPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOIDCProviderMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOIDCProviderPropsMixin
type jsiiProxy_CfnOIDCProviderPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOIDCProviderPropsMixin) Props() *CfnOIDCProviderMixinProps {
	var returns *CfnOIDCProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOIDCProviderPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IAM::OIDCProvider`.
func NewCfnOIDCProviderPropsMixin(props *CfnOIDCProviderMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOIDCProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOIDCProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOIDCProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnOIDCProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IAM::OIDCProvider`.
func NewCfnOIDCProviderPropsMixin_Override(c CfnOIDCProviderPropsMixin, props *CfnOIDCProviderMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnOIDCProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOIDCProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOIDCProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnOIDCProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOIDCProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iam.mixins.CfnOIDCProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOIDCProviderPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnOIDCProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

