package mixinsawsglue


// Properties for CfnConnectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var athenaProperties interface{}
//   var connectionProperties interface{}
//   var customAuthenticationCredentials interface{}
//   var pythonProperties interface{}
//   var sparkProperties interface{}
//   var tokenUrlParametersMap interface{}
//
//   cfnConnectionMixinProps := &CfnConnectionMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	ConnectionInput: &ConnectionInputProperty{
//   		AthenaProperties: athenaProperties,
//   		AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//   			BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   			CustomAuthenticationCredentials: customAuthenticationCredentials,
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			OAuth2Properties: &OAuth2PropertiesInputProperty{
//   				AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   					AuthorizationCode: jsii.String("authorizationCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   					AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   					UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   				},
//   				OAuth2Credentials: &OAuth2CredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					JwtToken: jsii.String("jwtToken"),
//   					RefreshToken: jsii.String("refreshToken"),
//   					UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   				},
//   				OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   				TokenUrl: jsii.String("tokenUrl"),
//   				TokenUrlParametersMap: tokenUrlParametersMap,
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		ConnectionProperties: connectionProperties,
//   		ConnectionType: jsii.String("connectionType"),
//   		Description: jsii.String("description"),
//   		MatchCriteria: []*string{
//   			jsii.String("matchCriteria"),
//   		},
//   		Name: jsii.String("name"),
//   		PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			SecurityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   		PythonProperties: pythonProperties,
//   		SparkProperties: sparkProperties,
//   		ValidateCredentials: jsii.Boolean(false),
//   		ValidateForComputeEnvironments: []*string{
//   			jsii.String("validateForComputeEnvironments"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html
//
type CfnConnectionMixinProps struct {
	// The ID of the data catalog to create the catalog object in.
	//
	// Currently, this should be the AWS account ID.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html#cfn-glue-connection-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The connection that you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html#cfn-glue-connection-connectioninput
	//
	ConnectionInput interface{} `field:"optional" json:"connectionInput" yaml:"connectionInput"`
}

