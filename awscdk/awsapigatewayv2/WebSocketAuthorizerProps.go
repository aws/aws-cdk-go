package awsapigatewayv2


// Properties to initialize an instance of `WebSocketAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var webSocketApi webSocketApi
//
//   webSocketAuthorizerProps := &WebSocketAuthorizerProps{
//   	IdentitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	Type: awscdk.Aws_apigatewayv2.WebSocketAuthorizerType_LAMBDA,
//   	WebSocketApi: webSocketApi,
//
//   	// the properties below are optional
//   	AuthorizerName: jsii.String("authorizerName"),
//   	AuthorizerUri: jsii.String("authorizerUri"),
//   }
//
type WebSocketAuthorizerProps struct {
	// The identity source for which authorization is requested.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identitysource
	//
	IdentitySource *[]*string `field:"required" json:"identitySource" yaml:"identitySource"`
	// The type of authorizer.
	Type WebSocketAuthorizerType `field:"required" json:"type" yaml:"type"`
	// WebSocket Api to attach the authorizer to.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
	// Name of the authorizer.
	// Default: - id of the WebSocketAuthorizer construct.
	//
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// For REQUEST authorizers, this must be a well-formed Lambda function URI.
	// Default: - required for Request authorizer types.
	//
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
}

