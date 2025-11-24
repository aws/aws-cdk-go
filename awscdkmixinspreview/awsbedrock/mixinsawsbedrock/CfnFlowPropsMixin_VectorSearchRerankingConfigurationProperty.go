package mixinsawsbedrock


// Configuration for reranking vector search results to improve relevance.
//
// Reranking applies additional relevance models to reorder the initial vector search results based on more sophisticated criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//
//   vectorSearchRerankingConfigurationProperty := &VectorSearchRerankingConfigurationProperty{
//   	BedrockRerankingConfiguration: &VectorSearchBedrockRerankingConfigurationProperty{
//   		MetadataConfiguration: &MetadataConfigurationForRerankingProperty{
//   			SelectionMode: jsii.String("selectionMode"),
//   			SelectiveModeConfiguration: &RerankingMetadataSelectiveModeConfigurationProperty{
//   				FieldsToExclude: []interface{}{
//   					&FieldForRerankingProperty{
//   						FieldName: jsii.String("fieldName"),
//   					},
//   				},
//   				FieldsToInclude: []interface{}{
//   					&FieldForRerankingProperty{
//   						FieldName: jsii.String("fieldName"),
//   					},
//   				},
//   			},
//   		},
//   		ModelConfiguration: &VectorSearchBedrockRerankingModelConfigurationProperty{
//   			AdditionalModelRequestFields: additionalModelRequestFields,
//   			ModelArn: jsii.String("modelArn"),
//   		},
//   		NumberOfRerankedResults: jsii.Number(123),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchrerankingconfiguration.html
//
type CfnFlowPropsMixin_VectorSearchRerankingConfigurationProperty struct {
	// Configuration for using Amazon Bedrock foundation models to rerank search results.
	//
	// This is required when the reranking type is set to BEDROCK.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchrerankingconfiguration-bedrockrerankingconfiguration
	//
	BedrockRerankingConfiguration interface{} `field:"optional" json:"bedrockRerankingConfiguration" yaml:"bedrockRerankingConfiguration"`
	// The type of reranking to apply to vector search results.
	//
	// Currently, the only supported value is BEDROCK, which uses Amazon Bedrock foundation models for reranking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchrerankingconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

