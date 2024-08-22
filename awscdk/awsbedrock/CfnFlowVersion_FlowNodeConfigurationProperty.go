package awsbedrock


// Contains configurations for a node in your flow.
//
// For more information, see [Node types in Amazon Bedrock works](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var input interface{}
//   var output interface{}
//
//   flowNodeConfigurationProperty := &FlowNodeConfigurationProperty{
//   	Condition: &ConditionFlowNodeConfigurationProperty{
//   		Conditions: []interface{}{
//   			&FlowConditionProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Expression: jsii.String("expression"),
//   			},
//   		},
//   	},
//   	Input: input,
//   	KnowledgeBase: &KnowledgeBaseFlowNodeConfigurationProperty{
//   		KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   		// the properties below are optional
//   		ModelId: jsii.String("modelId"),
//   	},
//   	LambdaFunction: &LambdaFunctionFlowNodeConfigurationProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   	},
//   	Lex: &LexFlowNodeConfigurationProperty{
//   		BotAliasArn: jsii.String("botAliasArn"),
//   		LocaleId: jsii.String("localeId"),
//   	},
//   	Output: output,
//   	Prompt: &PromptFlowNodeConfigurationProperty{
//   		SourceConfiguration: &PromptFlowNodeSourceConfigurationProperty{
//   			Inline: &PromptFlowNodeInlineConfigurationProperty{
//   				ModelId: jsii.String("modelId"),
//   				TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   					Text: &TextPromptTemplateConfigurationProperty{
//   						Text: jsii.String("text"),
//
//   						// the properties below are optional
//   						InputVariables: []interface{}{
//   							&PromptInputVariableProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   					},
//   				},
//   				TemplateType: jsii.String("templateType"),
//
//   				// the properties below are optional
//   				InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   					Text: &PromptModelInferenceConfigurationProperty{
//   						MaxTokens: jsii.Number(123),
//   						StopSequences: []*string{
//   							jsii.String("stopSequences"),
//   						},
//   						Temperature: jsii.Number(123),
//   						TopK: jsii.Number(123),
//   						TopP: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Resource: &PromptFlowNodeResourceConfigurationProperty{
//   				PromptArn: jsii.String("promptArn"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html
//
type CfnFlowVersion_FlowNodeConfigurationProperty struct {
	// Contains configurations for a Condition node in your flow.
	//
	// Defines conditions that lead to different branches of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html#cfn-bedrock-flowversion-flownodeconfiguration-condition
	//
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Contains configurations for an input flow node in your flow.
	//
	// The first node in the flow. `inputs` can't be specified for this node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html#cfn-bedrock-flowversion-flownodeconfiguration-input
	//
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// Contains configurations for a knowledge base node in your flow.
	//
	// Queries a knowledge base and returns the retrieved results or generated response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html#cfn-bedrock-flowversion-flownodeconfiguration-knowledgebase
	//
	KnowledgeBase interface{} `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
	// Contains configurations for a Lambda function node in your flow.
	//
	// Invokes an AWS Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html#cfn-bedrock-flowversion-flownodeconfiguration-lambdafunction
	//
	LambdaFunction interface{} `field:"optional" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Contains configurations for a Lex node in your flow.
	//
	// Invokes an Amazon Lex bot to identify the intent of the input and return the intent as the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html#cfn-bedrock-flowversion-flownodeconfiguration-lex
	//
	Lex interface{} `field:"optional" json:"lex" yaml:"lex"`
	// Contains configurations for an output flow node in your flow.
	//
	// The last node in the flow. `outputs` can't be specified for this node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html#cfn-bedrock-flowversion-flownodeconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// Contains configurations for a prompt node in your flow.
	//
	// Runs a prompt and generates the model response as the output. You can use a prompt from Prompt management or you can configure one in this node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeconfiguration.html#cfn-bedrock-flowversion-flownodeconfiguration-prompt
	//
	Prompt interface{} `field:"optional" json:"prompt" yaml:"prompt"`
}

