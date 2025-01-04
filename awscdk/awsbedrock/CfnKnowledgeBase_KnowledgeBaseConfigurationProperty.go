package awsbedrock


// Configurations to apply to a knowledge base attached to the agent during query.
//
// For more information, see [Knowledge base retrieval configurations](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-session-state.html#session-state-kb) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   knowledgeBaseConfigurationProperty := &KnowledgeBaseConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	KendraKnowledgeBaseConfiguration: &KendraKnowledgeBaseConfigurationProperty{
//   		KendraIndexArn: jsii.String("kendraIndexArn"),
//   	},
//   	VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   		EmbeddingModelArn: jsii.String("embeddingModelArn"),
//
//   		// the properties below are optional
//   		EmbeddingModelConfiguration: &EmbeddingModelConfigurationProperty{
//   			BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   				Dimensions: jsii.Number(123),
//   			},
//   		},
//   		SupplementalDataStorageConfiguration: &SupplementalDataStorageConfigurationProperty{
//   			SupplementalDataStorageLocations: []interface{}{
//   				&SupplementalDataStorageLocationProperty{
//   					SupplementalDataStorageLocationType: jsii.String("supplementalDataStorageLocationType"),
//
//   					// the properties below are optional
//   					S3Location: &S3LocationProperty{
//   						Uri: jsii.String("uri"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html
//
type CfnKnowledgeBase_KnowledgeBaseConfigurationProperty struct {
	// The type of data that the data source is converted into for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Settings for an Amazon Kendra knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-kendraknowledgebaseconfiguration
	//
	KendraKnowledgeBaseConfiguration interface{} `field:"optional" json:"kendraKnowledgeBaseConfiguration" yaml:"kendraKnowledgeBaseConfiguration"`
	// Contains details about the model that's used to convert the data source into vector embeddings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-vectorknowledgebaseconfiguration
	//
	VectorKnowledgeBaseConfiguration interface{} `field:"optional" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

