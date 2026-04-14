package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOAuth2CredentialProviderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnOAuth2CredentialProviderMixinProps := &CfnOAuth2CredentialProviderMixinProps{
//   	CredentialProviderVendor: jsii.String("credentialProviderVendor"),
//   	Name: jsii.String("name"),
//   	Oauth2ProviderConfigInput: &Oauth2ProviderConfigInputProperty{
//   		AtlassianOauth2ProviderConfig: &AtlassianOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		CustomOauth2ProviderConfig: &CustomOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			OauthDiscovery: &Oauth2DiscoveryProperty{
//   				AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   					AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   					Issuer: jsii.String("issuer"),
//   					ResponseTypes: []*string{
//   						jsii.String("responseTypes"),
//   					},
//   					TokenEndpoint: jsii.String("tokenEndpoint"),
//   				},
//   				DiscoveryUrl: jsii.String("discoveryUrl"),
//   			},
//   		},
//   		GithubOauth2ProviderConfig: &GithubOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		GoogleOauth2ProviderConfig: &GoogleOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		IncludedOauth2ProviderConfig: &IncludedOauth2ProviderConfigInputProperty{
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			Issuer: jsii.String("issuer"),
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   		LinkedinOauth2ProviderConfig: &LinkedinOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		MicrosoftOauth2ProviderConfig: &MicrosoftOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   			TenantId: jsii.String("tenantId"),
//   		},
//   		SalesforceOauth2ProviderConfig: &SalesforceOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
//   		},
//   		SlackOauth2ProviderConfig: &SlackOauth2ProviderConfigInputProperty{
//   			ClientId: jsii.String("clientId"),
//   			ClientSecret: jsii.String("clientSecret"),
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
type CfnOAuth2CredentialProviderMixinProps struct {
	// The vendor of the OAuth2 credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-credentialprovidervendor
	//
	CredentialProviderVendor *string `field:"optional" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// The name of the OAuth2 credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Input configuration for an OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput
	//
	Oauth2ProviderConfigInput interface{} `field:"optional" json:"oauth2ProviderConfigInput" yaml:"oauth2ProviderConfigInput"`
	// Tags to assign to the OAuth2 credential provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html#cfn-bedrockagentcore-oauth2credentialprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

