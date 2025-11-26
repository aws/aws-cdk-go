package previewawsquicksightmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsquicksight/previewawsquicksightmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnDataSourcePropsMixin(&CfnDataSourceMixinProps{
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
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Credentials: &DataSourceCredentialsProperty{
//   		CopySourceArn: jsii.String("copySourceArn"),
//   		CredentialPair: &CredentialPairProperty{
//   			AlternateDataSourceParameters: []interface{}{
//   				&DataSourceParametersProperty{
//   					AmazonElasticsearchParameters: &AmazonElasticsearchParametersProperty{
//   						Domain: jsii.String("domain"),
//   					},
//   					AmazonOpenSearchParameters: &AmazonOpenSearchParametersProperty{
//   						Domain: jsii.String("domain"),
//   					},
//   					AthenaParameters: &AthenaParametersProperty{
//   						IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   							EnableIdentityPropagation: jsii.Boolean(false),
//   						},
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
//   						ClusterId: jsii.String("clusterId"),
//   						Database: jsii.String("database"),
//   						Host: jsii.String("host"),
//   						IamParameters: &RedshiftIAMParametersProperty{
//   							AutoCreateDatabaseUser: jsii.Boolean(false),
//   							DatabaseGroups: []*string{
//   								jsii.String("databaseGroups"),
//   							},
//   							DatabaseUser: jsii.String("databaseUser"),
//   							RoleArn: jsii.String("roleArn"),
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
//   						RoleArn: jsii.String("roleArn"),
//   					},
//   					SnowflakeParameters: &SnowflakeParametersProperty{
//   						AuthenticationType: jsii.String("authenticationType"),
//   						Database: jsii.String("database"),
//   						DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   						Host: jsii.String("host"),
//   						OAuthParameters: &OAuthParametersProperty{
//   							IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   							IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   								VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   							},
//   							OAuthScope: jsii.String("oAuthScope"),
//   							TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   						},
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
//   					StarburstParameters: &StarburstParametersProperty{
//   						AuthenticationType: jsii.String("authenticationType"),
//   						Catalog: jsii.String("catalog"),
//   						DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   						Host: jsii.String("host"),
//   						OAuthParameters: &OAuthParametersProperty{
//   							IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   							IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   								VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   							},
//   							OAuthScope: jsii.String("oAuthScope"),
//   							TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   						},
//   						Port: jsii.Number(123),
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
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
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
//   			IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   				EnableIdentityPropagation: jsii.Boolean(false),
//   			},
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
//   			ClusterId: jsii.String("clusterId"),
//   			Database: jsii.String("database"),
//   			Host: jsii.String("host"),
//   			IamParameters: &RedshiftIAMParametersProperty{
//   				AutoCreateDatabaseUser: jsii.Boolean(false),
//   				DatabaseGroups: []*string{
//   					jsii.String("databaseGroups"),
//   				},
//   				DatabaseUser: jsii.String("databaseUser"),
//   				RoleArn: jsii.String("roleArn"),
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
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		SnowflakeParameters: &SnowflakeParametersProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//   			Database: jsii.String("database"),
//   			DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   			Host: jsii.String("host"),
//   			OAuthParameters: &OAuthParametersProperty{
//   				IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   				IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   					VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   				},
//   				OAuthScope: jsii.String("oAuthScope"),
//   				TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   			},
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
//   		StarburstParameters: &StarburstParametersProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//   			Catalog: jsii.String("catalog"),
//   			DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   			Host: jsii.String("host"),
//   			OAuthParameters: &OAuthParametersProperty{
//   				IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   				IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   					VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   				},
//   				OAuthScope: jsii.String("oAuthScope"),
//   				TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   			},
//   			Port: jsii.Number(123),
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
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   			Resource: jsii.String("resource"),
//   		},
//   	},
//   	SslProperties: &SslPropertiesProperty{
//   		DisableSsl: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	VpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   		VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html
//
type CfnDataSourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataSourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataSourcePropsMixin
type jsiiProxy_CfnDataSourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Props() *CfnDataSourceMixinProps {
	var returns *CfnDataSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::DataSource`.
func NewCfnDataSourcePropsMixin(props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::DataSource`.
func NewCfnDataSourcePropsMixin_Override(c CfnDataSourcePropsMixin, props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnDataSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnDataSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSourcePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

