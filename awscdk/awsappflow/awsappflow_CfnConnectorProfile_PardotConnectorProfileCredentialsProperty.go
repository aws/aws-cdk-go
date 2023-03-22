package awsappflow


// The connector-specific profile credentials required when using Salesforce Pardot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pardotConnectorProfileCredentialsProperty := &pardotConnectorProfileCredentialsProperty{
//   	accessToken: jsii.String("accessToken"),
//   	clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_PardotConnectorProfileCredentialsProperty struct {
	// The credentials used to access protected Salesforce Pardot resources.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The secret manager ARN, which contains the client ID and client secret of the connected app.
	ClientCredentialsArn *string `field:"optional" json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// `CfnConnectorProfile.PardotConnectorProfileCredentialsProperty.ConnectorOAuthRequest`.
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// The credentials used to acquire new access tokens.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

