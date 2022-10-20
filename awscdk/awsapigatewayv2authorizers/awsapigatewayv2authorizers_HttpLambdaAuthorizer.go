package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2authorizers/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Authorize Http Api routes via a lambda function.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // This function handles your auth logic
//   var authHandler function
//
//
//   authorizer := awscdk.NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &httpLambdaAuthorizerProps{
//   	responseTypes: []httpLambdaResponseType{
//   		awscdk.HttpLambdaResponseType_SIMPLE,
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
//
// Experimental.
type HttpLambdaAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpLambdaAuthorizer
type jsiiProxy_HttpLambdaAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Initialize a lambda authorizer to be bound with HTTP route.
// Experimental.
func NewHttpLambdaAuthorizer(id *string, handler awslambda.IFunction, props *HttpLambdaAuthorizerProps) HttpLambdaAuthorizer {
	_init_.Initialize()

	if err := validateNewHttpLambdaAuthorizerParameters(id, handler, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpLambdaAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpLambdaAuthorizer",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Initialize a lambda authorizer to be bound with HTTP route.
// Experimental.
func NewHttpLambdaAuthorizer_Override(h HttpLambdaAuthorizer, id *string, handler awslambda.IFunction, props *HttpLambdaAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpLambdaAuthorizer",
		[]interface{}{id, handler, props},
		h,
	)
}

func (h *jsiiProxy_HttpLambdaAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	if err := h.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

