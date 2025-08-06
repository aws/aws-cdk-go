package awsbedrockalpha


// Properties for configuring a Bedrock Foundation Model.
//
// These properties define the model's capabilities and supported features.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   bedrockFoundationModelProps := &BedrockFoundationModelProps{
//   	Legacy: jsii.Boolean(false),
//   	OptimizedForAgents: jsii.Boolean(false),
//   	SupportedVectorType: []vectorType{
//   		bedrock_alpha.*vectorType_FLOATING_POINT,
//   	},
//   	SupportsAgents: jsii.Boolean(false),
//   	SupportsCrossRegion: jsii.Boolean(false),
//   	SupportsKnowledgeBase: jsii.Boolean(false),
//   	VectorDimensions: jsii.Number(123),
//   }
//
// Experimental.
type BedrockFoundationModelProps struct {
	// https://docs.aws.amazon.com/bedrock/latest/userguide/model-lifecycle.html A version is marked Legacy when there is a more recent version which provides superior performance. Amazon Bedrock sets an EOL date for Legacy versions.
	// Default: - false.
	//
	// Experimental.
	Legacy *bool `field:"optional" json:"legacy" yaml:"legacy"`
	// Currently, some of the offered models are optimized with prompts/parsers fine-tuned for integrating with the agents architecture.
	//
	// When true, the model has been specifically optimized for agent interactions.
	// Default: - false.
	//
	// Experimental.
	OptimizedForAgents *bool `field:"optional" json:"optimizedForAgents" yaml:"optimizedForAgents"`
	// Embeddings models have different supported vector types.
	//
	// Defines whether the model supports floating-point or binary vectors.
	// Default: - undefined.
	//
	// Experimental.
	SupportedVectorType *[]VectorType `field:"optional" json:"supportedVectorType" yaml:"supportedVectorType"`
	// Bedrock Agents can use this model.
	//
	// When true, the model can be used with Bedrock Agents for automated task execution.
	// Default: - false.
	//
	// Experimental.
	SupportsAgents *bool `field:"optional" json:"supportsAgents" yaml:"supportsAgents"`
	// Can be used with a Cross-Region Inference Profile.
	//
	// When true, the model supports inference across different AWS regions.
	// Default: - false.
	//
	// Experimental.
	SupportsCrossRegion *bool `field:"optional" json:"supportsCrossRegion" yaml:"supportsCrossRegion"`
	// Bedrock Knowledge Base can use this model.
	//
	// When true, the model can be used for knowledge base operations.
	// Default: - false.
	//
	// Experimental.
	SupportsKnowledgeBase *bool `field:"optional" json:"supportsKnowledgeBase" yaml:"supportsKnowledgeBase"`
	// Embedding models have different vector dimensions.
	//
	// Only applicable for embedding models. Defines the dimensionality of the vector embeddings.
	// Default: - undefined.
	//
	// Experimental.
	VectorDimensions *float64 `field:"optional" json:"vectorDimensions" yaml:"vectorDimensions"`
}

