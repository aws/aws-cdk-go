package awsappflow


// The connector-specific profile credentials required when using Slack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackConnectorProfileCredentialsProperty := &slackConnectorProfileCredentialsProperty{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	accessToken: jsii.String("accessToken"),
//   	connectorOAuthRequest: &connectorOAuthRequestProperty{
//   		authCode: jsii.String("authCode"),
//   		redirectUri: jsii.String("redirectUri"),
//   	},
//   }
//
type CfnConnectorProfile_SlackConnectorProfileCredentialsProperty struct {
	// The identifier for the client.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret used by the OAuth client to authenticate to the authorization server.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The credentials used to access protected Slack resources.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
	ConnectorOAuthRequest interface{} `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
}

