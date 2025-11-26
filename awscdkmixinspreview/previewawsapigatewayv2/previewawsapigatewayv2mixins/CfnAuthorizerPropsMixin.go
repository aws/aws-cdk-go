package previewawsapigatewayv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigatewayv2/previewawsapigatewayv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGatewayV2::Authorizer` resource creates an authorizer for a WebSocket API or an HTTP API.
//
// To learn more, see [Controlling and managing access to a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-control-access.html) and [Controlling and managing access to an HTTP API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html) in the *API Gateway Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAuthorizerPropsMixin := awscdkmixinspreview.Mixins.NewCfnAuthorizerPropsMixin(&CfnAuthorizerMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	AuthorizerCredentialsArn: jsii.String("authorizerCredentialsArn"),
//   	AuthorizerPayloadFormatVersion: jsii.String("authorizerPayloadFormatVersion"),
//   	AuthorizerResultTtlInSeconds: jsii.Number(123),
//   	AuthorizerType: jsii.String("authorizerType"),
//   	AuthorizerUri: jsii.String("authorizerUri"),
//   	EnableSimpleResponses: jsii.Boolean(false),
//   	IdentitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	JwtConfiguration: &JWTConfigurationProperty{
//   		Audience: []*string{
//   			jsii.String("audience"),
//   		},
//   		Issuer: jsii.String("issuer"),
//   	},
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html
//
type CfnAuthorizerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAuthorizerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAuthorizerPropsMixin
type jsiiProxy_CfnAuthorizerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAuthorizerPropsMixin) Props() *CfnAuthorizerMixinProps {
	var returns *CfnAuthorizerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGatewayV2::Authorizer`.
func NewCfnAuthorizerPropsMixin(props *CfnAuthorizerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAuthorizerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAuthorizerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAuthorizerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnAuthorizerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGatewayV2::Authorizer`.
func NewCfnAuthorizerPropsMixin_Override(c CfnAuthorizerPropsMixin, props *CfnAuthorizerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnAuthorizerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAuthorizerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAuthorizerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnAuthorizerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAuthorizerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnAuthorizerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAuthorizerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAuthorizerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

