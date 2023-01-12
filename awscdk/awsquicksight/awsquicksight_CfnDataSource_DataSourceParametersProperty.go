package awsquicksight


// The parameters that Amazon QuickSight uses to connect to your underlying data source.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceParametersProperty := &dataSourceParametersProperty{
//   	amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   		domain: jsii.String("domain"),
//   	},
//   	amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   		domain: jsii.String("domain"),
//   	},
//   	athenaParameters: &athenaParametersProperty{
//   		workGroup: jsii.String("workGroup"),
//   	},
//   	auroraParameters: &auroraParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	databricksParameters: &databricksParametersProperty{
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   		sqlEndpointPath: jsii.String("sqlEndpointPath"),
//   	},
//   	mariaDbParameters: &mariaDbParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	mySqlParameters: &mySqlParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	oracleParameters: &oracleParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	postgreSqlParameters: &postgreSqlParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	prestoParameters: &prestoParametersProperty{
//   		catalog: jsii.String("catalog"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	rdsParameters: &rdsParametersProperty{
//   		database: jsii.String("database"),
//   		instanceId: jsii.String("instanceId"),
//   	},
//   	redshiftParameters: &redshiftParametersProperty{
//   		database: jsii.String("database"),
//
//   		// the properties below are optional
//   		clusterId: jsii.String("clusterId"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	s3Parameters: &s3ParametersProperty{
//   		manifestFileLocation: &manifestFileLocationProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   	snowflakeParameters: &snowflakeParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		warehouse: jsii.String("warehouse"),
//   	},
//   	sparkParameters: &sparkParametersProperty{
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	sqlServerParameters: &sqlServerParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	teradataParameters: &teradataParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   }
//
type CfnDataSource_DataSourceParametersProperty struct {
	// The parameters for OpenSearch.
	AmazonElasticsearchParameters interface{} `field:"optional" json:"amazonElasticsearchParameters" yaml:"amazonElasticsearchParameters"`
	// The parameters for OpenSearch.
	AmazonOpenSearchParameters interface{} `field:"optional" json:"amazonOpenSearchParameters" yaml:"amazonOpenSearchParameters"`
	// The parameters for Amazon Athena.
	AthenaParameters interface{} `field:"optional" json:"athenaParameters" yaml:"athenaParameters"`
	// The parameters for Amazon Aurora MySQL.
	AuroraParameters interface{} `field:"optional" json:"auroraParameters" yaml:"auroraParameters"`
	// The parameters for Amazon Aurora.
	AuroraPostgreSqlParameters interface{} `field:"optional" json:"auroraPostgreSqlParameters" yaml:"auroraPostgreSqlParameters"`
	// `CfnDataSource.DataSourceParametersProperty.DatabricksParameters`.
	DatabricksParameters interface{} `field:"optional" json:"databricksParameters" yaml:"databricksParameters"`
	// The parameters for MariaDB.
	MariaDbParameters interface{} `field:"optional" json:"mariaDbParameters" yaml:"mariaDbParameters"`
	// The parameters for MySQL.
	MySqlParameters interface{} `field:"optional" json:"mySqlParameters" yaml:"mySqlParameters"`
	// Oracle parameters.
	OracleParameters interface{} `field:"optional" json:"oracleParameters" yaml:"oracleParameters"`
	// The parameters for PostgreSQL.
	PostgreSqlParameters interface{} `field:"optional" json:"postgreSqlParameters" yaml:"postgreSqlParameters"`
	// The parameters for Presto.
	PrestoParameters interface{} `field:"optional" json:"prestoParameters" yaml:"prestoParameters"`
	// The parameters for Amazon RDS.
	RdsParameters interface{} `field:"optional" json:"rdsParameters" yaml:"rdsParameters"`
	// The parameters for Amazon Redshift.
	RedshiftParameters interface{} `field:"optional" json:"redshiftParameters" yaml:"redshiftParameters"`
	// The parameters for S3.
	S3Parameters interface{} `field:"optional" json:"s3Parameters" yaml:"s3Parameters"`
	// The parameters for Snowflake.
	SnowflakeParameters interface{} `field:"optional" json:"snowflakeParameters" yaml:"snowflakeParameters"`
	// The parameters for Spark.
	SparkParameters interface{} `field:"optional" json:"sparkParameters" yaml:"sparkParameters"`
	// The parameters for SQL Server.
	SqlServerParameters interface{} `field:"optional" json:"sqlServerParameters" yaml:"sqlServerParameters"`
	// The parameters for Teradata.
	TeradataParameters interface{} `field:"optional" json:"teradataParameters" yaml:"teradataParameters"`
}

