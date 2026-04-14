package awsbedrockagentcore


// Input configuration for a supported non-custom OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   includedOauth2ProviderConfigInputProperty := &IncludedOauth2ProviderConfigInputProperty{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	Issuer: jsii.String("issuer"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProviderPropsMixin_IncludedOauth2ProviderConfigInputProperty struct {
	// OAuth2 authorization endpoint for your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Token issuer of your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// OAuth2 token endpoint for your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-tokenendpoint
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

