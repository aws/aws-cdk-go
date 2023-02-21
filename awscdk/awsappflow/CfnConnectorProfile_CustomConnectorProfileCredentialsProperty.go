package awsappflow


// The connector-specific profile credentials that are required when using the custom connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorProfileCredentialsProperty := &CustomConnectorProfileCredentialsProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//
//   	// the properties below are optional
//   	ApiKey: &ApiKeyCredentialsProperty{
//   		ApiKey: jsii.String("apiKey"),
//
//   		// the properties below are optional
//   		ApiSecretKey: jsii.String("apiSecretKey"),
//   	},
//   	Basic: &BasicAuthCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	Custom: &CustomAuthCredentialsProperty{
//   		CustomAuthenticationType: jsii.String("customAuthenticationType"),
//
//   		// the properties below are optional
//   		CredentialsMap: map[string]*string{
//   			"credentialsMapKey": jsii.String("credentialsMap"),
//   		},
//   	},
//   	Oauth2: &OAuth2CredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		OAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   }
//
type CfnConnectorProfile_CustomConnectorProfileCredentialsProperty struct {
	// The authentication type that the custom connector uses for authenticating while creating a connector profile.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The API keys required for the authentication of the user.
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The basic credentials that are required for the authentication of the user.
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// If the connector uses the custom authentication mechanism, this holds the required credentials.
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// The OAuth 2.0 credentials required for the authentication of the user.
	Oauth2 interface{} `field:"optional" json:"oauth2" yaml:"oauth2"`
}

