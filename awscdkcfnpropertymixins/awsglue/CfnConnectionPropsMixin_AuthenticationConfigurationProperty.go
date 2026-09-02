package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var customAuthenticationCredentials interface{}
//   var tokenUrlParametersMap interface{}
//
//   authenticationConfigurationProperty := &AuthenticationConfigurationProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	CustomAuthenticationCredentials: customAuthenticationCredentials,
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfiguration.html
//
type CfnConnectionPropsMixin_AuthenticationConfigurationProperty struct {
	// A structure containing the authentication configuration in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfiguration.html#cfn-glue-connection-authenticationconfiguration-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// A structure containing the authentication credentials in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfiguration.html#cfn-glue-connection-authenticationconfiguration-basicauthenticationcredentials
	//
	BasicAuthenticationCredentials interface{} `field:"optional" json:"basicAuthenticationCredentials" yaml:"basicAuthenticationCredentials"`
	// A structure containing the authentication credentials in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfiguration.html#cfn-glue-connection-authenticationconfiguration-customauthenticationcredentials
	//
	CustomAuthenticationCredentials interface{} `field:"optional" json:"customAuthenticationCredentials" yaml:"customAuthenticationCredentials"`
	// The Amazon Resource Name (ARN) of the KMS key used in the authentication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfiguration.html#cfn-glue-connection-authenticationconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A structure containing the authentication credentials in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfiguration.html#cfn-glue-connection-authenticationconfiguration-oauth2properties
	//
	OAuth2Properties interface{} `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// The secret manager ARN to store credentials in the CreateConnection request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-authenticationconfiguration.html#cfn-glue-connection-authenticationconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

