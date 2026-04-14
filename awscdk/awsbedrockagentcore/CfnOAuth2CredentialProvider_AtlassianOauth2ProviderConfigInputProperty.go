package awsbedrockagentcore


// Input configuration for an Atlassian OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   atlassianOauth2ProviderConfigInputProperty := &AtlassianOauth2ProviderConfigInputProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-atlassianoauth2providerconfiginput.html
//
type CfnOAuth2CredentialProvider_AtlassianOauth2ProviderConfigInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-atlassianoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-atlassianoauth2providerconfiginput-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-atlassianoauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-atlassianoauth2providerconfiginput-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

