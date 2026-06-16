package awsbedrockagentcore


// Input configuration for a custom OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customOauth2ProviderConfigInputProperty := &CustomOauth2ProviderConfigInputProperty{
//   	OauthDiscovery: &Oauth2DiscoveryProperty{
//   		AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			Issuer: jsii.String("issuer"),
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//
//   			// the properties below are optional
//   			ResponseTypes: []*string{
//   				jsii.String("responseTypes"),
//   			},
//   		},
//   		DiscoveryUrl: jsii.String("discoveryUrl"),
//   	},
//
//   	// the properties below are optional
//   	ClientAuthenticationMethod: jsii.String("clientAuthenticationMethod"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	ClientSecretConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	ClientSecretSource: jsii.String("clientSecretSource"),
//   	OnBehalfOfTokenExchangeConfig: &OnBehalfOfTokenExchangeConfigProperty{
//   		GrantType: jsii.String("grantType"),
//
//   		// the properties below are optional
//   		TokenExchangeGrantTypeConfig: &TokenExchangeGrantTypeConfigProperty{
//   			ActorTokenContent: jsii.String("actorTokenContent"),
//
//   			// the properties below are optional
//   			ActorTokenScopes: []*string{
//   				jsii.String("actorTokenScopes"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProvider_CustomOauth2ProviderConfigInputProperty struct {
	// Discovery information for an OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-oauthdiscovery
	//
	OauthDiscovery interface{} `field:"required" json:"oauthDiscovery" yaml:"oauthDiscovery"`
	// The client authentication method to use when authenticating with the token endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-clientauthenticationmethod
	//
	ClientAuthenticationMethod *string `field:"optional" json:"clientAuthenticationMethod" yaml:"clientAuthenticationMethod"`
	// The client ID for the custom OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret for the custom OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-clientsecretconfig
	//
	ClientSecretConfig interface{} `field:"optional" json:"clientSecretConfig" yaml:"clientSecretConfig"`
	// The source of the client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-clientsecretsource
	//
	ClientSecretSource *string `field:"optional" json:"clientSecretSource" yaml:"clientSecretSource"`
	// Configuration for on-behalf-of token exchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-onbehalfoftokenexchangeconfig
	//
	OnBehalfOfTokenExchangeConfig interface{} `field:"optional" json:"onBehalfOfTokenExchangeConfig" yaml:"onBehalfOfTokenExchangeConfig"`
}

