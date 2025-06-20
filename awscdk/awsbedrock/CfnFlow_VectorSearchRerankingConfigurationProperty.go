package awsbedrock


// Contains configurations for reranking the retrieved results.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchrerankingconfiguration.html
//
type CfnFlow_VectorSearchRerankingConfigurationProperty struct {
	// The type of reranker model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchrerankingconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Contains configurations for an Amazon Bedrock reranker model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchrerankingconfiguration-bedrockrerankingconfiguration
	//
	BedrockRerankingConfiguration interface{} `field:"optional" json:"bedrockRerankingConfiguration" yaml:"bedrockRerankingConfiguration"`
}

