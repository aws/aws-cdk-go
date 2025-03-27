package awsdatazone


// Glue OAuth2 Credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueOAuth2CredentialsProperty := &GlueOAuth2CredentialsProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	JwtToken: jsii.String("jwtToken"),
//   	RefreshToken: jsii.String("refreshToken"),
//   	UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueoauth2credentials.html
//
type CfnConnection_GlueOAuth2CredentialsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueoauth2credentials.html#cfn-datazone-connection-glueoauth2credentials-accesstoken
	//
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueoauth2credentials.html#cfn-datazone-connection-glueoauth2credentials-jwttoken
	//
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueoauth2credentials.html#cfn-datazone-connection-glueoauth2credentials-refreshtoken
	//
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueoauth2credentials.html#cfn-datazone-connection-glueoauth2credentials-usermanagedclientapplicationclientsecret
	//
	UserManagedClientApplicationClientSecret *string `field:"optional" json:"userManagedClientApplicationClientSecret" yaml:"userManagedClientApplicationClientSecret"`
}

