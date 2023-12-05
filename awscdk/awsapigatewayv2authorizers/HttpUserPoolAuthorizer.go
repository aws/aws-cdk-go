package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   authorizer := awscdk.NewHttpUserPoolAuthorizer(jsii.String("BooksAuthorizer"), userPool)
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books"),
//   	Authorizer: Authorizer,
//   })
//
type HttpUserPoolAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpUserPoolAuthorizer
type jsiiProxy_HttpUserPoolAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Initialize a Cognito user pool authorizer to be bound with HTTP route.
func NewHttpUserPoolAuthorizer(id *string, pool awscognito.IUserPool, props *HttpUserPoolAuthorizerProps) HttpUserPoolAuthorizer {
	_init_.Initialize()

	if err := validateNewHttpUserPoolAuthorizerParameters(id, pool, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpUserPoolAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		[]interface{}{id, pool, props},
		&j,
	)

	return &j
}

// Initialize a Cognito user pool authorizer to be bound with HTTP route.
func NewHttpUserPoolAuthorizer_Override(h HttpUserPoolAuthorizer, id *string, pool awscognito.IUserPool, props *HttpUserPoolAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpUserPoolAuthorizer",
		[]interface{}{id, pool, props},
		h,
	)
}

func (h *jsiiProxy_HttpUserPoolAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
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

