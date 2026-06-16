package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore"
)

// Low-level properties when you need full control (prefer {@link OAuth2CredentialProvider.usingSlack} and other factories).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   oAuth2CredentialProviderProps := &OAuth2CredentialProviderProps{
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//   	Oauth2ProviderConfigInput: &Oauth2ProviderConfigInputProperty{
//   		AtlassianOauth2ProviderConfig: &AtlassianOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		CustomOauth2ProviderConfig: &CustomOauth2ProviderConfigInputProperty{
//   			OauthDiscovery: &Oauth2DiscoveryProperty{
//   				AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   					AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   					Issuer: jsii.String("issuer"),
//   					TokenEndpoint: jsii.String("tokenEndpoint"),
//
//   					// the properties below are optional
//   					ResponseTypes: []*string{
//   						jsii.String("responseTypes"),
//   					},
//   				},
//   				DiscoveryUrl: jsii.String("discoveryUrl"),
//   			},
//
//   			// the properties below are optional
//   			ClientAuthenticationMethod: jsii.String("clientAuthenticationMethod"),
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   			OnBehalfOfTokenExchangeConfig: &OnBehalfOfTokenExchangeConfigProperty{
//   				GrantType: jsii.String("grantType"),
//
//   				// the properties below are optional
//   				TokenExchangeGrantTypeConfig: &TokenExchangeGrantTypeConfigProperty{
//   					ActorTokenContent: jsii.String("actorTokenContent"),
//
//   					// the properties below are optional
//   					ActorTokenScopes: []*string{
//   						jsii.String("actorTokenScopes"),
//   					},
//   				},
//   			},
//   		},
//   		GithubOauth2ProviderConfig: &GithubOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		GoogleOauth2ProviderConfig: &GoogleOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		IncludedOauth2ProviderConfig: &IncludedOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   			Issuer: jsii.String("issuer"),
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   		LinkedinOauth2ProviderConfig: &LinkedinOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		MicrosoftOauth2ProviderConfig: &MicrosoftOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   			TenantId: jsii.String("tenantId"),
//   		},
//   		SalesforceOauth2ProviderConfig: &SalesforceOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   		SlackOauth2ProviderConfig: &SlackOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//
//   			// the properties below are optional
//   			ClientSecret: jsii.String("clientSecret"),
//   			ClientSecretConfig: &SecretReferenceProperty{
//   				JsonKey: jsii.String("jsonKey"),
//   				SecretId: jsii.String("secretId"),
//   			},
//   			ClientSecretSource: jsii.String("clientSecretSource"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	OAuth2CredentialProviderName: jsii.String("oAuth2CredentialProviderName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type OAuth2CredentialProviderProps struct {
	// OAuth2 vendor string for CloudFormation `CredentialProviderVendor`.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderVendor *string `field:"required" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// OAuth2 provider configuration passed through to `Oauth2ProviderConfigInput`.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Oauth2ProviderConfigInput *awsbedrockagentcore.CfnOAuth2CredentialProvider_Oauth2ProviderConfigInputProperty `field:"required" json:"oauth2ProviderConfigInput" yaml:"oauth2ProviderConfigInput"`
	// Name of the credential provider.
	// Default: a name generated by CDK.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	OAuth2CredentialProviderName *string `field:"optional" json:"oAuth2CredentialProviderName" yaml:"oAuth2CredentialProviderName"`
	// Tags for this credential provider.
	// Default: - no tags.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

