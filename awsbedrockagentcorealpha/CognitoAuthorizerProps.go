package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Properties for configuring a Cognito authorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gatewayCustomClaim GatewayCustomClaim
//   var userPool UserPool
//   var userPoolClient UserPoolClient
//
//   cognitoAuthorizerProps := &CognitoAuthorizerProps{
//   	UserPool: userPool,
//
//   	// the properties below are optional
//   	AllowedAudiences: []*string{
//   		jsii.String("allowedAudiences"),
//   	},
//   	AllowedClients: []IUserPoolClient{
//   		userPoolClient,
//   	},
//   	AllowedScopes: []*string{
//   		jsii.String("allowedScopes"),
//   	},
//   	CustomClaims: []GatewayCustomClaim{
//   		gatewayCustomClaim,
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type CognitoAuthorizerProps struct {
	// The Cognito User Pool to use for authentication.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// The allowed audiences for JWT validation.
	// Default: - No audience validation.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The allowed User Pool clients.
	// Default: - All clients are allowed.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedClients *[]awscognito.IUserPoolClient `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// The allowed scopes for JWT validation.
	// Default: - No scope validation.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Custom claims for additional JWT token validation.
	//
	// Allows you to validate additional fields in JWT tokens beyond the standard audience, client, and scope validations.
	// Default: - No custom claim validation.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CustomClaims *[]GatewayCustomClaim `field:"optional" json:"customClaims" yaml:"customClaims"`
}

