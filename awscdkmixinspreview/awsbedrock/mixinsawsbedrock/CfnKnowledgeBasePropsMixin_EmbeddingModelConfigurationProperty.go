package mixinsawsbedrock


// The configuration details for the embeddings model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   embeddingModelConfigurationProperty := &EmbeddingModelConfigurationProperty{
//   	BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   		Dimensions: jsii.Number(123),
//   		EmbeddingDataType: jsii.String("embeddingDataType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-embeddingmodelconfiguration.html
//
type CfnKnowledgeBasePropsMixin_EmbeddingModelConfigurationProperty struct {
	// The vector configuration details on the Bedrock embeddings model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-embeddingmodelconfiguration.html#cfn-bedrock-knowledgebase-embeddingmodelconfiguration-bedrockembeddingmodelconfiguration
	//
	BedrockEmbeddingModelConfiguration interface{} `field:"optional" json:"bedrockEmbeddingModelConfiguration" yaml:"bedrockEmbeddingModelConfiguration"`
}

