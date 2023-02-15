// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/internal"
)

// Authorize HTTP API Routes with IAM.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var principal anyPrincipal
//
//
//   authorizer := awscdkapigatewayv2authorizersalpha.NewHttpIamAuthorizer()
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"), &httpApiProps{
//   	defaultAuthorizer: authorizer,
//   })
//
//   routes := httpApi.addRoutes(&addRoutesOptions{
//   	integration: awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.example.com")),
//   	path: jsii.String("/books/{book}"),
//   })
//
//   routes[0].grantInvoke(principal)
//
// Experimental.
type HttpIamAuthorizer interface {
	awscdkapigatewayv2alpha.IHttpRouteAuthorizer
	// Bind this authorizer to a specified Http route.
	// Experimental.
	Bind(_options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig
}

// The jsii proxy struct for HttpIamAuthorizer
type jsiiProxy_HttpIamAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteAuthorizer
}

// Experimental.
func NewHttpIamAuthorizer() HttpIamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_HttpIamAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpIamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewHttpIamAuthorizer_Override(h HttpIamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.HttpIamAuthorizer",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_HttpIamAuthorizer) Bind(_options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig {
	if err := h.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.HttpRouteAuthorizerConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

