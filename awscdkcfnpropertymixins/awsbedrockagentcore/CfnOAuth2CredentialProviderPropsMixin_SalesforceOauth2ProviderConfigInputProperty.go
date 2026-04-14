package awsbedrockagentcore


// Input configuration for a Salesforce OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   salesforceOauth2ProviderConfigInputProperty := &SalesforceOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProviderPropsMixin_SalesforceOauth2ProviderConfigInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-salesforceoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

