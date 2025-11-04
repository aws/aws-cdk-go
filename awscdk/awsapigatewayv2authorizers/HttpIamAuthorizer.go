package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers/internal"
)

// Authorize HTTP API Routes with IAM.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var principal AnyPrincipal
//
//
//   authorizer := awscdk.NewHttpIamAuthorizer()
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
//   	DefaultAuthorizer: authorizer,
//   })
//
//   routes := httpApi.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	Path: jsii.String("/books/{book}"),
//   })
//
//   routes[0].GrantInvoke(principal)
//
type HttpIamAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// The authorizationType used for IAM Authorizer.
	AuthorizationType() awsapigatewayv2.HttpAuthorizerType
	// Bind this authorizer to a specified Http route.
	Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpIamAuthorizer
type jsiiProxy_HttpIamAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

func (j *jsiiProxy_HttpIamAuthorizer) AuthorizationType() awsapigatewayv2.HttpAuthorizerType {
	var returns awsapigatewayv2.HttpAuthorizerType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}


func NewHttpIamAuthorizer() HttpIamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpIamAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpIamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewHttpIamAuthorizer_Override(h HttpIamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.HttpIamAuthorizer",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_HttpIamAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
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

