package previewawsbedrockmixins


// The vector configuration details for the Bedrock embeddings model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockEmbeddingModelConfigurationProperty := &BedrockEmbeddingModelConfigurationProperty{
//   	Dimensions: jsii.Number(123),
//   	EmbeddingDataType: jsii.String("embeddingDataType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-bedrockembeddingmodelconfiguration.html
//
type CfnKnowledgeBasePropsMixin_BedrockEmbeddingModelConfigurationProperty struct {
	// The dimensions details for the vector configuration used on the Bedrock embeddings model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-bedrockembeddingmodelconfiguration.html#cfn-bedrock-knowledgebase-bedrockembeddingmodelconfiguration-dimensions
	//
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The data type for the vectors when using a model to convert text into vector embeddings.
	//
	// The model must support the specified data type for vector embeddings. Floating-point (float32) is the default data type, and is supported by most models for vector embeddings. See [Supported embeddings models](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-supported.html) for information on the available models and their vector data types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-bedrockembeddingmodelconfiguration.html#cfn-bedrock-knowledgebase-bedrockembeddingmodelconfiguration-embeddingdatatype
	//
	EmbeddingDataType *string `field:"optional" json:"embeddingDataType" yaml:"embeddingDataType"`
}

