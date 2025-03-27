package awsbedrock


// Additional Enrichment Configuration for example when using GraphRag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextEnrichmentConfigurationProperty := &ContextEnrichmentConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	BedrockFoundationModelConfiguration: &BedrockFoundationModelContextEnrichmentConfigurationProperty{
//   		EnrichmentStrategyConfiguration: &EnrichmentStrategyConfigurationProperty{
//   			Method: jsii.String("method"),
//   		},
//   		ModelArn: jsii.String("modelArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html
//
type CfnDataSource_ContextEnrichmentConfigurationProperty struct {
	// Enrichment type to be used for the vector database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html#cfn-bedrock-datasource-contextenrichmentconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Bedrock Foundation Model configuration to be used for Context Enrichment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html#cfn-bedrock-datasource-contextenrichmentconfiguration-bedrockfoundationmodelconfiguration
	//
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

