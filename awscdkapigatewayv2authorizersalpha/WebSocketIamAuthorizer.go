// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/internal"
)

// Authorize WebSocket API Routes with IAM.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   // This function handles your connect route
//   var connectHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"))
//
//   webSocketApi.AddRoute(jsii.String("$connect"), &WebSocketRouteOptions{
//   	Integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("Integration"), connectHandler),
//   	Authorizer: awscdkapigatewayv2authorizersalpha.NewWebSocketIamAuthorizer(),
//   })
//
//   // Create an IAM user (identity)
//   user := iam.NewUser(this, jsii.String("User"))
//
//   webSocketArn := awscdk.stack_Of(this).FormatArn(&ArnComponents{
//   	Service: jsii.String("execute-api"),
//   	Resource: webSocketApi.ApiId,
//   })
//
//   // Grant access to the IAM user
//   user.AttachInlinePolicy(iam.NewPolicy(this, jsii.String("AllowInvoke"), &PolicyProps{
//   	Statements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("execute-api:Invoke"),
//   			},
//   			Effect: iam.Effect_ALLOW,
//   			Resources: []*string{
//   				webSocketArn,
//   			},
//   		}),
//   	},
//   }))
//
// Experimental.
type WebSocketIamAuthorizer interface {
	awscdkapigatewayv2alpha.IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	// Experimental.
	Bind(_options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketIamAuthorizer
type jsiiProxy_WebSocketIamAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIWebSocketRouteAuthorizer
}

// Experimental.
func NewWebSocketIamAuthorizer() WebSocketIamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_WebSocketIamAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketIamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketIamAuthorizer_Override(w WebSocketIamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketIamAuthorizer",
		nil, // no parameters
		w,
	)
}

func (w *jsiiProxy_WebSocketIamAuthorizer) Bind(_options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig {
	if err := w.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

