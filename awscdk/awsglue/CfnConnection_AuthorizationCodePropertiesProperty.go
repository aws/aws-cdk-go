package awsglue


// The set of properties required for the the OAuth2 `AUTHORIZATION_CODE` grant type workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationCodePropertiesProperty := &AuthorizationCodePropertiesProperty{
//   	AuthorizationCode: jsii.String("authorizationCode"),
//   	RedirectUri: jsii.String("redirectUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authorizationcodeproperties.html
//
type CfnConnection_AuthorizationCodePropertiesProperty struct {
	// An authorization code to be used in the third leg of the `AUTHORIZATION_CODE` grant workflow.
	//
	// This is a single-use code which becomes invalid once exchanged for an access token, thus it is acceptable to have this value as a request parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authorizationcodeproperties.html#cfn-glue-connection-authorizationcodeproperties-authorizationcode
	//
	AuthorizationCode *string `field:"optional" json:"authorizationCode" yaml:"authorizationCode"`
	// The redirect URI where the user gets redirected to by authorization server when issuing an authorization code.
	//
	// The URI is subsequently used when the authorization code is exchanged for an access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authorizationcodeproperties.html#cfn-glue-connection-authorizationcodeproperties-redirecturi
	//
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

