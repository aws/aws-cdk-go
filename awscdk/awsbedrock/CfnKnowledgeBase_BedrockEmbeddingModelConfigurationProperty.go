package awsbedrock


// The vector configuration details for the Bedrock embeddings model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bedrockEmbeddingModelConfigurationProperty := &BedrockEmbeddingModelConfigurationProperty{
//   	Dimensions: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-bedrockembeddingmodelconfiguration.html
//
type CfnKnowledgeBase_BedrockEmbeddingModelConfigurationProperty struct {
	// The dimensions details for the vector configuration used on the Bedrock embeddings model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-bedrockembeddingmodelconfiguration.html#cfn-bedrock-knowledgebase-bedrockembeddingmodelconfiguration-dimensions
	//
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
}

