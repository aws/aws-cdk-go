package awsbedrockagentcorealpha


// Built-in OAuth2 vendors supported by `AWS::BedrockAgentCore::OAuth2CredentialProvider`.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-oauth2credentialprovider.html
//
// Experimental.
type OAuth2CredentialProviderVendor string

const (
	// Google OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_GOOGLE OAuth2CredentialProviderVendor = "GOOGLE"
	// GitHub OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_GITHUB OAuth2CredentialProviderVendor = "GITHUB"
	// Slack OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_SLACK OAuth2CredentialProviderVendor = "SLACK"
	// Salesforce OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_SALESFORCE OAuth2CredentialProviderVendor = "SALESFORCE"
	// Microsoft OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_MICROSOFT OAuth2CredentialProviderVendor = "MICROSOFT"
	// Custom OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_CUSTOM OAuth2CredentialProviderVendor = "CUSTOM"
	// Atlassian OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_ATLASSIAN OAuth2CredentialProviderVendor = "ATLASSIAN"
	// LinkedIn OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_LINKEDIN OAuth2CredentialProviderVendor = "LINKEDIN"
	// X (Twitter) OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_X OAuth2CredentialProviderVendor = "X"
	// Okta OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_OKTA OAuth2CredentialProviderVendor = "OKTA"
	// OneLogin OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_ONE_LOGIN OAuth2CredentialProviderVendor = "ONE_LOGIN"
	// PingOne OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_PING_ONE OAuth2CredentialProviderVendor = "PING_ONE"
	// Facebook OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_FACEBOOK OAuth2CredentialProviderVendor = "FACEBOOK"
	// Yandex OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_YANDEX OAuth2CredentialProviderVendor = "YANDEX"
	// Reddit OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_REDDIT OAuth2CredentialProviderVendor = "REDDIT"
	// Zoom OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_ZOOM OAuth2CredentialProviderVendor = "ZOOM"
	// Twitch OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_TWITCH OAuth2CredentialProviderVendor = "TWITCH"
	// Spotify OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_SPOTIFY OAuth2CredentialProviderVendor = "SPOTIFY"
	// Dropbox OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_DROPBOX OAuth2CredentialProviderVendor = "DROPBOX"
	// Notion OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_NOTION OAuth2CredentialProviderVendor = "NOTION"
	// HubSpot OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_HUBSPOT OAuth2CredentialProviderVendor = "HUBSPOT"
	// CyberArk OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_CYBER_ARK OAuth2CredentialProviderVendor = "CYBER_ARK"
	// FusionAuth OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_FUSION_AUTH OAuth2CredentialProviderVendor = "FUSION_AUTH"
	// Auth0 OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_AUTH0 OAuth2CredentialProviderVendor = "AUTH0"
	// Amazon Cognito OAuth2.
	// Experimental.
	OAuth2CredentialProviderVendor_COGNITO OAuth2CredentialProviderVendor = "COGNITO"
)

