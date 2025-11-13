package awsdatazone


// Properties for defining a `CfnConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionProps := &CfnConnectionProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AwsLocation: &AwsLocationProperty{
//   		AccessRole: jsii.String("accessRole"),
//   		AwsAccountId: jsii.String("awsAccountId"),
//   		AwsRegion: jsii.String("awsRegion"),
//   		IamConnectionId: jsii.String("iamConnectionId"),
//   	},
//   	Description: jsii.String("description"),
//   	EnableTrustedIdentityPropagation: jsii.Boolean(false),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   	Props: &ConnectionPropertiesInputProperty{
//   		AmazonQProperties: &AmazonQPropertiesInputProperty{
//   			AuthMode: jsii.String("authMode"),
//   			IsEnabled: jsii.Boolean(false),
//   			ProfileArn: jsii.String("profileArn"),
//   		},
//   		AthenaProperties: &AthenaPropertiesInputProperty{
//   			WorkgroupName: jsii.String("workgroupName"),
//   		},
//   		GlueProperties: &GluePropertiesInputProperty{
//   			GlueConnectionInput: &GlueConnectionInputProperty{
//   				AthenaProperties: map[string]*string{
//   					"athenaPropertiesKey": jsii.String("athenaProperties"),
//   				},
//   				AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   					AuthenticationType: jsii.String("authenticationType"),
//   					BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   						Password: jsii.String("password"),
//   						UserName: jsii.String("userName"),
//   					},
//   					CustomAuthenticationCredentials: map[string]*string{
//   						"customAuthenticationCredentialsKey": jsii.String("customAuthenticationCredentials"),
//   					},
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					OAuth2Properties: &OAuth2PropertiesProperty{
//   						AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   							AuthorizationCode: jsii.String("authorizationCode"),
//   							RedirectUri: jsii.String("redirectUri"),
//   						},
//   						OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   							AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   							UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   						},
//   						OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   							AccessToken: jsii.String("accessToken"),
//   							JwtToken: jsii.String("jwtToken"),
//   							RefreshToken: jsii.String("refreshToken"),
//   							UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   						},
//   						OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   						TokenUrl: jsii.String("tokenUrl"),
//   						TokenUrlParametersMap: map[string]*string{
//   							"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   						},
//   					},
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				ConnectionProperties: map[string]*string{
//   					"connectionPropertiesKey": jsii.String("connectionProperties"),
//   				},
//   				ConnectionType: jsii.String("connectionType"),
//   				Description: jsii.String("description"),
//   				MatchCriteria: jsii.String("matchCriteria"),
//   				Name: jsii.String("name"),
//   				PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					SecurityGroupIdList: []*string{
//   						jsii.String("securityGroupIdList"),
//   					},
//   					SubnetId: jsii.String("subnetId"),
//   					SubnetIdList: []*string{
//   						jsii.String("subnetIdList"),
//   					},
//   				},
//   				PythonProperties: map[string]*string{
//   					"pythonPropertiesKey": jsii.String("pythonProperties"),
//   				},
//   				SparkProperties: map[string]*string{
//   					"sparkPropertiesKey": jsii.String("sparkProperties"),
//   				},
//   				ValidateCredentials: jsii.Boolean(false),
//   				ValidateForComputeEnvironments: []*string{
//   					jsii.String("validateForComputeEnvironments"),
//   				},
//   			},
//   		},
//   		HyperPodProperties: &HyperPodPropertiesInputProperty{
//   			ClusterName: jsii.String("clusterName"),
//   		},
//   		IamProperties: &IamPropertiesInputProperty{
//   			GlueLineageSyncEnabled: jsii.Boolean(false),
//   		},
//   		RedshiftProperties: &RedshiftPropertiesInputProperty{
//   			Credentials: &RedshiftCredentialsProperty{
//   				SecretArn: jsii.String("secretArn"),
//   				UsernamePassword: &UsernamePasswordProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   			},
//   			DatabaseName: jsii.String("databaseName"),
//   			Host: jsii.String("host"),
//   			LineageSync: &RedshiftLineageSyncConfigurationInputProperty{
//   				Enabled: jsii.Boolean(false),
//   				Schedule: &LineageSyncScheduleProperty{
//   					Schedule: jsii.String("schedule"),
//   				},
//   			},
//   			Port: jsii.Number(123),
//   			Storage: &RedshiftStoragePropertiesProperty{
//   				ClusterName: jsii.String("clusterName"),
//   				WorkgroupName: jsii.String("workgroupName"),
//   			},
//   		},
//   		S3Properties: &S3PropertiesInputProperty{
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			S3AccessGrantLocationId: jsii.String("s3AccessGrantLocationId"),
//   		},
//   		SparkEmrProperties: &SparkEmrPropertiesInputProperty{
//   			ComputeArn: jsii.String("computeArn"),
//   			InstanceProfileArn: jsii.String("instanceProfileArn"),
//   			JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   			LogUri: jsii.String("logUri"),
//   			PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   			RuntimeRole: jsii.String("runtimeRole"),
//   			TrustedCertificatesS3Uri: jsii.String("trustedCertificatesS3Uri"),
//   		},
//   		SparkGlueProperties: &SparkGluePropertiesInputProperty{
//   			AdditionalArgs: &SparkGlueArgsProperty{
//   				Connection: jsii.String("connection"),
//   			},
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   			GlueVersion: jsii.String("glueVersion"),
//   			IdleTimeout: jsii.Number(123),
//   			JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   			NumberOfWorkers: jsii.Number(123),
//   			PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   			WorkerType: jsii.String("workerType"),
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html
//
type CfnConnectionProps struct {
	// The ID of the domain where the connection is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The location where the connection is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-awslocation
	//
	AwsLocation interface{} `field:"optional" json:"awsLocation" yaml:"awsLocation"`
	// Connection description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the trusted identity propagation is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-enabletrustedidentitypropagation
	//
	EnableTrustedIdentityPropagation interface{} `field:"optional" json:"enableTrustedIdentityPropagation" yaml:"enableTrustedIdentityPropagation"`
	// The ID of the environment where the connection is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-environmentidentifier
	//
	EnvironmentIdentifier *string `field:"optional" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The identifier of the project in which the connection should be created.
	//
	// If.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-projectidentifier
	//
	ProjectIdentifier *string `field:"optional" json:"projectIdentifier" yaml:"projectIdentifier"`
	// Connection props.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-props
	//
	Props interface{} `field:"optional" json:"props" yaml:"props"`
	// The scope of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html#cfn-datazone-connection-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

