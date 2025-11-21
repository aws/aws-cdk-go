package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// ****************************************************************************                              Factory ***************************************************************************.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
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
//   }
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
}

