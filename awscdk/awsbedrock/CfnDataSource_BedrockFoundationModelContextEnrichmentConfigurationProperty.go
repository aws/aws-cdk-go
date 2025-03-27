package awsbedrock


// Bedrock Foundation Model configuration to be used for Context Enrichment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnDataSource_BedrockFoundationModelContextEnrichmentConfigurationProperty struct {
	// Strategy to be used when using Bedrock Foundation Model for Context Enrichment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration-enrichmentstrategyconfiguration
	//
	EnrichmentStrategyConfiguration interface{} `field:"required" json:"enrichmentStrategyConfiguration" yaml:"enrichmentStrategyConfiguration"`
	// The model's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration.html#cfn-bedrock-datasource-bedrockfoundationmodelcontextenrichmentconfiguration-modelarn
	//
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
}

