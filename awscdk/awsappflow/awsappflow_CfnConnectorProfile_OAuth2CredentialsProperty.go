package awsappflow


// The OAuth 2.0 credentials required for OAuth 2.0 authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuth2CredentialsProperty := &oAuth2CredentialsProperty{
//   	accessToken: jsii.String("accessToken"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	oAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_OAuth2CredentialsProperty struct {
	// The access token used to access the connector on your behalf.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The identifier for the desired client.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// `CfnConnectorProfile.OAuth2CredentialsProperty.OAuthRequest`.
	OAuthRequest interface{} `field:"optional" json:"oAuthRequest" yaml:"oAuthRequest"`
	// The refresh token used to refresh an expired access token.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

