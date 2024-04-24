package awsbedrock


// Contains details about the embeddings configuration of the knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   knowledgeBaseConfigurationProperty := &KnowledgeBaseConfigurationProperty{
//   	Type: jsii.String("type"),
//   	VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   		EmbeddingModelArn: jsii.String("embeddingModelArn"),
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
	// Contains details about the embeddings model that'sused to convert the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-knowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration-vectorknowledgebaseconfiguration
	//
	VectorKnowledgeBaseConfiguration interface{} `field:"required" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

