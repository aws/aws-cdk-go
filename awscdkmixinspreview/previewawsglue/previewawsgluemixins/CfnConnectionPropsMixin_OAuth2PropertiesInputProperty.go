package previewawsgluemixins


// A structure containing properties for OAuth2 in the CreateConnection request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tokenUrlParametersMap interface{}
//
//   oAuth2PropertiesInputProperty := &OAuth2PropertiesInputProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2propertiesinput.html
//
type CfnConnectionPropsMixin_OAuth2PropertiesInputProperty struct {
	// The set of properties required for the the OAuth2 `AUTHORIZATION_CODE` grant type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2propertiesinput.html#cfn-glue-connection-oauth2propertiesinput-authorizationcodeproperties
	//
	AuthorizationCodeProperties interface{} `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// The client application type in the CreateConnection request.
	//
	// For example, `AWS_MANAGED` or `USER_MANAGED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2propertiesinput.html#cfn-glue-connection-oauth2propertiesinput-oauth2clientapplication
	//
	OAuth2ClientApplication interface{} `field:"optional" json:"oAuth2ClientApplication" yaml:"oAuth2ClientApplication"`
	// The credentials used when the authentication type is OAuth2 authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2propertiesinput.html#cfn-glue-connection-oauth2propertiesinput-oauth2credentials
	//
	OAuth2Credentials interface{} `field:"optional" json:"oAuth2Credentials" yaml:"oAuth2Credentials"`
	// The OAuth2 grant type in the CreateConnection request.
	//
	// For example, `AUTHORIZATION_CODE` , `JWT_BEARER` , or `CLIENT_CREDENTIALS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2propertiesinput.html#cfn-glue-connection-oauth2propertiesinput-oauth2granttype
	//
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// The URL of the provider's authentication server, to exchange an authorization code for an access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2propertiesinput.html#cfn-glue-connection-oauth2propertiesinput-tokenurl
	//
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// A map of parameters that are added to the token `GET` request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2propertiesinput.html#cfn-glue-connection-oauth2propertiesinput-tokenurlparametersmap
	//
	TokenUrlParametersMap interface{} `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

