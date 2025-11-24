package mixinsawsbedrock


// Configures how the knowledge base orchestrates the retrieval and generation process, allowing for customization of prompts, inference parameters, and performance settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//
//   knowledgeBaseOrchestrationConfigurationProperty := &KnowledgeBaseOrchestrationConfigurationProperty{
//   	AdditionalModelRequestFields: additionalModelRequestFields,
//   	InferenceConfig: &PromptInferenceConfigurationProperty{
//   		Text: &PromptModelInferenceConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			StopSequences: []*string{
//   				jsii.String("stopSequences"),
//   			},
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//   	PerformanceConfig: &PerformanceConfigurationProperty{
//   		Latency: jsii.String("latency"),
//   	},
//   	PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   		TextPromptTemplate: jsii.String("textPromptTemplate"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-knowledgebaseorchestrationconfiguration.html
//
type CfnFlowPropsMixin_KnowledgeBaseOrchestrationConfigurationProperty struct {
	// The additional model-specific request parameters as key-value pairs to be included in the request to the foundation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-knowledgebaseorchestrationconfiguration.html#cfn-bedrock-flow-knowledgebaseorchestrationconfiguration-additionalmodelrequestfields
	//
	AdditionalModelRequestFields interface{} `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// Contains inference configurations for the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-knowledgebaseorchestrationconfiguration.html#cfn-bedrock-flow-knowledgebaseorchestrationconfiguration-inferenceconfig
	//
	InferenceConfig interface{} `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
	// The performance configuration options for the knowledge base retrieval and generation process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-knowledgebaseorchestrationconfiguration.html#cfn-bedrock-flow-knowledgebaseorchestrationconfiguration-performanceconfig
	//
	PerformanceConfig interface{} `field:"optional" json:"performanceConfig" yaml:"performanceConfig"`
	// A custom prompt template for orchestrating the retrieval and generation process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-knowledgebaseorchestrationconfiguration.html#cfn-bedrock-flow-knowledgebaseorchestrationconfiguration-prompttemplate
	//
	PromptTemplate interface{} `field:"optional" json:"promptTemplate" yaml:"promptTemplate"`
}

