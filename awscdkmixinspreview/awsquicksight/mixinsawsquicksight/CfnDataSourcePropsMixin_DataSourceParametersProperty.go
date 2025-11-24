package mixinsawsquicksight


// The parameters that Quick Sight uses to connect to your underlying data source.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceParametersProperty := &DataSourceParametersProperty{
//   	AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   		Domain: jsii.String("domain"),
//   	},
//   	AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   		Domain: jsii.String("domain"),
//   	},
//   	AthenaParameters: &AthenaParametersProperty{
//   		IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   			EnableIdentityPropagation: jsii.Boolean(false),
//   		},
//   		RoleArn: jsii.String("roleArn"),
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
//   		UseServiceName: jsii.Boolean(false),
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
//   		ClusterId: jsii.String("clusterId"),
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		IamParameters: &RedshiftIAMParametersProperty{
//   			AutoCreateDatabaseUser: jsii.Boolean(false),
//   			DatabaseGroups: []*string{
//   				jsii.String("databaseGroups"),
//   			},
//   			DatabaseUser: jsii.String("databaseUser"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   			EnableIdentityPropagation: jsii.Boolean(false),
//   		},
//   		Port: jsii.Number(123),
//   	},
//   	S3Parameters: &S3ParametersProperty{
//   		ManifestFileLocation: &ManifestFileLocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SnowflakeParameters: &SnowflakeParametersProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   		Database: jsii.String("database"),
//   		DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   		Host: jsii.String("host"),
//   		OAuthParameters: &OAuthParametersProperty{
//   			IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   			IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   				VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   			},
//   			OAuthScope: jsii.String("oAuthScope"),
//   			TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   		},
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
//   	StarburstParameters: &StarburstParametersProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   		Catalog: jsii.String("catalog"),
//   		DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   		Host: jsii.String("host"),
//   		OAuthParameters: &OAuthParametersProperty{
//   			IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   			IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   				VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   			},
//   			OAuthScope: jsii.String("oAuthScope"),
//   			TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   		},
//   		Port: jsii.Number(123),
//   		ProductType: jsii.String("productType"),
//   	},
//   	TeradataParameters: &TeradataParametersProperty{
//   		Database: jsii.String("database"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   	TrinoParameters: &TrinoParametersProperty{
//   		Catalog: jsii.String("catalog"),
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html
//
type CfnDataSourcePropsMixin_DataSourceParametersProperty struct {
	// The parameters for OpenSearch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-amazonelasticsearchparameters
	//
	AmazonElasticsearchParameters interface{} `field:"optional" json:"amazonElasticsearchParameters" yaml:"amazonElasticsearchParameters"`
	// The parameters for OpenSearch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-amazonopensearchparameters
	//
	AmazonOpenSearchParameters interface{} `field:"optional" json:"amazonOpenSearchParameters" yaml:"amazonOpenSearchParameters"`
	// The parameters for Amazon Athena.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-athenaparameters
	//
	AthenaParameters interface{} `field:"optional" json:"athenaParameters" yaml:"athenaParameters"`
	// The parameters for Amazon Aurora MySQL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-auroraparameters
	//
	AuroraParameters interface{} `field:"optional" json:"auroraParameters" yaml:"auroraParameters"`
	// The parameters for Amazon Aurora.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-aurorapostgresqlparameters
	//
	AuroraPostgreSqlParameters interface{} `field:"optional" json:"auroraPostgreSqlParameters" yaml:"auroraPostgreSqlParameters"`
	// The required parameters that are needed to connect to a Databricks data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-databricksparameters
	//
	DatabricksParameters interface{} `field:"optional" json:"databricksParameters" yaml:"databricksParameters"`
	// The parameters for MariaDB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-mariadbparameters
	//
	MariaDbParameters interface{} `field:"optional" json:"mariaDbParameters" yaml:"mariaDbParameters"`
	// The parameters for MySQL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-mysqlparameters
	//
	MySqlParameters interface{} `field:"optional" json:"mySqlParameters" yaml:"mySqlParameters"`
	// Oracle parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-oracleparameters
	//
	OracleParameters interface{} `field:"optional" json:"oracleParameters" yaml:"oracleParameters"`
	// The parameters for PostgreSQL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-postgresqlparameters
	//
	PostgreSqlParameters interface{} `field:"optional" json:"postgreSqlParameters" yaml:"postgreSqlParameters"`
	// The parameters for Presto.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-prestoparameters
	//
	PrestoParameters interface{} `field:"optional" json:"prestoParameters" yaml:"prestoParameters"`
	// The parameters for Amazon RDS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-rdsparameters
	//
	RdsParameters interface{} `field:"optional" json:"rdsParameters" yaml:"rdsParameters"`
	// The parameters for Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-redshiftparameters
	//
	RedshiftParameters interface{} `field:"optional" json:"redshiftParameters" yaml:"redshiftParameters"`
	// The parameters for S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-s3parameters
	//
	S3Parameters interface{} `field:"optional" json:"s3Parameters" yaml:"s3Parameters"`
	// The parameters for Snowflake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-snowflakeparameters
	//
	SnowflakeParameters interface{} `field:"optional" json:"snowflakeParameters" yaml:"snowflakeParameters"`
	// The parameters for Spark.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-sparkparameters
	//
	SparkParameters interface{} `field:"optional" json:"sparkParameters" yaml:"sparkParameters"`
	// The parameters for SQL Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-sqlserverparameters
	//
	SqlServerParameters interface{} `field:"optional" json:"sqlServerParameters" yaml:"sqlServerParameters"`
	// The parameters that are required to connect to a Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-starburstparameters
	//
	StarburstParameters interface{} `field:"optional" json:"starburstParameters" yaml:"starburstParameters"`
	// The parameters for Teradata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-teradataparameters
	//
	TeradataParameters interface{} `field:"optional" json:"teradataParameters" yaml:"teradataParameters"`
	// The parameters that are required to connect to a Trino data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-datasourceparameters.html#cfn-quicksight-datasource-datasourceparameters-trinoparameters
	//
	TrinoParameters interface{} `field:"optional" json:"trinoParameters" yaml:"trinoParameters"`
}

