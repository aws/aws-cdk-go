package awsappflow


// The connector-specific profile credentials required when using SAPOData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataConnectorProfileCredentialsProperty := &SAPODataConnectorProfileCredentialsProperty{
//   	BasicAuthCredentials: &BasicAuthCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	OAuthCredentials: &OAuthCredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   }
//
type CfnConnectorProfile_SAPODataConnectorProfileCredentialsProperty struct {
	// The SAPOData basic authentication credentials.
	BasicAuthCredentials interface{} `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// The SAPOData OAuth type authentication credentials.
	OAuthCredentials interface{} `field:"optional" json:"oAuthCredentials" yaml:"oAuthCredentials"`
}

