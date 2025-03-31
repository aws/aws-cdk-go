package awsdatazone


// The OAuth2 properties.
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
	// The authorization code properties of the OAuth2 properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-authorizationcodeproperties
	//
	AuthorizationCodeProperties interface{} `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// The OAuth2 client application of the OAuth2 properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-oauth2clientapplication
	//
	OAuth2ClientApplication interface{} `field:"optional" json:"oAuth2ClientApplication" yaml:"oAuth2ClientApplication"`
	// The OAuth2 credentials of the OAuth2 properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-oauth2credentials
	//
	OAuth2Credentials interface{} `field:"optional" json:"oAuth2Credentials" yaml:"oAuth2Credentials"`
	// The OAuth2 grant type of the OAuth2 properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-oauth2granttype
	//
	OAuth2GrantType *string `field:"optional" json:"oAuth2GrantType" yaml:"oAuth2GrantType"`
	// The OAuth2 token URL of the OAuth2 properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-tokenurl
	//
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// The OAuth2 token URL parameter map of the OAuth2 properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-oauth2properties.html#cfn-datazone-connection-oauth2properties-tokenurlparametersmap
	//
	TokenUrlParametersMap interface{} `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

