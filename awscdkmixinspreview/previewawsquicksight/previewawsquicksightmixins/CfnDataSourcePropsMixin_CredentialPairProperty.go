package previewawsquicksightmixins


// The combination of user name and password that are used as credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   credentialPairProperty := &CredentialPairProperty{
//   	AlternateDataSourceParameters: []interface{}{
//   		&DataSourceParametersProperty{
//   			AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   				Domain: jsii.String("domain"),
//   			},
//   			AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   				Domain: jsii.String("domain"),
//   			},
//   			AthenaParameters: &AthenaParametersProperty{
//   				IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   					EnableIdentityPropagation: jsii.Boolean(false),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   				WorkGroup: jsii.String("workGroup"),
//   			},
//   			AuroraParameters: &AuroraParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			AuroraPostgreSqlParameters: &AuroraPostgreSqlParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			DatabricksParameters: &DatabricksParametersProperty{
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   				SqlEndpointPath: jsii.String("sqlEndpointPath"),
//   			},
//   			MariaDbParameters: &MariaDbParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			MySqlParameters: &MySqlParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			OracleParameters: &OracleParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   				UseServiceName: jsii.Boolean(false),
//   			},
//   			PostgreSqlParameters: &PostgreSqlParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			PrestoParameters: &PrestoParametersProperty{
//   				Catalog: jsii.String("catalog"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			RdsParameters: &RdsParametersProperty{
//   				Database: jsii.String("database"),
//   				InstanceId: jsii.String("instanceId"),
//   			},
//   			RedshiftParameters: &RedshiftParametersProperty{
//   				ClusterId: jsii.String("clusterId"),
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				IamParameters: &RedshiftIAMParametersProperty{
//   					AutoCreateDatabaseUser: jsii.Boolean(false),
//   					DatabaseGroups: []*string{
//   						jsii.String("databaseGroups"),
//   					},
//   					DatabaseUser: jsii.String("databaseUser"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   					EnableIdentityPropagation: jsii.Boolean(false),
//   				},
//   				Port: jsii.Number(123),
//   			},
//   			S3Parameters: &S3ParametersProperty{
//   				ManifestFileLocation: &ManifestFileLocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			SnowflakeParameters: &SnowflakeParametersProperty{
//   				AuthenticationType: jsii.String("authenticationType"),
//   				Database: jsii.String("database"),
//   				DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   				Host: jsii.String("host"),
//   				OAuthParameters: &OAuthParametersProperty{
//   					IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   					IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   						VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   					},
//   					OAuthScope: jsii.String("oAuthScope"),
//   					TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   				},
//   				Warehouse: jsii.String("warehouse"),
//   			},
//   			SparkParameters: &SparkParametersProperty{
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			SqlServerParameters: &SqlServerParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			StarburstParameters: &StarburstParametersProperty{
//   				AuthenticationType: jsii.String("authenticationType"),
//   				Catalog: jsii.String("catalog"),
//   				DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   				Host: jsii.String("host"),
//   				OAuthParameters: &OAuthParametersProperty{
//   					IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   					IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   						VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   					},
//   					OAuthScope: jsii.String("oAuthScope"),
//   					TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   				},
//   				Port: jsii.Number(123),
//   				ProductType: jsii.String("productType"),
//   			},
//   			TeradataParameters: &TeradataParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			TrinoParameters: &TrinoParametersProperty{
//   				Catalog: jsii.String("catalog"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html
//
type CfnDataSourcePropsMixin_CredentialPairProperty struct {
	// A set of alternate data source parameters that you want to share for these credentials.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the new data source with the existing credentials. If the `AlternateDataSourceParameters` list is null, the `DataSourceParameters` originally used with these `Credentials` is automatically allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html#cfn-quicksight-datasource-credentialpair-alternatedatasourceparameters
	//
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
	// Password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html#cfn-quicksight-datasource-credentialpair-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// User name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html#cfn-quicksight-datasource-credentialpair-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

