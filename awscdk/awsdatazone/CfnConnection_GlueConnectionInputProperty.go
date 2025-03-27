package awsdatazone


// Glue Connection Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueConnectionInputProperty := &GlueConnectionInputProperty{
//   	AthenaProperties: map[string]*string{
//   		"athenaPropertiesKey": jsii.String("athenaProperties"),
//   	},
//   	AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   		BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   			Password: jsii.String("password"),
//   			UserName: jsii.String("userName"),
//   		},
//   		CustomAuthenticationCredentials: map[string]*string{
//   			"customAuthenticationCredentialsKey": jsii.String("customAuthenticationCredentials"),
//   		},
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		OAuth2Properties: &OAuth2PropertiesProperty{
//   			AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   				AuthorizationCode: jsii.String("authorizationCode"),
//   				RedirectUri: jsii.String("redirectUri"),
//   			},
//   			OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   				AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   				UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   			},
//   			OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   				AccessToken: jsii.String("accessToken"),
//   				JwtToken: jsii.String("jwtToken"),
//   				RefreshToken: jsii.String("refreshToken"),
//   				UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   			},
//   			OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   			TokenUrl: jsii.String("tokenUrl"),
//   			TokenUrlParametersMap: map[string]*string{
//   				"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   			},
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	ConnectionProperties: map[string]*string{
//   		"connectionPropertiesKey": jsii.String("connectionProperties"),
//   	},
//   	ConnectionType: jsii.String("connectionType"),
//   	Description: jsii.String("description"),
//   	MatchCriteria: jsii.String("matchCriteria"),
//   	Name: jsii.String("name"),
//   	PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		SecurityGroupIdList: []*string{
//   			jsii.String("securityGroupIdList"),
//   		},
//   		SubnetId: jsii.String("subnetId"),
//   		SubnetIdList: []*string{
//   			jsii.String("subnetIdList"),
//   		},
//   	},
//   	PythonProperties: map[string]*string{
//   		"pythonPropertiesKey": jsii.String("pythonProperties"),
//   	},
//   	SparkProperties: map[string]*string{
//   		"sparkPropertiesKey": jsii.String("sparkProperties"),
//   	},
//   	ValidateCredentials: jsii.Boolean(false),
//   	ValidateForComputeEnvironments: []*string{
//   		jsii.String("validateForComputeEnvironments"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html
//
type CfnConnection_GlueConnectionInputProperty struct {
	// Property Map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-athenaproperties
	//
	AthenaProperties interface{} `field:"optional" json:"athenaProperties" yaml:"athenaProperties"`
	// Authentication Configuration Input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Connection Properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-connectionproperties
	//
	ConnectionProperties interface{} `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// Glue Connection Type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-connectiontype
	//
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-matchcriteria
	//
	MatchCriteria *string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Physical Connection Requirements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-physicalconnectionrequirements
	//
	PhysicalConnectionRequirements interface{} `field:"optional" json:"physicalConnectionRequirements" yaml:"physicalConnectionRequirements"`
	// Property Map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-pythonproperties
	//
	PythonProperties interface{} `field:"optional" json:"pythonProperties" yaml:"pythonProperties"`
	// Property Map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-sparkproperties
	//
	SparkProperties interface{} `field:"optional" json:"sparkProperties" yaml:"sparkProperties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-validatecredentials
	//
	ValidateCredentials interface{} `field:"optional" json:"validateCredentials" yaml:"validateCredentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-glueconnectioninput.html#cfn-datazone-connection-glueconnectioninput-validateforcomputeenvironments
	//
	ValidateForComputeEnvironments *[]*string `field:"optional" json:"validateForComputeEnvironments" yaml:"validateForComputeEnvironments"`
}

