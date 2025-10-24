package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations/internal"
)

// The HTTP Proxy integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // This function handles your auth logic
//   var authHandler Function
//
//
//   authorizer := awscdk.NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &HttpLambdaAuthorizerProps{
//   	ResponseTypes: []HttpLambdaResponseType{
//   		awscdk.HttpLambdaResponseType_SIMPLE,
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books"),
//   	Authorizer: Authorizer,
//   })
//
type HttpUrlIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Bind this integration to the route.
	Bind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpUrlIntegration
type jsiiProxy_HttpUrlIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

func NewHttpUrlIntegration(id *string, url *string, props *HttpUrlIntegrationProps) HttpUrlIntegration {
	_init_.Initialize()

	if err := validateNewHttpUrlIntegrationParameters(id, url, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpUrlIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpUrlIntegration",
		[]interface{}{id, url, props},
		&j,
	)

	return &j
}

func NewHttpUrlIntegration_Override(h HttpUrlIntegration, id *string, url *string, props *HttpUrlIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpUrlIntegration",
		[]interface{}{id, url, props},
		h,
	)
}

func (h *jsiiProxy_HttpUrlIntegration) Bind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpUrlIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

