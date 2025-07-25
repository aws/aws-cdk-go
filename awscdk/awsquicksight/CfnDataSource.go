package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSource := awscdk.Aws_quicksight.NewCfnDataSource(this, jsii.String("MyCfnDataSource"), &CfnDataSourceProps{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
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
//
//   				// the properties below are optional
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
//   				Database: jsii.String("database"),
//
//   				// the properties below are optional
//   				ClusterId: jsii.String("clusterId"),
//   				Host: jsii.String("host"),
//   				IamParameters: &RedshiftIAMParametersProperty{
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					AutoCreateDatabaseUser: jsii.Boolean(false),
//   					DatabaseGroups: []*string{
//   						jsii.String("databaseGroups"),
//   					},
//   					DatabaseUser: jsii.String("databaseUser"),
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
//
//   				// the properties below are optional
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			SnowflakeParameters: &SnowflakeParametersProperty{
//   				Database: jsii.String("database"),
//   				Host: jsii.String("host"),
//   				Warehouse: jsii.String("warehouse"),
//
//   				// the properties below are optional
//   				AuthenticationType: jsii.String("authenticationType"),
//   				DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   				OAuthParameters: &OAuthParametersProperty{
//   					TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   					// the properties below are optional
//   					IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   					IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   						VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   					},
//   					OAuthScope: jsii.String("oAuthScope"),
//   				},
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
//   				Catalog: jsii.String("catalog"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//
//   				// the properties below are optional
//   				AuthenticationType: jsii.String("authenticationType"),
//   				DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   				OAuthParameters: &OAuthParametersProperty{
//   					TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   					// the properties below are optional
//   					IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   					IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   						VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   					},
//   					OAuthScope: jsii.String("oAuthScope"),
//   				},
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
//
//   						// the properties below are optional
//   						UseServiceName: jsii.Boolean(false),
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
//   						IamParameters: &RedshiftIAMParametersProperty{
//   							RoleArn: jsii.String("roleArn"),
//
//   							// the properties below are optional
//   							AutoCreateDatabaseUser: jsii.Boolean(false),
//   							DatabaseGroups: []*string{
//   								jsii.String("databaseGroups"),
//   							},
//   							DatabaseUser: jsii.String("databaseUser"),
//   						},
//   						IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   							EnableIdentityPropagation: jsii.Boolean(false),
//   						},
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
//
//   						// the properties below are optional
//   						AuthenticationType: jsii.String("authenticationType"),
//   						DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   						OAuthParameters: &OAuthParametersProperty{
//   							TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   							// the properties below are optional
//   							IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   							IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   								VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   							},
//   							OAuthScope: jsii.String("oAuthScope"),
//   						},
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
//   					StarburstParameters: &StarburstParametersProperty{
//   						Catalog: jsii.String("catalog"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//
//   						// the properties below are optional
//   						AuthenticationType: jsii.String("authenticationType"),
//   						DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   						OAuthParameters: &OAuthParametersProperty{
//   							TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   							// the properties below are optional
//   							IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   							IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   								VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   							},
//   							OAuthScope: jsii.String("oAuthScope"),
//   						},
//   						ProductType: jsii.String("productType"),
//   					},
//   					TeradataParameters: &TeradataParametersProperty{
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   					TrinoParameters: &TrinoParametersProperty{
//   						Catalog: jsii.String("catalog"),
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
//
//   			// the properties below are optional
//   			UseServiceName: jsii.Boolean(false),
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
//   			IamParameters: &RedshiftIAMParametersProperty{
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				AutoCreateDatabaseUser: jsii.Boolean(false),
//   				DatabaseGroups: []*string{
//   					jsii.String("databaseGroups"),
//   				},
//   				DatabaseUser: jsii.String("databaseUser"),
//   			},
//   			IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   				EnableIdentityPropagation: jsii.Boolean(false),
//   			},
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
//
//   			// the properties below are optional
//   			AuthenticationType: jsii.String("authenticationType"),
//   			DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   			OAuthParameters: &OAuthParametersProperty{
//   				TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   				// the properties below are optional
//   				IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   				IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   					VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   				},
//   				OAuthScope: jsii.String("oAuthScope"),
//   			},
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
//   		StarburstParameters: &StarburstParametersProperty{
//   			Catalog: jsii.String("catalog"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//
//   			// the properties below are optional
//   			AuthenticationType: jsii.String("authenticationType"),
//   			DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   			OAuthParameters: &OAuthParametersProperty{
//   				TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   				// the properties below are optional
//   				IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   				IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   					VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   				},
//   				OAuthScope: jsii.String("oAuthScope"),
//   			},
//   			ProductType: jsii.String("productType"),
//   		},
//   		TeradataParameters: &TeradataParametersProperty{
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		TrinoParameters: &TrinoParametersProperty{
//   			Catalog: jsii.String("catalog"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   	ErrorInfo: &DataSourceErrorInfoProperty{
//   		Message: jsii.String("message"),
//   		Type: jsii.String("type"),
//   	},
//   	FolderArns: []*string{
//   		jsii.String("folderArns"),
//   	},
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//
//   			// the properties below are optional
//   			Resource: jsii.String("resource"),
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
//   	VpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   		VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html
//
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// A set of alternate data source parameters that you want to share for the credentials stored with this data source.
	AlternateDataSourceParameters() interface{}
	SetAlternateDataSourceParameters(val interface{})
	// The Amazon Resource Name (ARN) of the dataset.
	AttrArn() *string
	// The time that this data source was created.
	AttrCreatedTime() *string
	// The last time that this data source was updated.
	AttrLastUpdatedTime() *string
	// The HTTP status of the request.
	AttrStatus() *string
	// The AWS account ID.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The credentials Amazon QuickSight that uses to connect to your underlying source.
	Credentials() interface{}
	SetCredentials(val interface{})
	// An ID for the data source.
	DataSourceId() *string
	SetDataSourceId(val *string)
	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	DataSourceParameters() interface{}
	SetDataSourceParameters(val interface{})
	// Error information from the last update or the creation of the data source.
	ErrorInfo() interface{}
	SetErrorInfo(val interface{})
	FolderArns() *[]*string
	SetFolderArns(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// A display name for the data source.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A list of resource permissions on the data source.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
	SslProperties() interface{}
	SetSslProperties(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The type of the data source.
	//
	// To return a list of all data sources, use `ListDataSources` .
	Type() *string
	SetType(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
	VpcConnectionProperties() interface{}
	SetVpcConnectionProperties(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnDataSource) AlternateDataSourceParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alternateDataSourceParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Credentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ErrorInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) FolderArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"folderArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) SslProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) VpcConnectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConnectionProperties",
		&returns,
	)
	return returns
}


func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	if err := validateNewCfnDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource)SetAlternateDataSourceParameters(val interface{}) {
	if err := j.validateSetAlternateDataSourceParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternateDataSourceParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetCredentials(val interface{}) {
	if err := j.validateSetCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetDataSourceId(val *string) {
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetDataSourceParameters(val interface{}) {
	if err := j.validateSetDataSourceParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetErrorInfo(val interface{}) {
	if err := j.validateSetErrorInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorInfo",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetFolderArns(val *[]*string) {
	_jsii_.Set(
		j,
		"folderArns",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetPermissions(val interface{}) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetSslProperties(val interface{}) {
	if err := j.validateSetSslPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetVpcConnectionProperties(val interface{}) {
	if err := j.validateSetVpcConnectionPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnectionProperties",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDataSource_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSource) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

