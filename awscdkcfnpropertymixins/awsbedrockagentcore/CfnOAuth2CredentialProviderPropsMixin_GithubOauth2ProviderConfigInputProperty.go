package awsbedrockagentcore


// Input configuration for a GitHub OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   githubOauth2ProviderConfigInputProperty := &GithubOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-githuboauth2providerconfiginput.html
//
type CfnOAuth2CredentialProviderPropsMixin_GithubOauth2ProviderConfigInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-githuboauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-githuboauth2providerconfiginput-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-githuboauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-githuboauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

