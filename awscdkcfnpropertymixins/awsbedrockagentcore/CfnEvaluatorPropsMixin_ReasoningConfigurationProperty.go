package awsbedrockagentcore


// The reasoning configuration for reasoning models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   reasoningConfigurationProperty := &ReasoningConfigurationProperty{
//   	Effort: jsii.String("effort"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-reasoningconfiguration.html
//
type CfnEvaluatorPropsMixin_ReasoningConfigurationProperty struct {
	// The level of reasoning effort the model applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-reasoningconfiguration.html#cfn-bedrockagentcore-evaluator-reasoningconfiguration-effort
	//
	Effort *string `field:"optional" json:"effort" yaml:"effort"`
}

