package awsbedrockagentcore


// Input configuration for a Slack OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   slackOauth2ProviderConfigInputProperty := &SlackOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-slackoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProviderPropsMixin_SlackOauth2ProviderConfigInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-slackoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-slackoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-slackoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-slackoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

