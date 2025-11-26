package previewawsapigatewayv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigatewayv2/previewawsapigatewayv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGatewayV2::Integration` resource creates an integration for an API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var requestParameters interface{}
//   var requestTemplates interface{}
//   var responseParameters interface{}
//
//   cfnIntegrationPropsMixin := awscdkmixinspreview.Mixins.NewCfnIntegrationPropsMixin(&CfnIntegrationMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	ConnectionId: jsii.String("connectionId"),
//   	ConnectionType: jsii.String("connectionType"),
//   	ContentHandlingStrategy: jsii.String("contentHandlingStrategy"),
//   	CredentialsArn: jsii.String("credentialsArn"),
//   	Description: jsii.String("description"),
//   	IntegrationMethod: jsii.String("integrationMethod"),
//   	IntegrationSubtype: jsii.String("integrationSubtype"),
//   	IntegrationType: jsii.String("integrationType"),
//   	IntegrationUri: jsii.String("integrationUri"),
//   	PassthroughBehavior: jsii.String("passthroughBehavior"),
//   	PayloadFormatVersion: jsii.String("payloadFormatVersion"),
//   	RequestParameters: requestParameters,
//   	RequestTemplates: requestTemplates,
//   	ResponseParameters: responseParameters,
//   	TemplateSelectionExpression: jsii.String("templateSelectionExpression"),
//   	TimeoutInMillis: jsii.Number(123),
//   	TlsConfig: &TlsConfigProperty{
//   		ServerNameToVerify: jsii.String("serverNameToVerify"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html
//
type CfnIntegrationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIntegrationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIntegrationPropsMixin
type jsiiProxy_CfnIntegrationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIntegrationPropsMixin) Props() *CfnIntegrationMixinProps {
	var returns *CfnIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGatewayV2::Integration`.
func NewCfnIntegrationPropsMixin(props *CfnIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGatewayV2::Integration`.
func NewCfnIntegrationPropsMixin_Override(c CfnIntegrationPropsMixin, props *CfnIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

