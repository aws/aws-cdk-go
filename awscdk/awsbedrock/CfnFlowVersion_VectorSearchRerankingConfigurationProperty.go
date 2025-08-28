package awsbedrock


// Configuration for reranking vector search results to improve relevance.
//
// Reranking applies additional relevance models to reorder the initial vector search results based on more sophisticated criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   vectorSearchRerankingConfigurationProperty := &VectorSearchRerankingConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	BedrockRerankingConfiguration: &VectorSearchBedrockRerankingConfigurationProperty{
//   		ModelConfiguration: &VectorSearchBedrockRerankingModelConfigurationProperty{
//   			ModelArn: jsii.String("modelArn"),
//
//   			// the properties below are optional
//   			AdditionalModelRequestFields: additionalModelRequestFields,
//   		},
//
//   		// the properties below are optional
//   		MetadataConfiguration: &MetadataConfigurationForRerankingProperty{
//   			SelectionMode: jsii.String("selectionMode"),
//
//   			// the properties below are optional
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
//   		NumberOfRerankedResults: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-vectorsearchrerankingconfiguration.html
//
type CfnFlowVersion_VectorSearchRerankingConfigurationProperty struct {
	// The type of reranking to apply to vector search results.
	//
	// Currently, the only supported value is BEDROCK, which uses Amazon Bedrock foundation models for reranking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-vectorsearchrerankingconfiguration.html#cfn-bedrock-flowversion-vectorsearchrerankingconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Configuration for using Amazon Bedrock foundation models to rerank search results.
	//
	// This is required when the reranking type is set to BEDROCK.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-vectorsearchrerankingconfiguration.html#cfn-bedrock-flowversion-vectorsearchrerankingconfiguration-bedrockrerankingconfiguration
	//
	BedrockRerankingConfiguration interface{} `field:"optional" json:"bedrockRerankingConfiguration" yaml:"bedrockRerankingConfiguration"`
}

