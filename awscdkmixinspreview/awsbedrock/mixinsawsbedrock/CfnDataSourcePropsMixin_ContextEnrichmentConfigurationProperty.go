package mixinsawsbedrock


// Context enrichment configuration is used to provide additional context to the RAG application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contextEnrichmentConfigurationProperty := &ContextEnrichmentConfigurationProperty{
//   	BedrockFoundationModelConfiguration: &BedrockFoundationModelContextEnrichmentConfigurationProperty{
//   		EnrichmentStrategyConfiguration: &EnrichmentStrategyConfigurationProperty{
//   			Method: jsii.String("method"),
//   		},
//   		ModelArn: jsii.String("modelArn"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html
//
type CfnDataSourcePropsMixin_ContextEnrichmentConfigurationProperty struct {
	// The configuration of the Amazon Bedrock foundation model used for context enrichment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html#cfn-bedrock-datasource-contextenrichmentconfiguration-bedrockfoundationmodelconfiguration
	//
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
	// The method used for context enrichment.
	//
	// It must be Amazon Bedrock foundation models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html#cfn-bedrock-datasource-contextenrichmentconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

