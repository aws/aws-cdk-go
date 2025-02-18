package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The interface that various route integration classes will inherit.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bookStoreDefaultFn function
//
//
//   getBooksIntegration := awscdk.NewHttpUrlIntegration(jsii.String("GetBooksIntegration"), jsii.String("https://get-books-proxy.example.com"))
//   bookStoreDefaultIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), bookStoreDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	Integration: getBooksIntegration,
//   })
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []*httpMethod{
//   		apigwv2.*httpMethod_ANY,
//   	},
//   	Integration: bookStoreDefaultIntegration,
//   })
//
type HttpRouteIntegration interface {
	// Bind this integration to the route.
	Bind(options *HttpRouteIntegrationBindOptions) *HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	CompleteBind(_options *HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpRouteIntegration
type jsiiProxy_HttpRouteIntegration struct {
	_ byte // padding
}

// Initialize an integration for a route on http api.
func NewHttpRouteIntegration_Override(h HttpRouteIntegration, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.HttpRouteIntegration",
		[]interface{}{id},
		h,
	)
}

func (h *jsiiProxy_HttpRouteIntegration) Bind(options *HttpRouteIntegrationBindOptions) *HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpRouteIntegration) CompleteBind(_options *HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

