package awsbedrockagentcore


// Input configuration for an OAuth2 provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oauth2ProviderConfigInputProperty := &Oauth2ProviderConfigInputProperty{
//   	AtlassianOauth2ProviderConfig: &AtlassianOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   	CustomOauth2ProviderConfig: &CustomOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		OauthDiscovery: &Oauth2DiscoveryProperty{
//   			AuthorizationServerMetadata: &Oauth2AuthorizationServerMetadataProperty{
//   				AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				Issuer: jsii.String("issuer"),
//   				TokenEndpoint: jsii.String("tokenEndpoint"),
//
//   				// the properties below are optional
//   				ResponseTypes: []*string{
//   					jsii.String("responseTypes"),
//   				},
//   			},
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//   		},
//   	},
//   	GithubOauth2ProviderConfig: &GithubOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   	GoogleOauth2ProviderConfig: &GoogleOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   	IncludedOauth2ProviderConfig: &IncludedOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		Issuer: jsii.String("issuer"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   	},
//   	LinkedinOauth2ProviderConfig: &LinkedinOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   	MicrosoftOauth2ProviderConfig: &MicrosoftOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//
//   		// the properties below are optional
//   		TenantId: jsii.String("tenantId"),
//   	},
//   	SalesforceOauth2ProviderConfig: &SalesforceOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   	SlackOauth2ProviderConfig: &SlackOauth2ProviderConfigInputProperty{
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html
//
type CfnOAuth2CredentialProvider_Oauth2ProviderConfigInputProperty struct {
	// Input configuration for an Atlassian OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-atlassianoauth2providerconfig
	//
	AtlassianOauth2ProviderConfig interface{} `field:"optional" json:"atlassianOauth2ProviderConfig" yaml:"atlassianOauth2ProviderConfig"`
	// Input configuration for a custom OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-customoauth2providerconfig
	//
	CustomOauth2ProviderConfig interface{} `field:"optional" json:"customOauth2ProviderConfig" yaml:"customOauth2ProviderConfig"`
	// Input configuration for a GitHub OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-githuboauth2providerconfig
	//
	GithubOauth2ProviderConfig interface{} `field:"optional" json:"githubOauth2ProviderConfig" yaml:"githubOauth2ProviderConfig"`
	// Input configuration for a Google OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-googleoauth2providerconfig
	//
	GoogleOauth2ProviderConfig interface{} `field:"optional" json:"googleOauth2ProviderConfig" yaml:"googleOauth2ProviderConfig"`
	// Input configuration for a supported non-custom OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-includedoauth2providerconfig
	//
	IncludedOauth2ProviderConfig interface{} `field:"optional" json:"includedOauth2ProviderConfig" yaml:"includedOauth2ProviderConfig"`
	// Input configuration for a LinkedIn OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-linkedinoauth2providerconfig
	//
	LinkedinOauth2ProviderConfig interface{} `field:"optional" json:"linkedinOauth2ProviderConfig" yaml:"linkedinOauth2ProviderConfig"`
	// Input configuration for a Microsoft OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-microsoftoauth2providerconfig
	//
	MicrosoftOauth2ProviderConfig interface{} `field:"optional" json:"microsoftOauth2ProviderConfig" yaml:"microsoftOauth2ProviderConfig"`
	// Input configuration for a Salesforce OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-salesforceoauth2providerconfig
	//
	SalesforceOauth2ProviderConfig interface{} `field:"optional" json:"salesforceOauth2ProviderConfig" yaml:"salesforceOauth2ProviderConfig"`
	// Input configuration for a Slack OAuth2 provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput.html#cfn-bedrockagentcore-oauth2credentialprovider-oauth2providerconfiginput-slackoauth2providerconfig
	//
	SlackOauth2ProviderConfig interface{} `field:"optional" json:"slackOauth2ProviderConfig" yaml:"slackOauth2ProviderConfig"`
}

