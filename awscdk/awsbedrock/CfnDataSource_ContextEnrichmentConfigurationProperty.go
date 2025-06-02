package awsbedrock


// Context enrichment configuration is used to provide additional context to the RAG application.
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
	// The method used for context enrichment.
	//
	// It must be Amazon Bedrock foundation models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html#cfn-bedrock-datasource-contextenrichmentconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration of the Amazon Bedrock foundation model used for context enrichment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html#cfn-bedrock-datasource-contextenrichmentconfiguration-bedrockfoundationmodelconfiguration
	//
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

