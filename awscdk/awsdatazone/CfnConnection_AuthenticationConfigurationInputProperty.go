package awsdatazone


// The authentication configuration of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationConfigurationInputProperty := &AuthenticationConfigurationInputProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   		Password: jsii.String("password"),
//   		UserName: jsii.String("userName"),
//   	},
//   	CustomAuthenticationCredentials: map[string]*string{
//   		"customAuthenticationCredentialsKey": jsii.String("customAuthenticationCredentials"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	OAuth2Properties: &OAuth2PropertiesProperty{
//   		AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   			AuthorizationCode: jsii.String("authorizationCode"),
//   			RedirectUri: jsii.String("redirectUri"),
//   		},
//   		OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   			AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   			UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   		},
//   		OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			JwtToken: jsii.String("jwtToken"),
//   			RefreshToken: jsii.String("refreshToken"),
//   			UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   		},
//   		OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   		TokenUrl: jsii.String("tokenUrl"),
//   		TokenUrlParametersMap: map[string]*string{
//   			"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   		},
//   	},
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authenticationconfigurationinput.html
//
type CfnConnection_AuthenticationConfigurationInputProperty struct {
	// The authentication type of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authenticationconfigurationinput.html#cfn-datazone-connection-authenticationconfigurationinput-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The basic authentication credentials of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authenticationconfigurationinput.html#cfn-datazone-connection-authenticationconfigurationinput-basicauthenticationcredentials
	//
	BasicAuthenticationCredentials interface{} `field:"optional" json:"basicAuthenticationCredentials" yaml:"basicAuthenticationCredentials"`
	// The custom authentication credentials of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authenticationconfigurationinput.html#cfn-datazone-connection-authenticationconfigurationinput-customauthenticationcredentials
	//
	CustomAuthenticationCredentials interface{} `field:"optional" json:"customAuthenticationCredentials" yaml:"customAuthenticationCredentials"`
	// The KMS key ARN of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authenticationconfigurationinput.html#cfn-datazone-connection-authenticationconfigurationinput-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The oAuth2 properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authenticationconfigurationinput.html#cfn-datazone-connection-authenticationconfigurationinput-oauth2properties
	//
	OAuth2Properties interface{} `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// The secret ARN of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-authenticationconfigurationinput.html#cfn-datazone-connection-authenticationconfigurationinput-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

