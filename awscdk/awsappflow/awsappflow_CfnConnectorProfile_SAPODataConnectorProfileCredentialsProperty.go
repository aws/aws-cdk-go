package awsappflow


// The connector-specific profile credentials required when using SAPOData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAPODataConnectorProfileCredentialsProperty := &sAPODataConnectorProfileCredentialsProperty{
//   	basicAuthCredentials: &basicAuthCredentialsProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	oAuthCredentials: &oAuthCredentialsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		clientId: jsii.String("clientId"),
//   		clientSecret: jsii.String("clientSecret"),
//   		connectorOAuthRequest: &connectorOAuthRequestProperty{
//   			authCode: jsii.String("authCode"),
//   			redirectUri: jsii.String("redirectUri"),
//   		},
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   }
//
type CfnConnectorProfile_SAPODataConnectorProfileCredentialsProperty struct {
	// The SAPOData basic authentication credentials.
	BasicAuthCredentials interface{} `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// The SAPOData OAuth type authentication credentials.
	OAuthCredentials interface{} `field:"optional" json:"oAuthCredentials" yaml:"oAuthCredentials"`
}

