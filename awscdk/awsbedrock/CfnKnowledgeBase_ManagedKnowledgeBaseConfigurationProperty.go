package awsbedrock


// Contains details about the model used to create vector embeddings for a managed knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedKnowledgeBaseConfigurationProperty := &ManagedKnowledgeBaseConfigurationProperty{
//   	EmbeddingModelArn: jsii.String("embeddingModelArn"),
//   	EmbeddingModelConfiguration: &EmbeddingModelConfigurationProperty{
//   		BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   			Audio: []interface{}{
//   				&AudioConfigurationProperty{
//   					SegmentationConfiguration: &AudioSegmentationConfigurationProperty{
//   						FixedLengthDuration: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Dimensions: jsii.Number(123),
//   			EmbeddingDataType: jsii.String("embeddingDataType"),
//   			Video: []interface{}{
//   				&VideoConfigurationProperty{
//   					SegmentationConfiguration: &VideoSegmentationConfigurationProperty{
//   						FixedLengthDuration: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	EmbeddingModelType: jsii.String("embeddingModelType"),
//   	ServerSideEncryptionConfiguration: &ManagedKnowledgeBaseServerSideEncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-managedknowledgebaseconfiguration.html
//
type CfnKnowledgeBase_ManagedKnowledgeBaseConfigurationProperty struct {
	// The ARN of the model used to create vector embeddings for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-managedknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-managedknowledgebaseconfiguration-embeddingmodelarn
	//
	EmbeddingModelArn *string `field:"optional" json:"embeddingModelArn" yaml:"embeddingModelArn"`
	// The embeddings model configuration details for the vector model used in Knowledge Base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-managedknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-managedknowledgebaseconfiguration-embeddingmodelconfiguration
	//
	EmbeddingModelConfiguration interface{} `field:"optional" json:"embeddingModelConfiguration" yaml:"embeddingModelConfiguration"`
	// The type of embedding model to use for the managed knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-managedknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-managedknowledgebaseconfiguration-embeddingmodeltype
	//
	EmbeddingModelType *string `field:"optional" json:"embeddingModelType" yaml:"embeddingModelType"`
	// Contains details about the server-side encryption for the managed knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-managedknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-managedknowledgebaseconfiguration-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

