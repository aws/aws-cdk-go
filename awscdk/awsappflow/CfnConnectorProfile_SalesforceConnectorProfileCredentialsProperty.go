package awsappflow


// The connector-specific profile credentials required when using Salesforce.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceConnectorProfileCredentialsProperty := &SalesforceConnectorProfileCredentialsProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	ClientCredentialsArn: jsii.String("clientCredentialsArn"),
//   	ConnectorOAuthRequest: &ConnectorOAuthRequestProperty{
//   		AuthCode: jsii.String("authCode"),
//   		RedirectUri: jsii.String("redirectUri"),
//   	},
//   	JwtToken: jsii.String("jwtToken"),
//   	OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   	RefreshToken: jsii.String("refreshToken"),
//   }
//
type CfnConnectorProfile_SalesforceConnectorProfileCredentialsProperty struct {
	// The credentials used to access protected Salesforce resources.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The secret manager ARN, which contains the client ID and client secret of the connected app.
	ClientCredentialsArn *string `field:"optional" json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
	// `CfnConnectorProfile.SalesforceConnectorProfileCredentialsProperty.JwtToken`.
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// `CfnConnectorProfile.SalesforceConnectorProfileCredentialsProperty.OAuth2GrantType`.
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// The credentials used to acquire new access tokens.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

