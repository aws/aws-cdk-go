package awsdatazone


// The properties of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionPropertiesInputProperty := &ConnectionPropertiesInputProperty{
//   	AmazonQProperties: &AmazonQPropertiesInputProperty{
//   		AuthMode: jsii.String("authMode"),
//   		IsEnabled: jsii.Boolean(false),
//   		ProfileArn: jsii.String("profileArn"),
//   	},
//   	AthenaProperties: &AthenaPropertiesInputProperty{
//   		WorkgroupName: jsii.String("workgroupName"),
//   	},
//   	GlueProperties: &GluePropertiesInputProperty{
//   		GlueConnectionInput: &GlueConnectionInputProperty{
//   			AthenaProperties: map[string]*string{
//   				"athenaPropertiesKey": jsii.String("athenaProperties"),
//   			},
//   			AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   				AuthenticationType: jsii.String("authenticationType"),
//   				BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   					Password: jsii.String("password"),
//   					UserName: jsii.String("userName"),
//   				},
//   				CustomAuthenticationCredentials: map[string]*string{
//   					"customAuthenticationCredentialsKey": jsii.String("customAuthenticationCredentials"),
//   				},
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   				OAuth2Properties: &OAuth2PropertiesProperty{
//   					AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   						AuthorizationCode: jsii.String("authorizationCode"),
//   						RedirectUri: jsii.String("redirectUri"),
//   					},
//   					OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   						AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   						UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   					},
//   					OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   						AccessToken: jsii.String("accessToken"),
//   						JwtToken: jsii.String("jwtToken"),
//   						RefreshToken: jsii.String("refreshToken"),
//   						UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   					},
//   					OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   					TokenUrl: jsii.String("tokenUrl"),
//   					TokenUrlParametersMap: map[string]*string{
//   						"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   					},
//   				},
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   			ConnectionProperties: map[string]*string{
//   				"connectionPropertiesKey": jsii.String("connectionProperties"),
//   			},
//   			ConnectionType: jsii.String("connectionType"),
//   			Description: jsii.String("description"),
//   			MatchCriteria: jsii.String("matchCriteria"),
//   			Name: jsii.String("name"),
//   			PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   				AvailabilityZone: jsii.String("availabilityZone"),
//   				SecurityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   				SubnetIdList: []*string{
//   					jsii.String("subnetIdList"),
//   				},
//   			},
//   			PythonProperties: map[string]*string{
//   				"pythonPropertiesKey": jsii.String("pythonProperties"),
//   			},
//   			SparkProperties: map[string]*string{
//   				"sparkPropertiesKey": jsii.String("sparkProperties"),
//   			},
//   			ValidateCredentials: jsii.Boolean(false),
//   			ValidateForComputeEnvironments: []*string{
//   				jsii.String("validateForComputeEnvironments"),
//   			},
//   		},
//   	},
//   	HyperPodProperties: &HyperPodPropertiesInputProperty{
//   		ClusterName: jsii.String("clusterName"),
//   	},
//   	IamProperties: &IamPropertiesInputProperty{
//   		GlueLineageSyncEnabled: jsii.Boolean(false),
//   	},
//   	RedshiftProperties: &RedshiftPropertiesInputProperty{
//   		Credentials: &RedshiftCredentialsProperty{
//   			SecretArn: jsii.String("secretArn"),
//   			UsernamePassword: &UsernamePasswordProperty{
//   				Password: jsii.String("password"),
//   				Username: jsii.String("username"),
//   			},
//   		},
//   		DatabaseName: jsii.String("databaseName"),
//   		Host: jsii.String("host"),
//   		LineageSync: &RedshiftLineageSyncConfigurationInputProperty{
//   			Enabled: jsii.Boolean(false),
//   			Schedule: &LineageSyncScheduleProperty{
//   				Schedule: jsii.String("schedule"),
//   			},
//   		},
//   		Port: jsii.Number(123),
//   		Storage: &RedshiftStoragePropertiesProperty{
//   			ClusterName: jsii.String("clusterName"),
//   			WorkgroupName: jsii.String("workgroupName"),
//   		},
//   	},
//   	S3Properties: &S3PropertiesInputProperty{
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		S3AccessGrantLocationId: jsii.String("s3AccessGrantLocationId"),
//   	},
//   	SparkEmrProperties: &SparkEmrPropertiesInputProperty{
//   		ComputeArn: jsii.String("computeArn"),
//   		InstanceProfileArn: jsii.String("instanceProfileArn"),
//   		JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   		LogUri: jsii.String("logUri"),
//   		ManagedEndpointArn: jsii.String("managedEndpointArn"),
//   		PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   		RuntimeRole: jsii.String("runtimeRole"),
//   		TrustedCertificatesS3Uri: jsii.String("trustedCertificatesS3Uri"),
//   	},
//   	SparkGlueProperties: &SparkGluePropertiesInputProperty{
//   		AdditionalArgs: &SparkGlueArgsProperty{
//   			Connection: jsii.String("connection"),
//   		},
//   		GlueConnectionName: jsii.String("glueConnectionName"),
//   		GlueVersion: jsii.String("glueVersion"),
//   		IdleTimeout: jsii.Number(123),
//   		JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   		NumberOfWorkers: jsii.Number(123),
//   		PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   		WorkerType: jsii.String("workerType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html
//
type CfnConnection_ConnectionPropertiesInputProperty struct {
	// Amazon Q properties of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-amazonqproperties
	//
	AmazonQProperties interface{} `field:"optional" json:"amazonQProperties" yaml:"amazonQProperties"`
	// The Amazon Athena properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-athenaproperties
	//
	AthenaProperties interface{} `field:"optional" json:"athenaProperties" yaml:"athenaProperties"`
	// The AWS Glue properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-glueproperties
	//
	GlueProperties interface{} `field:"optional" json:"glueProperties" yaml:"glueProperties"`
	// The hyper pod properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-hyperpodproperties
	//
	HyperPodProperties interface{} `field:"optional" json:"hyperPodProperties" yaml:"hyperPodProperties"`
	// The IAM properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-iamproperties
	//
	IamProperties interface{} `field:"optional" json:"iamProperties" yaml:"iamProperties"`
	// The Amazon Redshift properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-redshiftproperties
	//
	RedshiftProperties interface{} `field:"optional" json:"redshiftProperties" yaml:"redshiftProperties"`
	// S3 Properties Input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-s3properties
	//
	S3Properties interface{} `field:"optional" json:"s3Properties" yaml:"s3Properties"`
	// The Spark EMR properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-sparkemrproperties
	//
	SparkEmrProperties interface{} `field:"optional" json:"sparkEmrProperties" yaml:"sparkEmrProperties"`
	// The Spark AWS Glue properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-connectionpropertiesinput.html#cfn-datazone-connection-connectionpropertiesinput-sparkglueproperties
	//
	SparkGlueProperties interface{} `field:"optional" json:"sparkGlueProperties" yaml:"sparkGlueProperties"`
}

