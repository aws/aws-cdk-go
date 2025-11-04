package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers/internal"
)

// Authorize WebSocket API Routes with IAM.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // This function handles your connect route
//   var connectHandler Function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"))
//
//   webSocketApi.AddRoute(jsii.String("$connect"), &WebSocketRouteOptions{
//   	Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("Integration"), connectHandler),
//   	Authorizer: awscdk.NewWebSocketIamAuthorizer(),
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
//   	Statements: []PolicyStatement{
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
type WebSocketIamAuthorizer interface {
	awsapigatewayv2.IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	Bind(options *awsapigatewayv2.WebSocketRouteAuthorizerBindOptions) *awsapigatewayv2.WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketIamAuthorizer
type jsiiProxy_WebSocketIamAuthorizer struct {
	internal.Type__awsapigatewayv2IWebSocketRouteAuthorizer
}

func NewWebSocketIamAuthorizer() WebSocketIamAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_WebSocketIamAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.WebSocketIamAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewWebSocketIamAuthorizer_Override(w WebSocketIamAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.WebSocketIamAuthorizer",
		nil, // no parameters
		w,
	)
}

func (w *jsiiProxy_WebSocketIamAuthorizer) Bind(options *awsapigatewayv2.WebSocketRouteAuthorizerBindOptions) *awsapigatewayv2.WebSocketRouteAuthorizerConfig {
	if err := w.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

