package awsquicksight


// Data source credentials.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceCredentialsProperty := &DataSourceCredentialsProperty{
//   	CopySourceArn: jsii.String("copySourceArn"),
//   	CredentialPair: &CredentialPairProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//
//   		// the properties below are optional
//   		AlternateDataSourceParameters: []interface{}{
//   			&DataSourceParametersProperty{
//   				AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   					Domain: jsii.String("domain"),
//   				},
//   				AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   					Domain: jsii.String("domain"),
//   				},
//   				AthenaParameters: &AthenaParametersProperty{
//   					IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   						EnableIdentityPropagation: jsii.Boolean(false),
//   					},
//   					RoleArn: jsii.String("roleArn"),
//   					WorkGroup: jsii.String("workGroup"),
//   				},
//   				AuroraParameters: &AuroraParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				AuroraPostgreSqlParameters: &AuroraPostgreSqlParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				DatabricksParameters: &DatabricksParametersProperty{
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   					SqlEndpointPath: jsii.String("sqlEndpointPath"),
//   				},
//   				MariaDbParameters: &MariaDbParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				MySqlParameters: &MySqlParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				OracleParameters: &OracleParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//
//   					// the properties below are optional
//   					UseServiceName: jsii.Boolean(false),
//   				},
//   				PostgreSqlParameters: &PostgreSqlParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				PrestoParameters: &PrestoParametersProperty{
//   					Catalog: jsii.String("catalog"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				RdsParameters: &RdsParametersProperty{
//   					Database: jsii.String("database"),
//   					InstanceId: jsii.String("instanceId"),
//   				},
//   				RedshiftParameters: &RedshiftParametersProperty{
//   					Database: jsii.String("database"),
//
//   					// the properties below are optional
//   					ClusterId: jsii.String("clusterId"),
//   					Host: jsii.String("host"),
//   					IamParameters: &RedshiftIAMParametersProperty{
//   						RoleArn: jsii.String("roleArn"),
//
//   						// the properties below are optional
//   						AutoCreateDatabaseUser: jsii.Boolean(false),
//   						DatabaseGroups: []*string{
//   							jsii.String("databaseGroups"),
//   						},
//   						DatabaseUser: jsii.String("databaseUser"),
//   					},
//   					IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   						EnableIdentityPropagation: jsii.Boolean(false),
//   					},
//   					Port: jsii.Number(123),
//   				},
//   				S3Parameters: &S3ParametersProperty{
//   					ManifestFileLocation: &ManifestFileLocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//
//   					// the properties below are optional
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				SnowflakeParameters: &SnowflakeParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Warehouse: jsii.String("warehouse"),
//
//   					// the properties below are optional
//   					AuthenticationType: jsii.String("authenticationType"),
//   					DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   					OAuthParameters: &OAuthParametersProperty{
//   						TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   						// the properties below are optional
//   						IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   						IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   							VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   						},
//   						OAuthScope: jsii.String("oAuthScope"),
//   					},
//   				},
//   				SparkParameters: &SparkParametersProperty{
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				SqlServerParameters: &SqlServerParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				StarburstParameters: &StarburstParametersProperty{
//   					Catalog: jsii.String("catalog"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//
//   					// the properties below are optional
//   					AuthenticationType: jsii.String("authenticationType"),
//   					DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   					OAuthParameters: &OAuthParametersProperty{
//   						TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   						// the properties below are optional
//   						IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   						IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   							VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   						},
//   						OAuthScope: jsii.String("oAuthScope"),
//   					},
//   					ProductType: jsii.String("productType"),
//   				},
//   				TeradataParameters: &TeradataParametersProperty{
//   					Database: jsii.String("database"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   				TrinoParameters: &TrinoParametersProperty{
//   					Catalog: jsii.String("catalog"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourcecredentials.html
//
type CfnDataSource_DataSourceCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of a data source that has the credential pair that you want to use.
	//
	// When `CopySourceArn` is not null, the credential pair from the data source in the ARN is used as the credentials for the `DataSourceCredentials` structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourcecredentials.html#cfn-quicksight-datasource-datasourcecredentials-copysourcearn
	//
	CopySourceArn *string `field:"optional" json:"copySourceArn" yaml:"copySourceArn"`
	// Credential pair.
	//
	// For more information, see `[CredentialPair](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CredentialPair.html)` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourcecredentials.html#cfn-quicksight-datasource-datasourcecredentials-credentialpair
	//
	CredentialPair interface{} `field:"optional" json:"credentialPair" yaml:"credentialPair"`
	// The Amazon Resource Name (ARN) of the secret associated with the data source in AWS Secrets Manager .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourcecredentials.html#cfn-quicksight-datasource-datasourcecredentials-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

