package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// The Lambda Proxy integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var booksDefaultFn function
//
//   booksIntegration := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	Integration: booksIntegration,
//   })
//
// Deprecated.
type HttpLambdaIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	// Bind this integration to the route.
	// Deprecated.
	Bind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Deprecated.
	CompleteBind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpLambdaIntegration
type jsiiProxy_HttpLambdaIntegration struct {
	internal.Type__awscdkapigatewayv2alphaHttpRouteIntegration
}

// Deprecated.
func NewHttpLambdaIntegration(id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) HttpLambdaIntegration {
	_init_.Initialize()

	if err := validateNewHttpLambdaIntegrationParameters(id, handler, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpLambdaIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Deprecated.
func NewHttpLambdaIntegration_Override(h HttpLambdaIntegration, id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		h,
	)
}

func (h *jsiiProxy_HttpLambdaIntegration) Bind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpLambdaIntegration) CompleteBind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{options},
	)
}

