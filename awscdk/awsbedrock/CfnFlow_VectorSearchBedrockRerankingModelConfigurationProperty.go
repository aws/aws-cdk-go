package awsbedrock


// Configuration for the Amazon Bedrock foundation model used for reranking vector search results.
//
// This specifies which model to use and any additional parameters required by the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   vectorSearchBedrockRerankingModelConfigurationProperty := &VectorSearchBedrockRerankingModelConfigurationProperty{
//   	ModelArn: jsii.String("modelArn"),
//
//   	// the properties below are optional
//   	AdditionalModelRequestFields: additionalModelRequestFields,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingmodelconfiguration.html
//
type CfnFlow_VectorSearchBedrockRerankingModelConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the foundation model to use for reranking.
	//
	// This model processes the query and search results to determine a more relevant ordering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingmodelconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingmodelconfiguration-modelarn
	//
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// A list of additional fields to include in the model request during reranking.
	//
	// These fields provide extra context or configuration options specific to the selected foundation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-vectorsearchbedrockrerankingmodelconfiguration.html#cfn-bedrock-flow-vectorsearchbedrockrerankingmodelconfiguration-additionalmodelrequestfields
	//
	AdditionalModelRequestFields interface{} `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
}

