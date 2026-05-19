package awsbedrockagentcorealpha


// OAuth configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   oAuthConfiguration := &OAuthConfiguration{
//   	ProviderArn: jsii.String("providerArn"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   	SecretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	CustomParameters: map[string]*string{
//   		"customParametersKey": jsii.String("customParameters"),
//   	},
//   }
//
// Experimental.
type OAuthConfiguration struct {
	// The OAuth credential provider ARN.
	//
	// This is returned when creating the OAuth credential provider via Console or API.
	// Format: arn:aws:bedrock-agentcore:region:account:token-vault/id/oauth2credentialprovider/name
	// Required: Yes.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// The OAuth scopes for the credential provider. These scopes define the level of access requested from the OAuth provider.
	//
	// Array Members: Minimum number of 0 items. Maximum number of 100 items.
	// Length Constraints: Minimum length of 1. Maximum length of 64.
	// Required: Yes.
	// Experimental.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The ARN of the Secrets Manager secret containing OAuth credentials (client ID and secret).
	//
	// This is returned when creating the OAuth credential provider via Console or API.
	// Format: arn:aws:secretsmanager:region:account:secret:name
	// Required: Yes.
	// Experimental.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Custom parameters for the OAuth flow.
	// Default: - No custom parameters.
	//
	// Experimental.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
}

