package awsbedrock


// Properties for defining a `CfnKnowledgeBase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKnowledgeBaseProps := &CfnKnowledgeBaseProps{
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		KendraKnowledgeBaseConfiguration: &KendraKnowledgeBaseConfigurationProperty{
//   			KendraIndexArn: jsii.String("kendraIndexArn"),
//   		},
//   		SqlKnowledgeBaseConfiguration: &SqlKnowledgeBaseConfigurationProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			RedshiftConfiguration: &RedshiftConfigurationProperty{
//   				QueryEngineConfiguration: &RedshiftQueryEngineConfigurationProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   						AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							DatabaseUser: jsii.String("databaseUser"),
//   							UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   						},
//   						ClusterIdentifier: jsii.String("clusterIdentifier"),
//   					},
//   					ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   						AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   						},
//   						WorkgroupArn: jsii.String("workgroupArn"),
//   					},
//   				},
//   				StorageConfigurations: []interface{}{
//   					&RedshiftQueryEngineStorageConfigurationProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						AwsDataCatalogConfiguration: &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   							TableNames: []*string{
//   								jsii.String("tableNames"),
//   							},
//   						},
//   						RedshiftConfiguration: &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   							DatabaseName: jsii.String("databaseName"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
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
//   								Name: jsii.String("name"),
//
//   								// the properties below are optional
//   								Columns: []interface{}{
//   									&QueryGenerationColumnProperty{
//   										Description: jsii.String("description"),
//   										Inclusion: jsii.String("inclusion"),
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Description: jsii.String("description"),
//   								Inclusion: jsii.String("inclusion"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   			EmbeddingModelArn: jsii.String("embeddingModelArn"),
//
//   			// the properties below are optional
//   			EmbeddingModelConfiguration: &EmbeddingModelConfigurationProperty{
//   				BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   					Dimensions: jsii.Number(123),
//   					EmbeddingDataType: jsii.String("embeddingDataType"),
//   				},
//   			},
//   			SupplementalDataStorageConfiguration: &SupplementalDataStorageConfigurationProperty{
//   				SupplementalDataStorageLocations: []interface{}{
//   					&SupplementalDataStorageLocationProperty{
//   						SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//
//   						// the properties below are optional
//   						S3Location: &S3LocationProperty{
//   							Uri: jsii.String("uri"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	StorageConfiguration: &StorageConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		MongoDbAtlasConfiguration: &MongoDbAtlasConfigurationProperty{
//   			CollectionName: jsii.String("collectionName"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Endpoint: jsii.String("endpoint"),
//   			FieldMapping: &MongoDbAtlasFieldMappingProperty{
//   				MetadataField: jsii.String("metadataField"),
//   				TextField: jsii.String("textField"),
//   				VectorField: jsii.String("vectorField"),
//   			},
//   			VectorIndexName: jsii.String("vectorIndexName"),
//
//   			// the properties below are optional
//   			EndpointServiceName: jsii.String("endpointServiceName"),
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
//
//   			// the properties below are optional
//   			Namespace: jsii.String("namespace"),
//   		},
//   		RdsConfiguration: &RdsConfigurationProperty{
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			FieldMapping: &RdsFieldMappingProperty{
//   				MetadataField: jsii.String("metadataField"),
//   				PrimaryKeyField: jsii.String("primaryKeyField"),
//   				TextField: jsii.String("textField"),
//   				VectorField: jsii.String("vectorField"),
//   			},
//   			ResourceArn: jsii.String("resourceArn"),
//   			TableName: jsii.String("tableName"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html
//
type CfnKnowledgeBaseProps struct {
	// Contains details about the embeddings configuration of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration
	//
	KnowledgeBaseConfiguration interface{} `field:"required" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// The name of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The description of the knowledge base associated with the inline agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
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

