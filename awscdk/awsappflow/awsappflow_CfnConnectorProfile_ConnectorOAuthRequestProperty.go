package awsappflow


// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorOAuthRequestProperty := &connectorOAuthRequestProperty{
//   	authCode: jsii.String("authCode"),
//   	redirectUri: jsii.String("redirectUri"),
//   }
//
type CfnConnectorProfile_ConnectorOAuthRequestProperty struct {
	// The code provided by the connector when it has been authenticated via the connected app.
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// The URL to which the authentication server redirects the browser after authorization has been granted.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

