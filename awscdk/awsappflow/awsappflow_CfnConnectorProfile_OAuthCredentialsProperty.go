package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuthCredentialsProperty := &oAuthCredentialsProperty{
//   	accessToken: jsii.String("accessToken"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_OAuthCredentialsProperty struct {
	// `CfnConnectorProfile.OAuthCredentialsProperty.AccessToken`.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// `CfnConnectorProfile.OAuthCredentialsProperty.ClientId`.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// `CfnConnectorProfile.OAuthCredentialsProperty.ClientSecret`.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// `CfnConnectorProfile.OAuthCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// `CfnConnectorProfile.OAuthCredentialsProperty.RefreshToken`.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

