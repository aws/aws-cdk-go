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
//   		VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   			EmbeddingModelArn: jsii.String("embeddingModelArn"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	StorageConfiguration: &StorageConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
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
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
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
	// Contains details about the storage configuration of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-storageconfiguration
	//
	StorageConfiguration interface{} `field:"required" json:"storageConfiguration" yaml:"storageConfiguration"`
	// The description of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

