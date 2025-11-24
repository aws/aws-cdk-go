package mixinsawsappflow


// The connector-specific profile credentials required when using ServiceNow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceNowConnectorProfileCredentialsProperty := &ServiceNowConnectorProfileCredentialsProperty{
//   	OAuth2Credentials: &OAuth2CredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		OAuthRequest: &ConnectorOAuthRequestProperty{
//   			AuthCode: jsii.String("authCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-servicenowconnectorprofilecredentials.html
//
type CfnConnectorProfilePropsMixin_ServiceNowConnectorProfileCredentialsProperty struct {
	// The OAuth 2.0 credentials required to authenticate the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-servicenowconnectorprofilecredentials.html#cfn-appflow-connectorprofile-servicenowconnectorprofilecredentials-oauth2credentials
	//
	OAuth2Credentials interface{} `field:"optional" json:"oAuth2Credentials" yaml:"oAuth2Credentials"`
	// The password that corresponds to the user name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-servicenowconnectorprofilecredentials.html#cfn-appflow-connectorprofile-servicenowconnectorprofilecredentials-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-servicenowconnectorprofilecredentials.html#cfn-appflow-connectorprofile-servicenowconnectorprofilecredentials-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

