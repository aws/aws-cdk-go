package awsglue


// A structure containing the authentication credentials in the CreateConnection request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tokenUrlParametersMap interface{}
//
//   oAuth2PropertiesProperty := &OAuth2PropertiesProperty{
//   	AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   		AuthorizationCode: jsii.String("authorizationCode"),
//   		RedirectUri: jsii.String("redirectUri"),
//   	},
//   	OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   		AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   		UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   	},
//   	OAuth2Credentials: &OAuth2CredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		JwtToken: jsii.String("jwtToken"),
//   		RefreshToken: jsii.String("refreshToken"),
//   		UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   	},
//   	OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   	TokenUrl: jsii.String("tokenUrl"),
//   	TokenUrlParametersMap: tokenUrlParametersMap,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2properties.html
//
type CfnConnection_OAuth2PropertiesProperty struct {
	// A structure containing the authorization code used in the authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2properties.html#cfn-glue-connection-oauth2properties-authorizationcodeproperties
	//
	AuthorizationCodeProperties interface{} `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// The OAuth2 client app used for the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2properties.html#cfn-glue-connection-oauth2properties-oauth2clientapplication
	//
	OAuth2ClientApplication interface{} `field:"optional" json:"oAuth2ClientApplication" yaml:"oAuth2ClientApplication"`
	// A structure containing the OAuth2 credentials used in the authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2properties.html#cfn-glue-connection-oauth2properties-oauth2credentials
	//
	OAuth2Credentials interface{} `field:"optional" json:"oAuth2Credentials" yaml:"oAuth2Credentials"`
	// The grant type used in the authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2properties.html#cfn-glue-connection-oauth2properties-oauth2granttype
	//
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// The URL used in the authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2properties.html#cfn-glue-connection-oauth2properties-tokenurl
	//
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// A map of key-value pairs used in the authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2properties.html#cfn-glue-connection-oauth2properties-tokenurlparametersmap
	//
	TokenUrlParametersMap interface{} `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

