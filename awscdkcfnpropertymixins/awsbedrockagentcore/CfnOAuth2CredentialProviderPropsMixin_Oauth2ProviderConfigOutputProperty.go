package awsbedrockagentcore


// Output configuration for an OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   oauth2ProviderConfigOutputProperty := &Oauth2ProviderConfigOutputProperty{
//   	ClientId: jsii.String("clientId"),
//   	OauthDiscovery: &Oauth2DiscoveryProperty{
//   		AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   			AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   			Issuer: jsii.String("issuer"),
//   			ResponseTypes: []*string{
//   				jsii.String("responseTypes"),
//   			},
//   			TokenEndpoint: jsii.String("tokenEndpoint"),
//   		},
//   		DiscoveryUrl: jsii.String("discoveryUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html
//
type CfnOAuth2CredentialProviderPropsMixin_Oauth2ProviderConfigOutputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Discovery information for an OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfigoutput-oauthdiscovery
	//
	OauthDiscovery interface{} `field:"optional" json:"oauthDiscovery" yaml:"oauthDiscovery"`
}

