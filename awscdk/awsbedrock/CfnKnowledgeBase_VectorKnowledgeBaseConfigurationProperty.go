package awsbedrock


// Contains details about the model used to create vector embeddings for the knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vectorKnowledgeBaseConfigurationProperty := &VectorKnowledgeBaseConfigurationProperty{
//   	EmbeddingModelArn: jsii.String("embeddingModelArn"),
//
//   	// the properties below are optional
//   	EmbeddingModelConfiguration: &EmbeddingModelConfigurationProperty{
//   		BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   			Dimensions: jsii.Number(123),
//   		},
//   	},
//   	SupplementalDataStorageConfiguration: &SupplementalDataStorageConfigurationProperty{
//   		SupplementalDataStorageLocations: []interface{}{
//   			&SupplementalDataStorageLocationProperty{
//   				SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//
//   				// the properties below are optional
//   				S3Location: &S3LocationProperty{
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-vectorknowledgebaseconfiguration.html
//
type CfnKnowledgeBase_VectorKnowledgeBaseConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the model or inference profile used to create vector embeddings for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-vectorknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-vectorknowledgebaseconfiguration-embeddingmodelarn
	//
	EmbeddingModelArn *string `field:"required" json:"embeddingModelArn" yaml:"embeddingModelArn"`
	// The embeddings model configuration details for the vector model used in Knowledge Base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-vectorknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-vectorknowledgebaseconfiguration-embeddingmodelconfiguration
	//
	EmbeddingModelConfiguration interface{} `field:"optional" json:"embeddingModelConfiguration" yaml:"embeddingModelConfiguration"`
	// If you include multimodal data from your data source, use this object to specify configurations for the storage location of the images extracted from your documents.
	//
	// These images can be retrieved and returned to the end user. They can also be used in generation when using [RetrieveAndGenerate](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_RetrieveAndGenerate.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-vectorknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-vectorknowledgebaseconfiguration-supplementaldatastorageconfiguration
	//
	SupplementalDataStorageConfiguration interface{} `field:"optional" json:"supplementalDataStorageConfiguration" yaml:"supplementalDataStorageConfiguration"`
}

