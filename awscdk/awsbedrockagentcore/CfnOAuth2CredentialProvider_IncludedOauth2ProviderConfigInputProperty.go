package awsbedrockagentcore


// Input configuration for a supported non-custom OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   includedOauth2ProviderConfigInputProperty := &IncludedOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	Issuer: jsii.String("issuer"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProvider_IncludedOauth2ProviderConfigInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// OAuth2 authorization endpoint for your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Token issuer of your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// OAuth2 token endpoint for your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-tokenendpoint
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

