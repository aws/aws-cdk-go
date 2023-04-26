package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2authorizers/internal"
)

// Authorize HTTP API Routes with IAM.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var principal anyPrincipal
//
//
//   authorizer := awscdk.NewHttpIamAuthorizer()
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &HttpApiProps{
//   	DefaultAuthorizer: authorizer,
//   })
//
//   routes := httpApi.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	Path: jsii.String("/books/{book}"),
//   })
//
//   routes[0].GrantInvoke(principal)
//
// Experimental.
type HttpIamAuthorizer interface {
	awsapigatewayv2.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(_options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpIamAuthorizer
type jsiiProxy_HttpIamAuthorizer struct {
	internal.Type__awsapigatewayv2IHttpRouteAuthorizer
}

// Experimental.
func NewHttpIamAuthorizer() HttpIamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpIamAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpIamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewHttpIamAuthorizer_Override(h HttpIamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_authorizers.HttpIamAuthorizer",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_HttpIamAuthorizer) Bind(_options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	if err := h.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

