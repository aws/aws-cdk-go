package mixinsawsbedrock


// The definition of the nodes and connections between nodes in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//   var collector interface{}
//   var flowDefinitionProperty_ FlowDefinitionProperty
//   var input interface{}
//   var iterator interface{}
//   var loopInput interface{}
//   var output interface{}
//
//   flowDefinitionProperty := &FlowDefinitionProperty{
//   	Connections: []interface{}{
//   		&FlowConnectionProperty{
//   			Configuration: &FlowConnectionConfigurationProperty{
//   				Conditional: &FlowConditionalConnectionConfigurationProperty{
//   					Condition: jsii.String("condition"),
//   				},
//   				Data: &FlowDataConnectionConfigurationProperty{
//   					SourceOutput: jsii.String("sourceOutput"),
//   					TargetInput: jsii.String("targetInput"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Nodes: []interface{}{
//   		&FlowNodeProperty{
//   			Configuration: &FlowNodeConfigurationProperty{
//   				Agent: &AgentFlowNodeConfigurationProperty{
//   					AgentAliasArn: jsii.String("agentAliasArn"),
//   				},
//   				Collector: collector,
//   				Condition: &ConditionFlowNodeConfigurationProperty{
//   					Conditions: []interface{}{
//   						&FlowConditionProperty{
//   							Expression: jsii.String("expression"),
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				InlineCode: &InlineCodeFlowNodeConfigurationProperty{
//   					Code: jsii.String("code"),
//   					Language: jsii.String("language"),
//   				},
//   				Input: input,
//   				Iterator: iterator,
//   				KnowledgeBase: &KnowledgeBaseFlowNodeConfigurationProperty{
//   					GuardrailConfiguration: &GuardrailConfigurationProperty{
//   						GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   						GuardrailVersion: jsii.String("guardrailVersion"),
//   					},
//   					InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   						Text: &PromptModelInferenceConfigurationProperty{
//   							MaxTokens: jsii.Number(123),
//   							StopSequences: []*string{
//   								jsii.String("stopSequences"),
//   							},
//   							Temperature: jsii.Number(123),
//   							TopP: jsii.Number(123),
//   						},
//   					},
//   					KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   					ModelId: jsii.String("modelId"),
//   					NumberOfResults: jsii.Number(123),
//   					OrchestrationConfiguration: &KnowledgeBaseOrchestrationConfigurationProperty{
//   						AdditionalModelRequestFields: additionalModelRequestFields,
//   						InferenceConfig: &PromptInferenceConfigurationProperty{
//   							Text: &PromptModelInferenceConfigurationProperty{
//   								MaxTokens: jsii.Number(123),
//   								StopSequences: []*string{
//   									jsii.String("stopSequences"),
//   								},
//   								Temperature: jsii.Number(123),
//   								TopP: jsii.Number(123),
//   							},
//   						},
//   						PerformanceConfig: &PerformanceConfigurationProperty{
//   							Latency: jsii.String("latency"),
//   						},
//   						PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   							TextPromptTemplate: jsii.String("textPromptTemplate"),
//   						},
//   					},
//   					PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   						TextPromptTemplate: jsii.String("textPromptTemplate"),
//   					},
//   					RerankingConfiguration: &VectorSearchRerankingConfigurationProperty{
//   						BedrockRerankingConfiguration: &VectorSearchBedrockRerankingConfigurationProperty{
//   							MetadataConfiguration: &MetadataConfigurationForRerankingProperty{
//   								SelectionMode: jsii.String("selectionMode"),
//   								SelectiveModeConfiguration: &RerankingMetadataSelectiveModeConfigurationProperty{
//   									FieldsToExclude: []interface{}{
//   										&FieldForRerankingProperty{
//   											FieldName: jsii.String("fieldName"),
//   										},
//   									},
//   									FieldsToInclude: []interface{}{
//   										&FieldForRerankingProperty{
//   											FieldName: jsii.String("fieldName"),
//   										},
//   									},
//   								},
//   							},
//   							ModelConfiguration: &VectorSearchBedrockRerankingModelConfigurationProperty{
//   								AdditionalModelRequestFields: additionalModelRequestFields,
//   								ModelArn: jsii.String("modelArn"),
//   							},
//   							NumberOfRerankedResults: jsii.Number(123),
//   						},
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				LambdaFunction: &LambdaFunctionFlowNodeConfigurationProperty{
//   					LambdaArn: jsii.String("lambdaArn"),
//   				},
//   				Lex: &LexFlowNodeConfigurationProperty{
//   					BotAliasArn: jsii.String("botAliasArn"),
//   					LocaleId: jsii.String("localeId"),
//   				},
//   				Loop: &LoopFlowNodeConfigurationProperty{
//   					Definition: flowDefinitionProperty_,
//   				},
//   				LoopController: &LoopControllerFlowNodeConfigurationProperty{
//   					ContinueCondition: &FlowConditionProperty{
//   						Expression: jsii.String("expression"),
//   						Name: jsii.String("name"),
//   					},
//   					MaxIterations: jsii.Number(123),
//   				},
//   				LoopInput: loopInput,
//   				Output: output,
//   				Prompt: &PromptFlowNodeConfigurationProperty{
//   					GuardrailConfiguration: &GuardrailConfigurationProperty{
//   						GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   						GuardrailVersion: jsii.String("guardrailVersion"),
//   					},
//   					SourceConfiguration: &PromptFlowNodeSourceConfigurationProperty{
//   						Inline: &PromptFlowNodeInlineConfigurationProperty{
//   							InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   								Text: &PromptModelInferenceConfigurationProperty{
//   									MaxTokens: jsii.Number(123),
//   									StopSequences: []*string{
//   										jsii.String("stopSequences"),
//   									},
//   									Temperature: jsii.Number(123),
//   									TopP: jsii.Number(123),
//   								},
//   							},
//   							ModelId: jsii.String("modelId"),
//   							TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   								Text: &TextPromptTemplateConfigurationProperty{
//   									InputVariables: []interface{}{
//   										&PromptInputVariableProperty{
//   											Name: jsii.String("name"),
//   										},
//   									},
//   									Text: jsii.String("text"),
//   								},
//   							},
//   							TemplateType: jsii.String("templateType"),
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
//   			Name: jsii.String("name"),
//   			Outputs: []interface{}{
//   				&FlowNodeOutputProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowdefinition.html
//
type CfnFlowVersionPropsMixin_FlowDefinitionProperty struct {
	// An array of connection definitions in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowdefinition.html#cfn-bedrock-flowversion-flowdefinition-connections
	//
	Connections interface{} `field:"optional" json:"connections" yaml:"connections"`
	// An array of node definitions in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flowdefinition.html#cfn-bedrock-flowversion-flowdefinition-nodes
	//
	Nodes interface{} `field:"optional" json:"nodes" yaml:"nodes"`
}

