package awsquicksight


// The combination of user name and password that are used as credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialPairProperty := &credentialPairProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//
//   	// the properties below are optional
//   	alternateDataSourceParameters: []interface{}{
//   		&dataSourceParametersProperty{
//   			amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			athenaParameters: &athenaParametersProperty{
//   				workGroup: jsii.String("workGroup"),
//   			},
//   			auroraParameters: &auroraParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			databricksParameters: &databricksParametersProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   				sqlEndpointPath: jsii.String("sqlEndpointPath"),
//   			},
//   			mariaDbParameters: &mariaDbParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mySqlParameters: &mySqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			oracleParameters: &oracleParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			postgreSqlParameters: &postgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			prestoParameters: &prestoParametersProperty{
//   				catalog: jsii.String("catalog"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			rdsParameters: &rdsParametersProperty{
//   				database: jsii.String("database"),
//   				instanceId: jsii.String("instanceId"),
//   			},
//   			redshiftParameters: &redshiftParametersProperty{
//   				database: jsii.String("database"),
//
//   				// the properties below are optional
//   				clusterId: jsii.String("clusterId"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			s3Parameters: &s3ParametersProperty{
//   				manifestFileLocation: &manifestFileLocationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			snowflakeParameters: &snowflakeParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				warehouse: jsii.String("warehouse"),
//   			},
//   			sparkParameters: &sparkParametersProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			sqlServerParameters: &sqlServerParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			teradataParameters: &teradataParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSource_CredentialPairProperty struct {
	// Password.
	Password *string `field:"required" json:"password" yaml:"password"`
	// User name.
	Username *string `field:"required" json:"username" yaml:"username"`
	// A set of alternate data source parameters that you want to share for these credentials.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the new data source with the existing credentials. If the `AlternateDataSourceParameters` list is null, the `DataSourceParameters` originally used with these `Credentials` is automatically allowed.
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
}

