package previewawsapigatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigateway/previewawsapigatewaymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGateway::Authorizer` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled.
//
// API Gateway activates the authorizer when a client calls those methods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAuthorizerPropsMixin := awscdkmixinspreview.Mixins.NewCfnAuthorizerPropsMixin(&CfnAuthorizerMixinProps{
//   	AuthorizerCredentials: jsii.String("authorizerCredentials"),
//   	AuthorizerResultTtlInSeconds: jsii.Number(123),
//   	AuthorizerUri: jsii.String("authorizerUri"),
//   	AuthType: jsii.String("authType"),
//   	IdentitySource: jsii.String("identitySource"),
//   	IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	Name: jsii.String("name"),
//   	ProviderArns: []*string{
//   		jsii.String("providerArns"),
//   	},
//   	RestApiId: jsii.String("restApiId"),
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html
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


// Create a mixin to apply properties to `AWS::ApiGateway::Authorizer`.
func NewCfnAuthorizerPropsMixin(props *CfnAuthorizerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAuthorizerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAuthorizerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAuthorizerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnAuthorizerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGateway::Authorizer`.
func NewCfnAuthorizerPropsMixin_Override(c CfnAuthorizerPropsMixin, props *CfnAuthorizerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnAuthorizerPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnAuthorizerPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnAuthorizerPropsMixin",
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

