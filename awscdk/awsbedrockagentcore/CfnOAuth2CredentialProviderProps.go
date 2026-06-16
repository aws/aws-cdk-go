package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOAuth2CredentialProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOAuth2CredentialProviderProps := &CfnOAuth2CredentialProviderProps{
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html
//
type CfnOAuth2CredentialProviderProps struct {
	// The vendor of the OAuth2 credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-credentialprovidervendor
	//
	CredentialProviderVendor *string `field:"required" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// The name of the OAuth2 credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Input configuration for an OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput
	//
	Oauth2ProviderConfigInput interface{} `field:"optional" json:"oauth2ProviderConfigInput" yaml:"oauth2ProviderConfigInput"`
	// Tags to assign to the OAuth2 credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

