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
//   dataSourceParametersProperty := &DataSourceParametersProperty{
//   	AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   		Domain: jsii.String("domain"),
//   	},
//   	AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   		Domain: jsii.String("domain"),
//   	},
//   	AthenaParameters: &AthenaParametersProperty{
//   		WorkGroup: jsii.String("workGroup"),
//   	},
//   	AuroraParameters: &AuroraParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	AuroraPostgreSqlParameters: &AuroraPostgreSqlParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	DatabricksParameters: &DatabricksParametersProperty{
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   		SqlEndpointPath: jsii.String("sqlEndpointPath"),
//   	},
//   	MariaDbParameters: &MariaDbParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	MySqlParameters: &MySqlParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	OracleParameters: &OracleParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	PostgreSqlParameters: &PostgreSqlParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	PrestoParameters: &PrestoParametersProperty{
//   		Catalog: jsii.String("catalog"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	RdsParameters: &RdsParametersProperty{
//   		Database: jsii.String("database"),
//   		InstanceId: jsii.String("instanceId"),
//   	},
//   	RedshiftParameters: &RedshiftParametersProperty{
//   		Database: jsii.String("database"),
//
//   		// the properties below are optional
//   		ClusterId: jsii.String("clusterId"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	S3Parameters: &S3ParametersProperty{
//   		ManifestFileLocation: &ManifestFileLocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	SnowflakeParameters: &SnowflakeParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Warehouse: jsii.String("warehouse"),
//   	},
//   	SparkParameters: &SparkParametersProperty{
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	SqlServerParameters: &SqlServerParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	TeradataParameters: &TeradataParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
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
	// The required parameters that are needed to connect to a Databricks data source.
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

