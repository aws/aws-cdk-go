package awsappflow


// The connector-specific profile credentials required by Google Analytics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   googleAnalyticsConnectorProfileCredentialsProperty := &GoogleAnalyticsConnectorProfileCredentialsProperty{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	AccessToken: jsii.String("accessToken"),
//   	ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   		AuthCode: jsii.String("authCode"),
//   		RedirectUri: jsii.String("redirectUri"),
//   	},
//   	RefreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_GoogleAnalyticsConnectorProfileCredentialsProperty struct {
	// The identifier for the desired client.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The credentials used to access protected Google Analytics resources.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// The credentials used to acquire new access tokens.
	//
	// This is required only for OAuth2 access tokens, and is not required for OAuth1 access tokens.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

