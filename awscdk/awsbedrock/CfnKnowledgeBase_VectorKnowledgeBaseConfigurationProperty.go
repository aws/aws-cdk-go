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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-vectorknowledgebaseconfiguration.html
//
type CfnKnowledgeBase_VectorKnowledgeBaseConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the model used to create vector embeddings for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-vectorknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-vectorknowledgebaseconfiguration-embeddingmodelarn
	//
	EmbeddingModelArn *string `field:"required" json:"embeddingModelArn" yaml:"embeddingModelArn"`
}

