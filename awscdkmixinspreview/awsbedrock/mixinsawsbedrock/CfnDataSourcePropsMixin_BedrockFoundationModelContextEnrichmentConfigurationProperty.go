package mixinsawsbedrock


// Context enrichment configuration is used to provide additional context to the RAG application using Amazon Bedrock foundation models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bedrockFoundationModelContextEnrichmentConfigurationProperty := &BedrockFoundationModelContextEnrichmentConfigurationProperty{
//   	EnrichmentStrategyConfiguration: &EnrichmentStrategyConfigurationProperty{
//   		Method: jsii.String("method"),
//   	},
//   	ModelArn: jsii.String("modelArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration.html
//
type CfnDataSourcePropsMixin_BedrockFoundationModelContextEnrichmentConfigurationProperty struct {
	// The enrichment stategy used to provide additional context.
	//
	// For example, Neptune GraphRAG uses Amazon Bedrock foundation models to perform chunk entity extraction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration-enrichmentstrategyconfiguration
	//
	EnrichmentStrategyConfiguration interface{} `field:"optional" json:"enrichmentStrategyConfiguration" yaml:"enrichmentStrategyConfiguration"`
	// The Amazon Resource Name (ARN) of the model used to create vector embeddings for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration-modelarn
	//
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
}

