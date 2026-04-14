package awsbedrockagentcore


// Input configuration for a Microsoft OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microsoftOauth2ProviderConfigInputProperty := &MicrosoftOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	TenantId: jsii.String("tenantId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProvider_MicrosoftOauth2ProviderConfigInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The Microsoft Entra ID tenant ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput-tenantid
	//
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

