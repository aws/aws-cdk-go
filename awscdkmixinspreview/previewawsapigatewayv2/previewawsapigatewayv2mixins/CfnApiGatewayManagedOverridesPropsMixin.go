package previewawsapigatewayv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapigatewayv2/previewawsapigatewayv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGatewayV2::ApiGatewayManagedOverrides` resource overrides the default properties of API Gateway-managed resources that are implicitly configured for you when you use quick create.
//
// When you create an API by using quick create, an `AWS::ApiGatewayV2::Route` , `AWS::ApiGatewayV2::Integration` , and `AWS::ApiGatewayV2::Stage` are created for you and associated with your `AWS::ApiGatewayV2::Api` . The `AWS::ApiGatewayV2::ApiGatewayManagedOverrides` resource enables you to set, or override the properties of these implicit resources. Supported only for HTTP APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var routeSettings interface{}
//   var stageVariables interface{}
//
//   cfnApiGatewayManagedOverridesPropsMixin := awscdkmixinspreview.Mixins.NewCfnApiGatewayManagedOverridesPropsMixin(&CfnApiGatewayManagedOverridesMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	Integration: &IntegrationOverridesProperty{
//   		Description: jsii.String("description"),
//   		IntegrationMethod: jsii.String("integrationMethod"),
//   		PayloadFormatVersion: jsii.String("payloadFormatVersion"),
//   		TimeoutInMillis: jsii.Number(123),
//   	},
//   	Route: &RouteOverridesProperty{
//   		AuthorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		AuthorizationType: jsii.String("authorizationType"),
//   		AuthorizerId: jsii.String("authorizerId"),
//   		OperationName: jsii.String("operationName"),
//   		Target: jsii.String("target"),
//   	},
//   	Stage: &StageOverridesProperty{
//   		AccessLogSettings: &AccessLogSettingsProperty{
//   			DestinationArn: jsii.String("destinationArn"),
//   			Format: jsii.String("format"),
//   		},
//   		AutoDeploy: jsii.Boolean(false),
//   		DefaultRouteSettings: &RouteSettingsProperty{
//   			DataTraceEnabled: jsii.Boolean(false),
//   			DetailedMetricsEnabled: jsii.Boolean(false),
//   			LoggingLevel: jsii.String("loggingLevel"),
//   			ThrottlingBurstLimit: jsii.Number(123),
//   			ThrottlingRateLimit: jsii.Number(123),
//   		},
//   		Description: jsii.String("description"),
//   		RouteSettings: routeSettings,
//   		StageVariables: stageVariables,
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apigatewaymanagedoverrides.html
//
type CfnApiGatewayManagedOverridesPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApiGatewayManagedOverridesMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApiGatewayManagedOverridesPropsMixin
type jsiiProxy_CfnApiGatewayManagedOverridesPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApiGatewayManagedOverridesPropsMixin) Props() *CfnApiGatewayManagedOverridesMixinProps {
	var returns *CfnApiGatewayManagedOverridesMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverridesPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGatewayV2::ApiGatewayManagedOverrides`.
func NewCfnApiGatewayManagedOverridesPropsMixin(props *CfnApiGatewayManagedOverridesMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApiGatewayManagedOverridesPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApiGatewayManagedOverridesPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApiGatewayManagedOverridesPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnApiGatewayManagedOverridesPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGatewayV2::ApiGatewayManagedOverrides`.
func NewCfnApiGatewayManagedOverridesPropsMixin_Override(c CfnApiGatewayManagedOverridesPropsMixin, props *CfnApiGatewayManagedOverridesMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnApiGatewayManagedOverridesPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApiGatewayManagedOverridesPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApiGatewayManagedOverridesPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnApiGatewayManagedOverridesPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiGatewayManagedOverridesPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_apigatewayv2.mixins.CfnApiGatewayManagedOverridesPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApiGatewayManagedOverridesPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnApiGatewayManagedOverridesPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

