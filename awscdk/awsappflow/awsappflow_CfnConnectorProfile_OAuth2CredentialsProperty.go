package awsappflow


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
	// `CfnConnectorProfile.OAuth2CredentialsProperty.AccessToken`.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// `CfnConnectorProfile.OAuth2CredentialsProperty.ClientId`.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// `CfnConnectorProfile.OAuth2CredentialsProperty.ClientSecret`.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// `CfnConnectorProfile.OAuth2CredentialsProperty.OAuthRequest`.
	OAuthRequest interface{} `field:"optional" json:"oAuthRequest" yaml:"oAuthRequest"`
	// `CfnConnectorProfile.OAuth2CredentialsProperty.RefreshToken`.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

