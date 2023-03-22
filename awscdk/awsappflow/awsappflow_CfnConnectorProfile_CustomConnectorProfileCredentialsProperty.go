package awsappflow


// The connector-specific profile credentials that are required when using the custom connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorProfileCredentialsProperty := &customConnectorProfileCredentialsProperty{
//   	authenticationType: jsii.String("authenticationType"),
//
//   	// the properties below are optional
//   	apiKey: &apiKeyCredentialsProperty{
//   		apiKey: jsii.String("apiKey"),
//
//   		// the properties below are optional
//   		apiSecretKey: jsii.String("apiSecretKey"),
//   	},
//   	basic: &basicAuthCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	custom: &customAuthCredentialsProperty{
//   		customAuthenticationType: jsii.String("customAuthenticationType"),
//
//   		// the properties below are optional
//   		credentialsMap: map[string]*string{
//   			"credentialsMapKey": jsii.String("credentialsMap"),
//   		},
//   	},
//   	oauth2: &oAuth2CredentialsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//   		oAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   		refreshToken: jsii.String("refreshToken"),
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

