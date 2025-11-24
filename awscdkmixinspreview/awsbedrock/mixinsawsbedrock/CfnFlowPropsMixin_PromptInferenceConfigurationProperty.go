package mixinsawsbedrock


// Contains inference configurations for the prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   promptInferenceConfigurationProperty := &PromptInferenceConfigurationProperty{
//   	Text: &PromptModelInferenceConfigurationProperty{
//   		MaxTokens: jsii.Number(123),
//   		StopSequences: []*string{
//   			jsii.String("stopSequences"),
//   		},
//   		Temperature: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinferenceconfiguration.html
//
type CfnFlowPropsMixin_PromptInferenceConfigurationProperty struct {
	// Contains inference configurations for a text prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinferenceconfiguration.html#cfn-bedrock-flow-promptinferenceconfiguration-text
	//
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

