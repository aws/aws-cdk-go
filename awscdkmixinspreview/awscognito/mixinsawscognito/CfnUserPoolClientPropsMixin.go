package mixinsawscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscognito/mixinsawscognito/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::UserPoolClient` resource specifies an Amazon Cognito user pool client.
//
// > If you don't specify a value for a parameter, Amazon Cognito sets it to a default value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPoolClientPropsMixin := awscdkmixinspreview.Mixins.NewCfnUserPoolClientPropsMixin(&CfnUserPoolClientMixinProps{
//   	AccessTokenValidity: jsii.Number(123),
//   	AllowedOAuthFlows: []*string{
//   		jsii.String("allowedOAuthFlows"),
//   	},
//   	AllowedOAuthFlowsUserPoolClient: jsii.Boolean(false),
//   	AllowedOAuthScopes: []*string{
//   		jsii.String("allowedOAuthScopes"),
//   	},
//   	AnalyticsConfiguration: &AnalyticsConfigurationProperty{
//   		ApplicationArn: jsii.String("applicationArn"),
//   		ApplicationId: jsii.String("applicationId"),
//   		ExternalId: jsii.String("externalId"),
//   		RoleArn: jsii.String("roleArn"),
//   		UserDataShared: jsii.Boolean(false),
//   	},
//   	AuthSessionValidity: jsii.Number(123),
//   	CallbackUrLs: []*string{
//   		jsii.String("callbackUrLs"),
//   	},
//   	ClientName: jsii.String("clientName"),
//   	DefaultRedirectUri: jsii.String("defaultRedirectUri"),
//   	EnablePropagateAdditionalUserContextData: jsii.Boolean(false),
//   	EnableTokenRevocation: jsii.Boolean(false),
//   	ExplicitAuthFlows: []*string{
//   		jsii.String("explicitAuthFlows"),
//   	},
//   	GenerateSecret: jsii.Boolean(false),
//   	IdTokenValidity: jsii.Number(123),
//   	LogoutUrLs: []*string{
//   		jsii.String("logoutUrLs"),
//   	},
//   	PreventUserExistenceErrors: jsii.String("preventUserExistenceErrors"),
//   	ReadAttributes: []*string{
//   		jsii.String("readAttributes"),
//   	},
//   	RefreshTokenRotation: &RefreshTokenRotationProperty{
//   		Feature: jsii.String("feature"),
//   		RetryGracePeriodSeconds: jsii.Number(123),
//   	},
//   	RefreshTokenValidity: jsii.Number(123),
//   	SupportedIdentityProviders: []*string{
//   		jsii.String("supportedIdentityProviders"),
//   	},
//   	TokenValidityUnits: &TokenValidityUnitsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		IdToken: jsii.String("idToken"),
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   	WriteAttributes: []*string{
//   		jsii.String("writeAttributes"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html
//
type CfnUserPoolClientPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnUserPoolClientMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPoolClientPropsMixin
type jsiiProxy_CfnUserPoolClientPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnUserPoolClientPropsMixin) Props() *CfnUserPoolClientMixinProps {
	var returns *CfnUserPoolClientMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClientPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClientPropsMixin(props *CfnUserPoolClientMixinProps, options *mixins.CfnPropertyMixinOptions) CfnUserPoolClientPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPoolClientPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPoolClientPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClientPropsMixin_Override(c CfnUserPoolClientPropsMixin, props *CfnUserPoolClientMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnUserPoolClientPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPoolClientPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolClientPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolClientPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolClientPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnUserPoolClientPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

