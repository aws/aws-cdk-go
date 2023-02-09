package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &cfnDataSourceProps{
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
//   	awsAccountId: jsii.String("awsAccountId"),
//   	credentials: &dataSourceCredentialsProperty{
//   		copySourceArn: jsii.String("copySourceArn"),
//   		credentialPair: &credentialPairProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//
//   			// the properties below are optional
//   			alternateDataSourceParameters: []interface{}{
//   				&dataSourceParametersProperty{
//   					amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					athenaParameters: &athenaParametersProperty{
//   						workGroup: jsii.String("workGroup"),
//   					},
//   					auroraParameters: &auroraParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					databricksParameters: &databricksParametersProperty{
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   						sqlEndpointPath: jsii.String("sqlEndpointPath"),
//   					},
//   					mariaDbParameters: &mariaDbParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					mySqlParameters: &mySqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					oracleParameters: &oracleParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					postgreSqlParameters: &postgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					prestoParameters: &prestoParametersProperty{
//   						catalog: jsii.String("catalog"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					rdsParameters: &rdsParametersProperty{
//   						database: jsii.String("database"),
//   						instanceId: jsii.String("instanceId"),
//   					},
//   					redshiftParameters: &redshiftParametersProperty{
//   						database: jsii.String("database"),
//
//   						// the properties below are optional
//   						clusterId: jsii.String("clusterId"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					s3Parameters: &s3ParametersProperty{
//   						manifestFileLocation: &manifestFileLocationProperty{
//   							bucket: jsii.String("bucket"),
//   							key: jsii.String("key"),
//   						},
//   					},
//   					snowflakeParameters: &snowflakeParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						warehouse: jsii.String("warehouse"),
//   					},
//   					sparkParameters: &sparkParametersProperty{
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					sqlServerParameters: &sqlServerParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					teradataParameters: &teradataParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		secretArn: jsii.String("secretArn"),
//   	},
//   	dataSourceId: jsii.String("dataSourceId"),
//   	dataSourceParameters: &dataSourceParametersProperty{
//   		amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		athenaParameters: &athenaParametersProperty{
//   			workGroup: jsii.String("workGroup"),
//   		},
//   		auroraParameters: &auroraParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		databricksParameters: &databricksParametersProperty{
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   			sqlEndpointPath: jsii.String("sqlEndpointPath"),
//   		},
//   		mariaDbParameters: &mariaDbParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		mySqlParameters: &mySqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		oracleParameters: &oracleParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		postgreSqlParameters: &postgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		prestoParameters: &prestoParametersProperty{
//   			catalog: jsii.String("catalog"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		rdsParameters: &rdsParametersProperty{
//   			database: jsii.String("database"),
//   			instanceId: jsii.String("instanceId"),
//   		},
//   		redshiftParameters: &redshiftParametersProperty{
//   			database: jsii.String("database"),
//
//   			// the properties below are optional
//   			clusterId: jsii.String("clusterId"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		s3Parameters: &s3ParametersProperty{
//   			manifestFileLocation: &manifestFileLocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   		},
//   		snowflakeParameters: &snowflakeParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			warehouse: jsii.String("warehouse"),
//   		},
//   		sparkParameters: &sparkParametersProperty{
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		sqlServerParameters: &sqlServerParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		teradataParameters: &teradataParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   	errorInfo: &dataSourceErrorInfoProperty{
//   		message: jsii.String("message"),
//   		type: jsii.String("type"),
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	sslProperties: &sslPropertiesProperty{
//   		disableSsl: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   	vpcConnectionProperties: &vpcConnectionPropertiesProperty{
//   		vpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   }
//
type CfnDataSourceProps struct {
	// A set of alternate data source parameters that you want to share for the credentials stored with this data source.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the credentials from this existing data source. If the `AlternateDataSourceParameters` list is null, the `Credentials` originally used with this `DataSourceParameters` are automatically allowed.
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
	// The AWS account ID.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The credentials Amazon QuickSight that uses to connect to your underlying source.
	//
	// Currently, only credentials based on user name and password are supported.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// An ID for the data source.
	//
	// This ID is unique per AWS Region for each AWS account.
	DataSourceId *string `field:"optional" json:"dataSourceId" yaml:"dataSourceId"`
	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	DataSourceParameters interface{} `field:"optional" json:"dataSourceParameters" yaml:"dataSourceParameters"`
	// Error information from the last update or the creation of the data source.
	ErrorInfo interface{} `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// A display name for the data source.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions on the data source.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
	SslProperties interface{} `field:"optional" json:"sslProperties" yaml:"sslProperties"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the data source. To return a list of all data sources, use `ListDataSources` .
	//
	// Use `AMAZON_ELASTICSEARCH` for Amazon OpenSearch Service.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
	VpcConnectionProperties interface{} `field:"optional" json:"vpcConnectionProperties" yaml:"vpcConnectionProperties"`
}

