package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a knowledge base as a resource in a top-level template. Minimally, you must specify the following properties:.
//
// - Name – Specify a name for the knowledge base.
// - RoleArn – Specify the Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the knowledge base. For more information, see [Create a service role for Knowledge base for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/kb-permissions.html) .
// - KnowledgeBaseConfiguration – Specify the embeddings configuration of the knowledge base. The following sub-properties are required:
//
// - Type – Specify the value `VECTOR` .
// - StorageConfiguration – Specify information about the vector store in which the data source is stored. The following sub-properties are required:
//
// - Type – Specify the vector store service that you are using.
//
// > Redis Enterprise Cloud vector stores are currently unsupported in CloudFormation .
//
// For more information about using knowledge bases in Amazon Bedrock , see [Knowledge base for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBasePropsMixin := awscdkmixinspreview.Mixins.NewCfnKnowledgeBasePropsMixin(&CfnKnowledgeBaseMixinProps{
//   	Description: jsii.String("description"),
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		KendraKnowledgeBaseConfiguration: &KendraKnowledgeBaseConfigurationProperty{
//   			KendraIndexArn: jsii.String("kendraIndexArn"),
//   		},
//   		SqlKnowledgeBaseConfiguration: &SqlKnowledgeBaseConfigurationProperty{
//   			RedshiftConfiguration: &RedshiftConfigurationProperty{
//   				QueryEngineConfiguration: &RedshiftQueryEngineConfigurationProperty{
//   					ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   						AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   							DatabaseUser: jsii.String("databaseUser"),
//   							Type: jsii.String("type"),
//   							UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   						},
//   						ClusterIdentifier: jsii.String("clusterIdentifier"),
//   					},
//   					ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   						AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   							Type: jsii.String("type"),
//   							UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   						},
//   						WorkgroupArn: jsii.String("workgroupArn"),
//   					},
//   					Type: jsii.String("type"),
//   				},
//   				QueryGenerationConfiguration: &QueryGenerationConfigurationProperty{
//   					ExecutionTimeoutSeconds: jsii.Number(123),
//   					GenerationContext: &QueryGenerationContextProperty{
//   						CuratedQueries: []interface{}{
//   							&CuratedQueryProperty{
//   								NaturalLanguage: jsii.String("naturalLanguage"),
//   								Sql: jsii.String("sql"),
//   							},
//   						},
//   						Tables: []interface{}{
//   							&QueryGenerationTableProperty{
//   								Columns: []interface{}{
//   									&QueryGenerationColumnProperty{
//   										Description: jsii.String("description"),
//   										Inclusion: jsii.String("inclusion"),
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Description: jsii.String("description"),
//   								Inclusion: jsii.String("inclusion"),
//   								Name: jsii.String("name"),
//   							},
//   						},
//   					},
//   				},
//   				StorageConfigurations: []interface{}{
//   					&RedshiftQueryEngineStorageConfigurationProperty{
//   						AwsDataCatalogConfiguration: &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   							TableNames: []*string{
//   								jsii.String("tableNames"),
//   							},
//   						},
//   						RedshiftConfiguration: &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   							DatabaseName: jsii.String("databaseName"),
//   						},
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		Type: jsii.String("type"),
//   		VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   			EmbeddingModelArn: jsii.String("embeddingModelArn"),
//   			EmbeddingModelConfiguration: &EmbeddingModelConfigurationProperty{
//   				BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   					Audio: []interface{}{
//   						&AudioConfigurationProperty{
//   							SegmentationConfiguration: &AudioSegmentationConfigurationProperty{
//   								FixedLengthDuration: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Dimensions: jsii.Number(123),
//   					EmbeddingDataType: jsii.String("embeddingDataType"),
//   					Video: []interface{}{
//   						&VideoConfigurationProperty{
//   							SegmentationConfiguration: &VideoSegmentationConfigurationProperty{
//   								FixedLengthDuration: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			SupplementalDataStorageConfiguration: &SupplementalDataStorageConfigurationProperty{
//   				SupplementalDataStorageLocations: []interface{}{
//   					&SupplementalDataStorageLocationProperty{
//   						S3Location: &S3LocationProperty{
//   							Uri: jsii.String("uri"),
//   						},
//   						SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	StorageConfiguration: &StorageConfigurationProperty{
//   		MongoDbAtlasConfiguration: &MongoDbAtlasConfigurationProperty{
//   			CollectionName: jsii.String("collectionName"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Endpoint: jsii.String("endpoint"),
//   			EndpointServiceName: jsii.String("endpointServiceName"),
//   			FieldMapping: &MongoDbAtlasFieldMappingProperty{
//   				MetadataField: jsii.String("metadataField"),
//   				TextField: jsii.String("textField"),
//   				VectorField: jsii.String("vectorField"),
//   			},
//   			TextIndexName: jsii.String("textIndexName"),
//   			VectorIndexName: jsii.String("vectorIndexName"),
//   		},
//   		NeptuneAnalyticsConfiguration: &NeptuneAnalyticsConfigurationProperty{
//   			FieldMapping: &NeptuneAnalyticsFieldMappingProperty{
//   				MetadataField: jsii.String("metadataField"),
//   				TextField: jsii.String("textField"),
//   			},
//   			GraphArn: jsii.String("graphArn"),
//   		},
//   		OpensearchManagedClusterConfiguration: &OpenSearchManagedClusterConfigurationProperty{
//   			DomainArn: jsii.String("domainArn"),
//   			DomainEndpoint: jsii.String("domainEndpoint"),
//   			FieldMapping: &OpenSearchManagedClusterFieldMappingProperty{
//   				MetadataField: jsii.String("metadataField"),
//   				TextField: jsii.String("textField"),
//   				VectorField: jsii.String("vectorField"),
//   			},
//   			VectorIndexName: jsii.String("vectorIndexName"),
//   		},
//   		OpensearchServerlessConfiguration: &OpenSearchServerlessConfigurationProperty{
//   			CollectionArn: jsii.String("collectionArn"),
//   			FieldMapping: &OpenSearchServerlessFieldMappingProperty{
//   				MetadataField: jsii.String("metadataField"),
//   				TextField: jsii.String("textField"),
//   				VectorField: jsii.String("vectorField"),
//   			},
//   			VectorIndexName: jsii.String("vectorIndexName"),
//   		},
//   		PineconeConfiguration: &PineconeConfigurationProperty{
//   			ConnectionString: jsii.String("connectionString"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			FieldMapping: &PineconeFieldMappingProperty{
//   				MetadataField: jsii.String("metadataField"),
//   				TextField: jsii.String("textField"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   		},
//   		RdsConfiguration: &RdsConfigurationProperty{
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			FieldMapping: &RdsFieldMappingProperty{
//   				CustomMetadataField: jsii.String("customMetadataField"),
//   				MetadataField: jsii.String("metadataField"),
//   				PrimaryKeyField: jsii.String("primaryKeyField"),
//   				TextField: jsii.String("textField"),
//   				VectorField: jsii.String("vectorField"),
//   			},
//   			ResourceArn: jsii.String("resourceArn"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		S3VectorsConfiguration: &S3VectorsConfigurationProperty{
//   			IndexArn: jsii.String("indexArn"),
//   			IndexName: jsii.String("indexName"),
//   			VectorBucketArn: jsii.String("vectorBucketArn"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html
//
type CfnKnowledgeBasePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnKnowledgeBaseMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnKnowledgeBasePropsMixin
type jsiiProxy_CfnKnowledgeBasePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnKnowledgeBasePropsMixin) Props() *CfnKnowledgeBaseMixinProps {
	var returns *CfnKnowledgeBaseMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBasePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::KnowledgeBase`.
func NewCfnKnowledgeBasePropsMixin(props *CfnKnowledgeBaseMixinProps, options *mixins.CfnPropertyMixinOptions) CfnKnowledgeBasePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnKnowledgeBasePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKnowledgeBasePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBasePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::KnowledgeBase`.
func NewCfnKnowledgeBasePropsMixin_Override(c CfnKnowledgeBasePropsMixin, props *CfnKnowledgeBaseMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBasePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnKnowledgeBasePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKnowledgeBasePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBasePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKnowledgeBasePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBasePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKnowledgeBasePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnKnowledgeBasePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

