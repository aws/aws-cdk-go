package awsbedrock


// Contains configurations for a node in your flow.
//
// For more information, see [Node types in a flow](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var collector interface{}
//   var input interface{}
//   var iterator interface{}
//   var output interface{}
//
//   flowNodeConfigurationProperty := &FlowNodeConfigurationProperty{
//   	Agent: &AgentFlowNodeConfigurationProperty{
//   		AgentAliasArn: jsii.String("agentAliasArn"),
//   	},
//   	Collector: collector,
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
//   	InlineCode: &InlineCodeFlowNodeConfigurationProperty{
//   		Code: jsii.String("code"),
//   		Language: jsii.String("language"),
//   	},
//   	Input: input,
//   	Iterator: iterator,
//   	KnowledgeBase: &KnowledgeBaseFlowNodeConfigurationProperty{
//   		KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   		// the properties below are optional
//   		GuardrailConfiguration: &GuardrailConfigurationProperty{
//   			GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   			GuardrailVersion: jsii.String("guardrailVersion"),
//   		},
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
//   						TopP: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Resource: &PromptFlowNodeResourceConfigurationProperty{
//   				PromptArn: jsii.String("promptArn"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		GuardrailConfiguration: &GuardrailConfigurationProperty{
//   			GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   			GuardrailVersion: jsii.String("guardrailVersion"),
//   		},
//   	},
//   	Retrieval: &RetrievalFlowNodeConfigurationProperty{
//   		ServiceConfiguration: &RetrievalFlowNodeServiceConfigurationProperty{
//   			S3: &RetrievalFlowNodeS3ConfigurationProperty{
//   				BucketName: jsii.String("bucketName"),
//   			},
//   		},
//   	},
//   	Storage: &StorageFlowNodeConfigurationProperty{
//   		ServiceConfiguration: &StorageFlowNodeServiceConfigurationProperty{
//   			S3: &StorageFlowNodeS3ConfigurationProperty{
//   				BucketName: jsii.String("bucketName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html
//
type CfnFlow_FlowNodeConfigurationProperty struct {
	// Contains configurations for an agent node in your flow.
	//
	// Invokes an alias of an agent and returns the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-agent
	//
	Agent interface{} `field:"optional" json:"agent" yaml:"agent"`
	// Contains configurations for a collector node in your flow.
	//
	// Collects an iteration of inputs and consolidates them into an array of outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-collector
	//
	Collector interface{} `field:"optional" json:"collector" yaml:"collector"`
	// Contains configurations for a condition node in your flow.
	//
	// Defines conditions that lead to different branches of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-condition
	//
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Contains configurations for an inline code node in your flow.
	//
	// Inline code nodes let you write and execute code directly within your flow, enabling data transformations, custom logic, and integrations without needing an external Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-inlinecode
	//
	InlineCode interface{} `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// Contains configurations for an input flow node in your flow.
	//
	// The first node in the flow. `inputs` can't be specified for this node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-input
	//
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// Contains configurations for an iterator node in your flow.
	//
	// Takes an input that is an array and iteratively sends each item of the array as an output to the following node. The size of the array is also returned in the output.
	//
	// The output flow node at the end of the flow iteration will return a response for each member of the array. To return only one response, you can include a collector node downstream from the iterator node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-iterator
	//
	Iterator interface{} `field:"optional" json:"iterator" yaml:"iterator"`
	// Contains configurations for a knowledge base node in your flow.
	//
	// Queries a knowledge base and returns the retrieved results or generated response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-knowledgebase
	//
	KnowledgeBase interface{} `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
	// Contains configurations for a Lambda function node in your flow.
	//
	// Invokes an AWS Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-lambdafunction
	//
	LambdaFunction interface{} `field:"optional" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Contains configurations for a Lex node in your flow.
	//
	// Invokes an Amazon Lex bot to identify the intent of the input and return the intent as the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-lex
	//
	Lex interface{} `field:"optional" json:"lex" yaml:"lex"`
	// Contains configurations for an output flow node in your flow.
	//
	// The last node in the flow. `outputs` can't be specified for this node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// Contains configurations for a prompt node in your flow.
	//
	// Runs a prompt and generates the model response as the output. You can use a prompt from Prompt management or you can configure one in this node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-prompt
	//
	Prompt interface{} `field:"optional" json:"prompt" yaml:"prompt"`
	// Contains configurations for a retrieval node in your flow.
	//
	// Retrieves data from an Amazon S3 location and returns it as the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-retrieval
	//
	Retrieval interface{} `field:"optional" json:"retrieval" yaml:"retrieval"`
	// Contains configurations for a storage node in your flow.
	//
	// Stores an input in an Amazon S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeconfiguration.html#cfn-bedrock-flow-flownodeconfiguration-storage
	//
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
}

