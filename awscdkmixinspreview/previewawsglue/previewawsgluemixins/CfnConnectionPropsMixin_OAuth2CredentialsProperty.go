package previewawsgluemixins


// The credentials used when the authentication type is OAuth2 authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   oAuth2CredentialsProperty := &OAuth2CredentialsProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	JwtToken: jsii.String("jwtToken"),
//   	RefreshToken: jsii.String("refreshToken"),
//   	UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2credentials.html
//
type CfnConnectionPropsMixin_OAuth2CredentialsProperty struct {
	// The access token used when the authentication type is OAuth2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2credentials.html#cfn-glue-connection-oauth2credentials-accesstoken
	//
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The JSON Web Token (JWT) used when the authentication type is OAuth2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2credentials.html#cfn-glue-connection-oauth2credentials-jwttoken
	//
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// The refresh token used when the authentication type is OAuth2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2credentials.html#cfn-glue-connection-oauth2credentials-refreshtoken
	//
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// The client application client secret if the client application is user managed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-oauth2credentials.html#cfn-glue-connection-oauth2credentials-usermanagedclientapplicationclientsecret
	//
	UserManagedClientApplicationClientSecret *string `field:"optional" json:"userManagedClientApplicationClientSecret" yaml:"userManagedClientApplicationClientSecret"`
}

