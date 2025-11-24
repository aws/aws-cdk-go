package mixinsawsappflow


// Used by select connectors for which the OAuth workflow is supported, such as Salesforce, Google Analytics, Marketo, Zendesk, and Slack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectorOAuthRequestProperty := &ConnectorOAuthRequestProperty{
//   	AuthCode: jsii.String("authCode"),
//   	RedirectUri: jsii.String("redirectUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectoroauthrequest.html
//
type CfnConnectorProfilePropsMixin_ConnectorOAuthRequestProperty struct {
	// The code provided by the connector when it has been authenticated via the connected app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectoroauthrequest.html#cfn-appflow-connectorprofile-connectoroauthrequest-authcode
	//
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// The URL to which the authentication server redirects the browser after authorization has been granted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectoroauthrequest.html#cfn-appflow-connectorprofile-connectoroauthrequest-redirecturi
	//
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

