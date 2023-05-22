package awsappflow


// The OAuth credentials required for OAuth type authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthCredentialsProperty := &OAuthCredentialsProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   		AuthCode: jsii.String("authCode"),
//   		RedirectUri: jsii.String("redirectUri"),
//   	},
//   	RefreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_OAuthCredentialsProperty struct {
	// The access token used to access protected SAPOData resources.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The identifier for the desired client.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// `CfnConnectorProfile.OAuthCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// The refresh token used to refresh expired access token.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

