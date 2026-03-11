package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::UserPoolClient` resource specifies an Amazon Cognito user pool client.
//
// > If you don't specify a value for a parameter, Amazon Cognito sets it to a default value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUserPoolClientPropsMixin := awscdkcfnpropertymixins.Aws_cognito.NewCfnUserPoolClientPropsMixin(&CfnUserPoolClientMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html
//
type CfnUserPoolClientPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserPoolClientMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPoolClientPropsMixin
type jsiiProxy_CfnUserPoolClientPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnUserPoolClientPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClientPropsMixin(props *CfnUserPoolClientMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserPoolClientPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPoolClientPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPoolClientPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClientPropsMixin_Override(c CfnUserPoolClientPropsMixin, props *CfnUserPoolClientMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserPoolClientPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPoolClientPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_cognito.CfnUserPoolClientPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolClientPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

