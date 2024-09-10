package awsbedrock


// The definition of the nodes and connections between nodes in the flow.
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
//   flowDefinitionProperty := &FlowDefinitionProperty{
//   	Connections: []interface{}{
//   		&FlowConnectionProperty{
//   			Name: jsii.String("name"),
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Configuration: &FlowConnectionConfigurationProperty{
//   				Conditional: &FlowConditionalConnectionConfigurationProperty{
//   					Condition: jsii.String("condition"),
//   				},
//   				Data: &FlowDataConnectionConfigurationProperty{
//   					SourceOutput: jsii.String("sourceOutput"),
//   					TargetInput: jsii.String("targetInput"),
//   				},
//   			},
//   		},
//   	},
//   	Nodes: []interface{}{
//   		&FlowNodeProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Configuration: &FlowNodeConfigurationProperty{
//   				Agent: &AgentFlowNodeConfigurationProperty{
//   					AgentAliasArn: jsii.String("agentAliasArn"),
//   				},
//   				Collector: collector,
//   				Condition: &ConditionFlowNodeConfigurationProperty{
//   					Conditions: []interface{}{
//   						&FlowConditionProperty{
//   							Name: jsii.String("name"),
//
//   							// the properties below are optional
//   							Expression: jsii.String("expression"),
//   						},
//   					},
//   				},
//   				Input: input,
//   				Iterator: iterator,
//   				KnowledgeBase: &KnowledgeBaseFlowNodeConfigurationProperty{
//   					KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   					// the properties below are optional
//   					ModelId: jsii.String("modelId"),
//   				},
//   				LambdaFunction: &LambdaFunctionFlowNodeConfigurationProperty{
//   					LambdaArn: jsii.String("lambdaArn"),
//   				},
//   				Lex: &LexFlowNodeConfigurationProperty{
//   					BotAliasArn: jsii.String("botAliasArn"),
//   					LocaleId: jsii.String("localeId"),
//   				},
//   				Output: output,
//   				Prompt: &PromptFlowNodeConfigurationProperty{
//   					SourceConfiguration: &PromptFlowNodeSourceConfigurationProperty{
//   						Inline: &PromptFlowNodeInlineConfigurationProperty{
//   							ModelId: jsii.String("modelId"),
//   							TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   								Text: &TextPromptTemplateConfigurationProperty{
//   									Text: jsii.String("text"),
//
//   									// the properties below are optional
//   									InputVariables: []interface{}{
//   										&PromptInputVariableProperty{
//   											Name: jsii.String("name"),
//   										},
//   									},
//   								},
//   							},
//   							TemplateType: jsii.String("templateType"),
//
//   							// the properties below are optional
//   							InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   								Text: &PromptModelInferenceConfigurationProperty{
//   									MaxTokens: jsii.Number(123),
//   									StopSequences: []*string{
//   										jsii.String("stopSequences"),
//   									},
//   									Temperature: jsii.Number(123),
//   									TopK: jsii.Number(123),
//   									TopP: jsii.Number(123),
//   								},
//   							},
//   						},
//   						Resource: &PromptFlowNodeResourceConfigurationProperty{
//   							PromptArn: jsii.String("promptArn"),
//   						},
//   					},
//   				},
//   				Retrieval: &RetrievalFlowNodeConfigurationProperty{
//   					ServiceConfiguration: &RetrievalFlowNodeServiceConfigurationProperty{
//   						S3: &RetrievalFlowNodeS3ConfigurationProperty{
//   							BucketName: jsii.String("bucketName"),
//   						},
//   					},
//   				},
//   				Storage: &StorageFlowNodeConfigurationProperty{
//   					ServiceConfiguration: &StorageFlowNodeServiceConfigurationProperty{
//   						S3: &StorageFlowNodeS3ConfigurationProperty{
//   							BucketName: jsii.String("bucketName"),
//   						},
//   					},
//   				},
//   			},
//   			Inputs: []interface{}{
//   				&FlowNodeInputProperty{
//   					Expression: jsii.String("expression"),
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Outputs: []interface{}{
//   				&FlowNodeOutputProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html
//
type CfnFlow_FlowDefinitionProperty struct {
	// An array of connection definitions in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html#cfn-bedrock-flow-flowdefinition-connections
	//
	Connections interface{} `field:"optional" json:"connections" yaml:"connections"`
	// An array of node definitions in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html#cfn-bedrock-flow-flowdefinition-nodes
	//
	Nodes interface{} `field:"optional" json:"nodes" yaml:"nodes"`
}

