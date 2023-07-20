package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &CfnDataSourceProps{
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
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Credentials: &DataSourceCredentialsProperty{
//   		CopySourceArn: jsii.String("copySourceArn"),
//   		CredentialPair: &CredentialPairProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//
//   			// the properties below are optional
//   			AlternateDataSourceParameters: []interface{}{
//   				&DataSourceParametersProperty{
//   					AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   						Domain: jsii.String("domain"),
//   					},
//   					AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   						Domain: jsii.String("domain"),
//   					},
//   					AthenaParameters: &AthenaParametersProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						WorkGroup: jsii.String("workGroup"),
//   					},
//   					AuroraParameters: &AuroraParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					AuroraPostgreSqlParameters: &AuroraPostgreSqlParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					DatabricksParameters: &DatabricksParametersProperty{
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   						SqlEndpointPath: jsii.String("sqlEndpointPath"),
//   					},
//   					MariaDbParameters: &MariaDbParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					MySqlParameters: &MySqlParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					OracleParameters: &OracleParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					PostgreSqlParameters: &PostgreSqlParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					PrestoParameters: &PrestoParametersProperty{
//   						Catalog: jsii.String("catalog"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					RdsParameters: &RdsParametersProperty{
//   						Database: jsii.String("database"),
//   						InstanceId: jsii.String("instanceId"),
//   					},
//   					RedshiftParameters: &RedshiftParametersProperty{
//   						Database: jsii.String("database"),
//
//   						// the properties below are optional
//   						ClusterId: jsii.String("clusterId"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					S3Parameters: &S3ParametersProperty{
//   						ManifestFileLocation: &ManifestFileLocationProperty{
//   							Bucket: jsii.String("bucket"),
//   							Key: jsii.String("key"),
//   						},
//
//   						// the properties below are optional
//   						RoleArn: jsii.String("roleArn"),
//   					},
//   					SnowflakeParameters: &SnowflakeParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Warehouse: jsii.String("warehouse"),
//   					},
//   					SparkParameters: &SparkParametersProperty{
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					SqlServerParameters: &SqlServerParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					TeradataParameters: &TeradataParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	DataSourceId: jsii.String("dataSourceId"),
//   	DataSourceParameters: &DataSourceParametersProperty{
//   		AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   			Domain: jsii.String("domain"),
//   		},
//   		AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   			Domain: jsii.String("domain"),
//   		},
//   		AthenaParameters: &AthenaParametersProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			WorkGroup: jsii.String("workGroup"),
//   		},
//   		AuroraParameters: &AuroraParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		AuroraPostgreSqlParameters: &AuroraPostgreSqlParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		DatabricksParameters: &DatabricksParametersProperty{
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   			SqlEndpointPath: jsii.String("sqlEndpointPath"),
//   		},
//   		MariaDbParameters: &MariaDbParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		MySqlParameters: &MySqlParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		OracleParameters: &OracleParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		PostgreSqlParameters: &PostgreSqlParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		PrestoParameters: &PrestoParametersProperty{
//   			Catalog: jsii.String("catalog"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		RdsParameters: &RdsParametersProperty{
//   			Database: jsii.String("database"),
//   			InstanceId: jsii.String("instanceId"),
//   		},
//   		RedshiftParameters: &RedshiftParametersProperty{
//   			Database: jsii.String("database"),
//
//   			// the properties below are optional
//   			ClusterId: jsii.String("clusterId"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		S3Parameters: &S3ParametersProperty{
//   			ManifestFileLocation: &ManifestFileLocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//
//   			// the properties below are optional
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		SnowflakeParameters: &SnowflakeParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Warehouse: jsii.String("warehouse"),
//   		},
//   		SparkParameters: &SparkParametersProperty{
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		SqlServerParameters: &SqlServerParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		TeradataParameters: &TeradataParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   	ErrorInfo: &DataSourceErrorInfoProperty{
//   		Message: jsii.String("message"),
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	SslProperties: &SslPropertiesProperty{
//   		DisableSsl: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	VpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   		VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html
//
type CfnDataSourceProps struct {
	// A set of alternate data source parameters that you want to share for the credentials stored with this data source.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the credentials from this existing data source. If the `AlternateDataSourceParameters` list is null, the `Credentials` originally used with this `DataSourceParameters` are automatically allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-alternatedatasourceparameters
	//
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The credentials Amazon QuickSight that uses to connect to your underlying source.
	//
	// Currently, only credentials based on user name and password are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// An ID for the data source.
	//
	// This ID is unique per AWS Region for each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceid
	//
	DataSourceId *string `field:"optional" json:"dataSourceId" yaml:"dataSourceId"`
	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceparameters
	//
	DataSourceParameters interface{} `field:"optional" json:"dataSourceParameters" yaml:"dataSourceParameters"`
	// Error information from the last update or the creation of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-errorinfo
	//
	ErrorInfo interface{} `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// A display name for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions on the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-sslproperties
	//
	SslProperties interface{} `field:"optional" json:"sslProperties" yaml:"sslProperties"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the data source. To return a list of all data sources, use `ListDataSources` .
	//
	// Use `AMAZON_ELASTICSEARCH` for Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-vpcconnectionproperties
	//
	VpcConnectionProperties interface{} `field:"optional" json:"vpcConnectionProperties" yaml:"vpcConnectionProperties"`
}

