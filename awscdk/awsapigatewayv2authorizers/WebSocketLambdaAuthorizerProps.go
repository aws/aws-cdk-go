package awsapigatewayv2authorizers


// Properties to initialize WebSocketTokenAuthorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketLambdaAuthorizerProps := &WebSocketLambdaAuthorizerProps{
//   	AuthorizerName: jsii.String("authorizerName"),
//   	IdentitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   }
//
type WebSocketLambdaAuthorizerProps struct {
	// The name of the authorizer.
	// Default: - same value as `id` passed in the constructor.
	//
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	//
	// Request parameter match `'route.request.querystring|header.[a-zA-z0-9._-]+'`.
	// Staged variable match `'stageVariables.[a-zA-Z0-9._-]+'`.
	// Context parameter match `'context.[a-zA-Z0-9._-]+'`.
	// Default: ['route.request.header.Authorization']
	//
	IdentitySource *[]*string `field:"optional" json:"identitySource" yaml:"identitySource"`
}

