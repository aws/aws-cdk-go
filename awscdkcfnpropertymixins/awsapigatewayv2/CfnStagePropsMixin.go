package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGatewayV2::Stage` resource specifies a stage for an API.
//
// Each stage is a named reference to a deployment of the API and is made available for client applications to call. To learn more, see [Working with stages for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-stages.html) and [Deploy a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-set-up-websocket-deployment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var routeSettings interface{}
//   var stageVariables interface{}
//   var tags interface{}
//
//   cfnStagePropsMixin := awscdkcfnpropertymixins.Aws_apigatewayv2.NewCfnStagePropsMixin(&CfnStageMixinProps{
//   	AccessLogSettings: &AccessLogSettingsProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		Format: jsii.String("format"),
//   	},
//   	AccessPolicyId: jsii.String("accessPolicyId"),
//   	ApiId: jsii.String("apiId"),
//   	AutoDeploy: jsii.Boolean(false),
//   	ClientCertificateId: jsii.String("clientCertificateId"),
//   	DefaultRouteSettings: &RouteSettingsProperty{
//   		DataTraceEnabled: jsii.Boolean(false),
//   		DetailedMetricsEnabled: jsii.Boolean(false),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   	},
//   	DeploymentId: jsii.String("deploymentId"),
//   	Description: jsii.String("description"),
//   	RouteSettings: routeSettings,
//   	StageName: jsii.String("stageName"),
//   	StageVariables: stageVariables,
//   	Tags: tags,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html
//
type CfnStagePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnStageMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStagePropsMixin
type jsiiProxy_CfnStagePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnStagePropsMixin) Props() *CfnStageMixinProps {
	var returns *CfnStageMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStagePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGatewayV2::Stage`.
func NewCfnStagePropsMixin(props *CfnStageMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnStagePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStagePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStagePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnStagePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGatewayV2::Stage`.
func NewCfnStagePropsMixin_Override(c CfnStagePropsMixin, props *CfnStageMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnStagePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnStagePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStagePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnStagePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStagePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_apigatewayv2.CfnStagePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStagePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnStagePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

