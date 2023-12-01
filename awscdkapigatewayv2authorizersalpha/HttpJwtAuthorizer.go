package awscdkapigatewayv2authorizersalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/internal"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   issuer := "https://test.us.auth0.com"
//   authorizer := awscdkapigatewayv2authorizersalpha.NewHttpJwtAuthorizer(jsii.String("BooksAuthorizer"), issuer, &HttpJwtAuthorizerProps{
//   	JwtAudience: []*string{
//   		jsii.String("3131231"),
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books"),
//   	Authorizer: Authorizer,
//   })
//
// Deprecated.
type HttpJwtAuthorizer interface {
	awscdkapigatewayv2alpha.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Deprecated.
	Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpJwtAuthorizer
type jsiiProxy_HttpJwtAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer
}

// Initialize a JWT authorizer to be bound with HTTP route.
// Deprecated.
func NewHttpJwtAuthorizer(id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) HttpJwtAuthorizer {
	_init_.Initialize()

	if err := validateNewHttpJwtAuthorizerParameters(id, jwtIssuer, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpJwtAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpJwtAuthorizer",
		[]interface{}{id, jwtIssuer, props},
		&j,
	)

	return &j
}

// Initialize a JWT authorizer to be bound with HTTP route.
// Deprecated.
func NewHttpJwtAuthorizer_Override(h HttpJwtAuthorizer, id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpJwtAuthorizer",
		[]interface{}{id, jwtIssuer, props},
		h,
	)
}

func (h *jsiiProxy_HttpJwtAuthorizer) Bind(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig {
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

