package awsdatazone


// Glue Properties Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gluePropertiesInputProperty := &GluePropertiesInputProperty{
//   	GlueConnectionInput: &GlueConnectionInputProperty{
//   		AthenaProperties: map[string]*string{
//   			"athenaPropertiesKey": jsii.String("athenaProperties"),
//   		},
//   		AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//   			BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   				Password: jsii.String("password"),
//   				UserName: jsii.String("userName"),
//   			},
//   			CustomAuthenticationCredentials: map[string]*string{
//   				"customAuthenticationCredentialsKey": jsii.String("customAuthenticationCredentials"),
//   			},
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			OAuth2Properties: &OAuth2PropertiesProperty{
//   				AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   					AuthorizationCode: jsii.String("authorizationCode"),
//   					RedirectUri: jsii.String("redirectUri"),
//   				},
//   				OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   					AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   					UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   				},
//   				OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   					AccessToken: jsii.String("accessToken"),
//   					JwtToken: jsii.String("jwtToken"),
//   					RefreshToken: jsii.String("refreshToken"),
//   					UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   				},
//   				OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   				TokenUrl: jsii.String("tokenUrl"),
//   				TokenUrlParametersMap: map[string]*string{
//   					"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   				},
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		ConnectionProperties: map[string]*string{
//   			"connectionPropertiesKey": jsii.String("connectionProperties"),
//   		},
//   		ConnectionType: jsii.String("connectionType"),
//   		Description: jsii.String("description"),
//   		MatchCriteria: jsii.String("matchCriteria"),
//   		Name: jsii.String("name"),
//   		PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			SecurityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   			SubnetIdList: []*string{
//   				jsii.String("subnetIdList"),
//   			},
//   		},
//   		PythonProperties: map[string]*string{
//   			"pythonPropertiesKey": jsii.String("pythonProperties"),
//   		},
//   		SparkProperties: map[string]*string{
//   			"sparkPropertiesKey": jsii.String("sparkProperties"),
//   		},
//   		ValidateCredentials: jsii.Boolean(false),
//   		ValidateForComputeEnvironments: []*string{
//   			jsii.String("validateForComputeEnvironments"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-gluepropertiesinput.html
//
type CfnConnection_GluePropertiesInputProperty struct {
	// Glue Connection Input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-gluepropertiesinput.html#cfn-datazone-connection-gluepropertiesinput-glueconnectioninput
	//
	GlueConnectionInput interface{} `field:"optional" json:"glueConnectionInput" yaml:"glueConnectionInput"`
}

