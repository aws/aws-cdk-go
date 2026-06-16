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
//
//   	// the properties below are optional
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	ClientSecretConfig: &SecretReferenceProperty{
//   		JsonKey: jsii.String("jsonKey"),
//   		SecretId: jsii.String("secretId"),
//   	},
//   	ClientSecretSource: jsii.String("clientSecretSource"),
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
	// OAuth2 authorization endpoint for your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A reference to a customer-provided secret stored in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-clientsecretconfig
	//
	ClientSecretConfig interface{} `field:"optional" json:"clientSecretConfig" yaml:"clientSecretConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-clientsecretsource
	//
	ClientSecretSource *string `field:"optional" json:"clientSecretSource" yaml:"clientSecretSource"`
	// Token issuer of your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// OAuth2 token endpoint for your isolated OAuth2 application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-includedoauth2providerconfiginput-tokenendpoint
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

