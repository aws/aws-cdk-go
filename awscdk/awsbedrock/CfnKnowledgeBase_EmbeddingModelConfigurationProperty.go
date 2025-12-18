package awsbedrock


// The configuration details for the embeddings model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   embeddingModelConfigurationProperty := &EmbeddingModelConfigurationProperty{
//   	BedrockEmbeddingModelConfiguration: &BedrockEmbeddingModelConfigurationProperty{
//   		Audio: []interface{}{
//   			&AudioConfigurationProperty{
//   				SegmentationConfiguration: &AudioSegmentationConfigurationProperty{
//   					FixedLengthDuration: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Dimensions: jsii.Number(123),
//   		EmbeddingDataType: jsii.String("embeddingDataType"),
//   		Video: []interface{}{
//   			&VideoConfigurationProperty{
//   				SegmentationConfiguration: &VideoSegmentationConfigurationProperty{
//   					FixedLengthDuration: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-embeddingmodelconfiguration.html
//
type CfnKnowledgeBase_EmbeddingModelConfigurationProperty struct {
	// The vector configuration details on the Bedrock embeddings model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-embeddingmodelconfiguration.html#cfn-bedrock-knowledgebase-embeddingmodelconfiguration-bedrockembeddingmodelconfiguration
	//
	BedrockEmbeddingModelConfiguration interface{} `field:"optional" json:"bedrockEmbeddingModelConfiguration" yaml:"bedrockEmbeddingModelConfiguration"`
}

