package previewawsbedrockmixins


// Properties for CfnKnowledgeBasePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseMixinProps := &CfnKnowledgeBaseMixinProps{
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
//   					Dimensions: jsii.Number(123),
//   					EmbeddingDataType: jsii.String("embeddingDataType"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html
//
type CfnKnowledgeBaseMixinProps struct {
	// The description of the knowledge base associated with the inline agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Contains details about the embeddings configuration of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration
	//
	KnowledgeBaseConfiguration interface{} `field:"optional" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// The name of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Contains details about the storage configuration of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-storageconfiguration
	//
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

