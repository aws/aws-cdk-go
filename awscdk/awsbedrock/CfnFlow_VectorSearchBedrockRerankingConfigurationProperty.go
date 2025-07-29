package awsbedrock


// Configuration for using Amazon Bedrock foundation models to rerank Knowledge Base vector search results.
//
// This enables more sophisticated relevance ranking using large language models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   vectorSearchBedrockRerankingConfigurationProperty := &VectorSearchBedrockRerankingConfigurationProperty{
//   	ModelConfiguration: &VectorSearchBedrockRerankingModelConfigurationProperty{
//   		ModelArn: jsii.String("modelArn"),
//
//   		// the properties below are optional
//   		AdditionalModelRequestFields: additionalModelRequestFields,
//   	},
//
//   	// the properties below are optional
//   	MetadataConfiguration: &MetadataConfigurationForRerankingProperty{
//   		SelectionMode: jsii.String("selectionMode"),
//
//   		// the properties below are optional
//   		SelectiveModeConfiguration: &RerankingMetadataSelectiveModeConfigurationProperty{
//   			FieldsToExclude: []interface{}{
//   				&FieldForRerankingProperty{
//   					FieldName: jsii.String("fieldName"),
//   				},
//   			},
//   			FieldsToInclude: []interface{}{
//   				&FieldForRerankingProperty{
//   					FieldName: jsii.String("fieldName"),
//   				},
//   			},
//   		},
//   	},
//   	NumberOfRerankedResults: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.html
//
type CfnFlow_VectorSearchBedrockRerankingConfigurationProperty struct {
	// Configuration for the Amazon Bedrock foundation model used for reranking.
	//
	// This includes the model ARN and any additional request fields required by the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingconfiguration-modelconfiguration
	//
	ModelConfiguration interface{} `field:"required" json:"modelConfiguration" yaml:"modelConfiguration"`
	// Configuration for how document metadata should be used during the reranking process.
	//
	// This determines which metadata fields are included when reordering search results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingconfiguration-metadataconfiguration
	//
	MetadataConfiguration interface{} `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// The maximum number of results to rerank.
	//
	// This limits how many of the initial vector search results will be processed by the reranking model. A smaller number improves performance but may exclude potentially relevant results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingconfiguration-numberofrerankedresults
	//
	NumberOfRerankedResults *float64 `field:"optional" json:"numberOfRerankedResults" yaml:"numberOfRerankedResults"`
}

