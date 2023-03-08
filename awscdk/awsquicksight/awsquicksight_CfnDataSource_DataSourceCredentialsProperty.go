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
//   dataSourceCredentialsProperty := &dataSourceCredentialsProperty{
//   	copySourceArn: jsii.String("copySourceArn"),
//   	credentialPair: &credentialPairProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//
//   		// the properties below are optional
//   		alternateDataSourceParameters: []interface{}{
//   			&dataSourceParametersProperty{
//   				amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   					domain: jsii.String("domain"),
//   				},
//   				amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   					domain: jsii.String("domain"),
//   				},
//   				athenaParameters: &athenaParametersProperty{
//   					workGroup: jsii.String("workGroup"),
//   				},
//   				auroraParameters: &auroraParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				databricksParameters: &databricksParametersProperty{
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   					sqlEndpointPath: jsii.String("sqlEndpointPath"),
//   				},
//   				mariaDbParameters: &mariaDbParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				mySqlParameters: &mySqlParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				oracleParameters: &oracleParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				postgreSqlParameters: &postgreSqlParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				prestoParameters: &prestoParametersProperty{
//   					catalog: jsii.String("catalog"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				rdsParameters: &rdsParametersProperty{
//   					database: jsii.String("database"),
//   					instanceId: jsii.String("instanceId"),
//   				},
//   				redshiftParameters: &redshiftParametersProperty{
//   					database: jsii.String("database"),
//
//   					// the properties below are optional
//   					clusterId: jsii.String("clusterId"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				s3Parameters: &s3ParametersProperty{
//   					manifestFileLocation: &manifestFileLocationProperty{
//   						bucket: jsii.String("bucket"),
//   						key: jsii.String("key"),
//   					},
//   				},
//   				snowflakeParameters: &snowflakeParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					warehouse: jsii.String("warehouse"),
//   				},
//   				sparkParameters: &sparkParametersProperty{
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				sqlServerParameters: &sqlServerParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				teradataParameters: &teradataParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	secretArn: jsii.String("secretArn"),
//   }
//
type CfnDataSource_DataSourceCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of a data source that has the credential pair that you want to use.
	//
	// When `CopySourceArn` is not null, the credential pair from the data source in the ARN is used as the credentials for the `DataSourceCredentials` structure.
	CopySourceArn *string `field:"optional" json:"copySourceArn" yaml:"copySourceArn"`
	// Credential pair.
	//
	// For more information, see `[CredentialPair](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CredentialPair.html)` .
	CredentialPair interface{} `field:"optional" json:"credentialPair" yaml:"credentialPair"`
	// `CfnDataSource.DataSourceCredentialsProperty.SecretArn`.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

