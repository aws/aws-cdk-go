package awsglue


// A structure containing the authentication configuration in the CreateConnection request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customAuthenticationCredentials interface{}
//   var tokenUrlParametersMap interface{}
//
//   authenticationConfigurationInputProperty := &AuthenticationConfigurationInputProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//
//   	// the properties below are optional
//   	BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	CustomAuthenticationCredentials: customAuthenticationCredentials,
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	OAuth2Properties: &OAuth2PropertiesInputProperty{
//   		AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   			AuthorizationCode: jsii.String("authorizationCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   			AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   			UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   		},
//   		OAuth2Credentials: &OAuth2CredentialsProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			JwtToken: jsii.String("jwtToken"),
//   			RefreshToken: jsii.String("refreshToken"),
//   			UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   		},
//   		OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   		TokenUrl: jsii.String("tokenUrl"),
//   		TokenUrlParametersMap: tokenUrlParametersMap,
//   	},
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfigurationinput.html
//
type CfnConnection_AuthenticationConfigurationInputProperty struct {
	// A structure containing the authentication configuration in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfigurationinput.html#cfn-glue-connection-authenticationconfigurationinput-authenticationtype
	//
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The credentials used when the authentication type is basic authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfigurationinput.html#cfn-glue-connection-authenticationconfigurationinput-basicauthenticationcredentials
	//
	BasicAuthenticationCredentials interface{} `field:"optional" json:"basicAuthenticationCredentials" yaml:"basicAuthenticationCredentials"`
	// The credentials used when the authentication type is custom authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfigurationinput.html#cfn-glue-connection-authenticationconfigurationinput-customauthenticationcredentials
	//
	CustomAuthenticationCredentials interface{} `field:"optional" json:"customAuthenticationCredentials" yaml:"customAuthenticationCredentials"`
	// The ARN of the KMS key used to encrypt the connection.
	//
	// Only taken an as input in the request and stored in the Secret Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfigurationinput.html#cfn-glue-connection-authenticationconfigurationinput-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The properties for OAuth2 authentication in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfigurationinput.html#cfn-glue-connection-authenticationconfigurationinput-oauth2properties
	//
	OAuth2Properties interface{} `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// The secret manager ARN to store credentials in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfigurationinput.html#cfn-glue-connection-authenticationconfigurationinput-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

