// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// The HTTP Proxy integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   // This function handles your auth logic
//   var authHandler function
//
//
//   authorizer := awscdkapigatewayv2authorizersalpha.NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &httpLambdaAuthorizerProps{
//   	responseTypes: []httpLambdaResponseType{
//   		awscdkapigatewayv2authorizersalpha.HttpLambdaResponseType_SIMPLE,
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
//
// Experimental.
type HttpUrlIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpUrlIntegration
type jsiiProxy_HttpUrlIntegration struct {
	internal.Type__awscdkapigatewayv2alphaHttpRouteIntegration
}

// Experimental.
func NewHttpUrlIntegration(id *string, url *string, props *HttpUrlIntegrationProps) HttpUrlIntegration {
	_init_.Initialize()

	if err := validateNewHttpUrlIntegrationParameters(id, url, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpUrlIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpUrlIntegration",
		[]interface{}{id, url, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpUrlIntegration_Override(h HttpUrlIntegration, id *string, url *string, props *HttpUrlIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpUrlIntegration",
		[]interface{}{id, url, props},
		h,
	)
}

func (h *jsiiProxy_HttpUrlIntegration) Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(_arg); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpUrlIntegration) CompleteBind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

