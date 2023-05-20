package awscdkapigatewayv2authorizersalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/internal"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   authorizer := awscdkapigatewayv2authorizersalpha.NewHttpUserPoolAuthorizer(jsii.String("BooksAuthorizer"), userPool)
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books"),
//   	Authorizer: Authorizer,
//   })
//
// Experimental.
type HttpUserPoolAuthorizer interface {
	awscdkapigatewayv2alpha.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpUserPoolAuthorizer
type jsiiProxy_HttpUserPoolAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer
}

// Initialize a Cognito user pool authorizer to be bound with HTTP route.
// Experimental.
func NewHttpUserPoolAuthorizer(id *string, pool awscognito.IUserPool, props *HttpUserPoolAuthorizerProps) HttpUserPoolAuthorizer {
	_init_.Initialize()

	if err := validateNewHttpUserPoolAuthorizerParameters(id, pool, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpUserPoolAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpUserPoolAuthorizer",
		[]interface{}{id, pool, props},
		&j,
	)

	return &j
}

// Initialize a Cognito user pool authorizer to be bound with HTTP route.
// Experimental.
func NewHttpUserPoolAuthorizer_Override(h HttpUserPoolAuthorizer, id *string, pool awscognito.IUserPool, props *HttpUserPoolAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpUserPoolAuthorizer",
		[]interface{}{id, pool, props},
		h,
	)
}

func (h *jsiiProxy_HttpUserPoolAuthorizer) Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig {
	if err := h.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

