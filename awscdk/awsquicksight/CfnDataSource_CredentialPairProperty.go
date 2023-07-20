package awsquicksight


// The combination of user name and password that are used as credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialPairProperty := &CredentialPairProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	AlternateDataSourceParameters: []interface{}{
//   		&DataSourceParametersProperty{
//   			AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   				Domain: jsii.String("domain"),
//   			},
//   			AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   				Domain: jsii.String("domain"),
//   			},
//   			AthenaParameters: &AthenaParametersProperty{
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
//   				Database: jsii.String("database"),
//
//   				// the properties below are optional
//   				ClusterId: jsii.String("clusterId"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			S3Parameters: &S3ParametersProperty{
//   				ManifestFileLocation: &ManifestFileLocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   				},
//
//   				// the properties below are optional
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			SnowflakeParameters: &SnowflakeParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
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
//   			TeradataParameters: &TeradataParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html
//
type CfnDataSource_CredentialPairProperty struct {
	// Password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html#cfn-quicksight-datasource-credentialpair-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// User name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html#cfn-quicksight-datasource-credentialpair-username
	//
	Username *string `field:"required" json:"username" yaml:"username"`
	// A set of alternate data source parameters that you want to share for these credentials.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the new data source with the existing credentials. If the `AlternateDataSourceParameters` list is null, the `DataSourceParameters` originally used with these `Credentials` is automatically allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-credentialpair.html#cfn-quicksight-datasource-credentialpair-alternatedatasourceparameters
	//
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
}

