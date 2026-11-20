package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   aIPromptInferenceConfigurationProperty := &AIPromptInferenceConfigurationProperty{
//   	MaxTokensToSample: jsii.Number(123),
//   	Temperature: jsii.Number(123),
//   	TopK: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-aipromptinferenceconfiguration.html
//
type CfnAIPromptPropsMixin_AIPromptInferenceConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-aipromptinferenceconfiguration.html#cfn-wisdom-aiprompt-aipromptinferenceconfiguration-maxtokenstosample
	//
	MaxTokensToSample *float64 `field:"optional" json:"maxTokensToSample" yaml:"maxTokensToSample"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-aipromptinferenceconfiguration.html#cfn-wisdom-aiprompt-aipromptinferenceconfiguration-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-aipromptinferenceconfiguration.html#cfn-wisdom-aiprompt-aipromptinferenceconfiguration-topk
	//
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-aipromptinferenceconfiguration.html#cfn-wisdom-aiprompt-aipromptinferenceconfiguration-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

