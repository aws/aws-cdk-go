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
//
//   	// the properties below are optional
//   	ClientSecret: jsii.String("clientSecret"),
//   	ClientSecretConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	ClientSecretSource: jsii.String("clientSecretSource"),
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
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput-clientsecretconfig
	//
	ClientSecretConfig interface{} `field:"optional" json:"clientSecretConfig" yaml:"clientSecretConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput-clientsecretsource
	//
	ClientSecretSource *string `field:"optional" json:"clientSecretSource" yaml:"clientSecretSource"`
	// The Microsoft Entra ID tenant ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-microsoftoauth2providerconfiginput-tenantid
	//
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

