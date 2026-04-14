package awsbedrockagentcore


// Input configuration for a Salesforce OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceOauth2ProviderConfigInputProperty := &SalesforceOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProvider_SalesforceOauth2ProviderConfigInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

