package awsbedrock


// Strategy to be used when using Bedrock Foundation Model for Context Enrichment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enrichmentStrategyConfigurationProperty := &EnrichmentStrategyConfigurationProperty{
//   	Method: jsii.String("method"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-enrichmentstrategyconfiguration.html
//
type CfnDataSource_EnrichmentStrategyConfigurationProperty struct {
	// Enrichment Strategy method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-enrichmentstrategyconfiguration.html#cfn-bedrock-datasource-enrichmentstrategyconfiguration-method
	//
	Method *string `field:"required" json:"method" yaml:"method"`
}

