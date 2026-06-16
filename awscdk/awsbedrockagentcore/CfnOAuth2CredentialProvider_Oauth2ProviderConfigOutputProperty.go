package awsbedrockagentcore


// Output configuration for an OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oauth2ProviderConfigOutputProperty := &Oauth2ProviderConfigOutputProperty{
//   	ClientAuthenticationMethod: jsii.String("clientAuthenticationMethod"),
//   	ClientId: jsii.String("clientId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html
//
type CfnOAuth2CredentialProvider_Oauth2ProviderConfigOutputProperty struct {
	// The client authentication method used when authenticating with the token endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-clientauthenticationmethod
	//
	ClientAuthenticationMethod *string `field:"optional" json:"clientAuthenticationMethod" yaml:"clientAuthenticationMethod"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Discovery information for an OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-oauthdiscovery
	//
	OauthDiscovery interface{} `field:"optional" json:"oauthDiscovery" yaml:"oauthDiscovery"`
	// Configuration for on-behalf-of token exchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-onbehalfoftokenexchangeconfig
	//
	OnBehalfOfTokenExchangeConfig interface{} `field:"optional" json:"onBehalfOfTokenExchangeConfig" yaml:"onBehalfOfTokenExchangeConfig"`
}

