package awsbedrock


// Contains configurations for reranking with an Amazon Bedrock reranker model.
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
	// Contains configurations for the reranker model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingconfiguration-modelconfiguration
	//
	ModelConfiguration interface{} `field:"required" json:"modelConfiguration" yaml:"modelConfiguration"`
	// Contains configurations for the metadata to use in reranking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingconfiguration-metadataconfiguration
	//
	MetadataConfiguration interface{} `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// The number of results to return after reranking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingconfiguration-numberofrerankedresults
	//
	NumberOfRerankedResults *float64 `field:"optional" json:"numberOfRerankedResults" yaml:"numberOfRerankedResults"`
}

