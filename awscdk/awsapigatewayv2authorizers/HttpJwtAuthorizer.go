package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers/internal"
)

// Authorize Http Api routes on whether the requester is registered as part of an AWS Cognito user pool.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   issuer := "https://test.us.auth0.com"
//   authorizer := awscdk.NewHttpJwtAuthorizer(jsii.String("BooksAuthorizer"), issuer, &HttpJwtAuthorizerProps{
//   	JwtAudience: []*string{
//   		jsii.String("3131231"),
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
type HttpJwtAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// The authorizationType used for JWT Authorizer.
	AuthorizationType() *string
	// Return the id of the authorizer if it's been constructed.
	AuthorizerId() *string
	// Bind this authorizer to a specified Http route.
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpJwtAuthorizer
type jsiiProxy_HttpJwtAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

func (j *jsiiProxy_HttpJwtAuthorizer) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpJwtAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}


// Initialize a JWT authorizer to be bound with HTTP route.
func NewHttpJwtAuthorizer(id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) HttpJwtAuthorizer {
	_init_.Initialize()

	if err := validateNewHttpJwtAuthorizerParameters(id, jwtIssuer, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpJwtAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		[]interface{}{id, jwtIssuer, props},
		&j,
	)

	return &j
}

// Initialize a JWT authorizer to be bound with HTTP route.
func NewHttpJwtAuthorizer_Override(h HttpJwtAuthorizer, id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpJwtAuthorizer",
		[]interface{}{id, jwtIssuer, props},
		h,
	)
}

func (h *jsiiProxy_HttpJwtAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
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

