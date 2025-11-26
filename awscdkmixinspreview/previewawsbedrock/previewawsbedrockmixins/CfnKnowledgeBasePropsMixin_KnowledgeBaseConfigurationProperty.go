package previewawsbedrockmixins


// Configurations to apply to a knowledge base attached to the agent during query.
//
// For more information, see [Knowledge base retrieval configurations](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-session-state.html#session-state-kb) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   knowledgeBaseConfigurationProperty := &KnowledgeBaseConfigurationProperty{
//   	KendraKnowledgeBaseConfiguration: &KendraKnowledgeBaseConfigurationProperty{
//   		KendraIndexArn: jsii.String("kendraIndexArn"),
//   	},
//   	SqlKnowledgeBaseConfiguration: &SqlKnowledgeBaseConfigurationProperty{
//   		RedshiftConfiguration: &RedshiftConfigurationProperty{
//   			QueryEngineConfiguration: &RedshiftQueryEngineConfigurationProperty{
//   				ProvisionedConfiguration: &RedshiftProvisionedConfigurationProperty{
//   					AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   						DatabaseUser: jsii.String("databaseUser"),
//   						Type: jsii.String("type"),
//   						UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   					},
//   					ClusterIdentifier: jsii.String("clusterIdentifier"),
//   				},
//   				ServerlessConfiguration: &RedshiftServerlessConfigurationProperty{
//   					AuthConfiguration: &RedshiftServerlessAuthConfigurationProperty{
//   						Type: jsii.String("type"),
//   						UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   					},
//   					WorkgroupArn: jsii.String("workgroupArn"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			QueryGenerationConfiguration: &QueryGenerationConfigurationProperty{
//   				ExecutionTimeoutSeconds: jsii.Number(123),
//   				GenerationContext: &QueryGenerationContextProperty{
//   					CuratedQueries: []interface{}{
//   						&CuratedQueryProperty{
//   							NaturalLanguage: jsii.String("naturalLanguage"),
//   							Sql: jsii.String("sql"),
//   						},
//   					},
//   					Tables: []interface{}{
//   						&QueryGenerationTableProperty{
//   							Columns: []interface{}{
//   								&QueryGenerationColumnProperty{
//   									Description: jsii.String("description"),
//   									Inclusion: jsii.String("inclusion"),
//   									Name: jsii.String("name"),
//   								},
//   							},
//   							Description: jsii.String("description"),
//   							Inclusion: jsii.String("inclusion"),
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   			StorageConfigurations: []interface{}{
//   				&RedshiftQueryEngineStorageConfigurationProperty{
//   					AwsDataCatalogConfiguration: &RedshiftQueryEngineAwsDataCatalogStorageConfigurationProperty{
//   						TableNames: []*string{
//   							jsii.String("tableNames"),
//   						},
//   					},
//   					RedshiftConfiguration: &RedshiftQueryEngineRedshiftStorageConfigurationProperty{
//   						DatabaseName: jsii.String("databaseName"),
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Type: jsii.String("type"),
//   	VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   		EmbeddingModelArn: jsii.String("embeddingModelArn"),
//   		EmbeddingModelConfiguration: &EmbeddingModelConfigurationProperty{
//   			BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   				Dimensions: jsii.Number(123),
//   				EmbeddingDataType: jsii.String("embeddingDataType"),
//   			},
//   		},
//   		SupplementalDataStorageConfiguration: &SupplementalDataStorageConfigurationProperty{
//   			SupplementalDataStorageLocations: []interface{}{
//   				&SupplementalDataStorageLocationProperty{
//   					S3Location: &S3LocationProperty{
//   						Uri: jsii.String("uri"),
//   					},
//   					SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html
//
type CfnKnowledgeBasePropsMixin_KnowledgeBaseConfigurationProperty struct {
	// Settings for an Amazon Kendra knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-kendraknowledgebaseconfiguration
	//
	KendraKnowledgeBaseConfiguration interface{} `field:"optional" json:"kendraKnowledgeBaseConfiguration" yaml:"kendraKnowledgeBaseConfiguration"`
	// Specifies configurations for a knowledge base connected to an SQL database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-sqlknowledgebaseconfiguration
	//
	SqlKnowledgeBaseConfiguration interface{} `field:"optional" json:"sqlKnowledgeBaseConfiguration" yaml:"sqlKnowledgeBaseConfiguration"`
	// The type of data that the data source is converted into for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Contains details about the model that's used to convert the data source into vector embeddings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-vectorknowledgebaseconfiguration
	//
	VectorKnowledgeBaseConfiguration interface{} `field:"optional" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

