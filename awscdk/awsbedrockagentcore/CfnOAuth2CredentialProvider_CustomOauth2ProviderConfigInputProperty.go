package awsbedrockagentcore


// Input configuration for a custom OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customOauth2ProviderConfigInputProperty := &CustomOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProvider_CustomOauth2ProviderConfigInputProperty struct {
	// The client ID for the custom OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret for the custom OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Discovery information for an OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-customoauth2providerconfiginput-oauthdiscovery
	//
	OauthDiscovery interface{} `field:"required" json:"oauthDiscovery" yaml:"oauthDiscovery"`
}

