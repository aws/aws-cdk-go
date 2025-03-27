package awsdatazone


// OAuth2 Properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		JwtToken: jsii.String("jwtToken"),
//   		RefreshToken: jsii.String("refreshToken"),
//   		UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   	},
//   	OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   	TokenUrl: jsii.String("tokenUrl"),
//   	TokenUrlParametersMap: map[string]*string{
//   		"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html
//
type CfnConnection_OAuth2PropertiesProperty struct {
	// Authorization Code Properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-authorizationcodeproperties
	//
	AuthorizationCodeProperties interface{} `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// OAuth2 Client Application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-oauth2clientapplication
	//
	OAuth2ClientApplication interface{} `field:"optional" json:"oAuth2ClientApplication" yaml:"oAuth2ClientApplication"`
	// Glue OAuth2 Credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-oauth2credentials
	//
	OAuth2Credentials interface{} `field:"optional" json:"oAuth2Credentials" yaml:"oAuth2Credentials"`
	// OAuth2 Grant Type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-oauth2granttype
	//
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-tokenurl
	//
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// The token URL parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-tokenurlparametersmap
	//
	TokenUrlParametersMap interface{} `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

