package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// ****************************************************************************                              Factory ***************************************************************************.
//
// Example:
//   var userPool UserPool
//   var userPoolClient UserPoolClient
//
//
//   // Optional: Create custom claims (CustomClaimOperator and GatewayCustomClaim from agentcore)
//   customClaims := []GatewayCustomClaim{
//   	agentcore.GatewayCustomClaim_WithStringValue(jsii.String("department"), jsii.String("engineering")),
//   	agentcore.GatewayCustomClaim_WithStringArrayValue(jsii.String("roles"), []*string{
//   		jsii.String("admin"),
//   	}, agentcore.CustomClaimOperator_CONTAINS),
//   	agentcore.GatewayCustomClaim_WithStringArrayValue(jsii.String("permissions"), []*string{
//   		jsii.String("read"),
//   		jsii.String("write"),
//   	}, agentcore.CustomClaimOperator_CONTAINS_ANY),
//   }
//
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	AuthorizerConfiguration: agentcore.GatewayAuthorizer_UsingCognito(&CognitoAuthorizerProps{
//   		UserPool: userPool,
//   		AllowedClients: []IUserPoolClient{
//   			userPoolClient,
//   		},
//   		AllowedAudiences: []*string{
//   			jsii.String("audience1"),
//   		},
//   		AllowedScopes: []*string{
//   			jsii.String("read"),
//   			jsii.String("write"),
//   		},
//   		CustomClaims: customClaims,
//   	}),
//   })
//
// Experimental.
type CognitoAuthorizerProps struct {
	// The Cognito User Pool to use for authentication.
	// Experimental.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// The allowed audiences for JWT validation.
	// Default: - No audience validation.
	//
	// Experimental.
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The allowed User Pool clients.
	// Default: - All clients are allowed.
	//
	// Experimental.
	AllowedClients *[]awscognito.IUserPoolClient `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// The allowed scopes for JWT validation.
	// Default: - No scope validation.
	//
	// Experimental.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Custom claims for additional JWT token validation.
	//
	// Allows you to validate additional fields in JWT tokens beyond the standard audience, client, and scope validations.
	// Default: - No custom claim validation.
	//
	// Experimental.
	CustomClaims *[]GatewayCustomClaim `field:"optional" json:"customClaims" yaml:"customClaims"`
}

