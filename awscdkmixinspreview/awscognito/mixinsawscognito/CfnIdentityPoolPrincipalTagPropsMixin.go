package mixinsawscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscognito/mixinsawscognito/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// A list of the identity pool principal tag assignments for attributes for access control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var principalTags interface{}
//
//   cfnIdentityPoolPrincipalTagPropsMixin := awscdkmixinspreview.Mixins.NewCfnIdentityPoolPrincipalTagPropsMixin(&CfnIdentityPoolPrincipalTagMixinProps{
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   	IdentityProviderName: jsii.String("identityProviderName"),
//   	PrincipalTags: principalTags,
//   	UseDefaults: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolprincipaltag.html
//
type CfnIdentityPoolPrincipalTagPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIdentityPoolPrincipalTagMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdentityPoolPrincipalTagPropsMixin
type jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin) Props() *CfnIdentityPoolPrincipalTagMixinProps {
	var returns *CfnIdentityPoolPrincipalTagMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::IdentityPoolPrincipalTag`.
func NewCfnIdentityPoolPrincipalTagPropsMixin(props *CfnIdentityPoolPrincipalTagMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIdentityPoolPrincipalTagPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdentityPoolPrincipalTagPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPrincipalTagPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::IdentityPoolPrincipalTag`.
func NewCfnIdentityPoolPrincipalTagPropsMixin_Override(c CfnIdentityPoolPrincipalTagPropsMixin, props *CfnIdentityPoolPrincipalTagMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPrincipalTagPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdentityPoolPrincipalTagPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdentityPoolPrincipalTagPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPrincipalTagPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentityPoolPrincipalTagPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPrincipalTagPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIdentityPoolPrincipalTagPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

