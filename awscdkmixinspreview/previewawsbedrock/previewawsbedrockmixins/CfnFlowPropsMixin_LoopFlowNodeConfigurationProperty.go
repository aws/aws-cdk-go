package previewawsbedrockmixins


// Contains configurations for the nodes of a DoWhile loop in your flow.
//
// A DoWhile loop is made up of the following nodes:
//
// - `Loop` - The container node that holds the loop's flow definition. This node encompasses the entire loop structure.
// - `LoopInput` - The entry point node for the loop. This node receives inputs from nodes outside the loop and from previous loop iterations.
// - Body nodes - The processing nodes that execute within each loop iteration. These can be nodes for handling data in your flow, such as a prompt or Lambda function nodes. Some node types aren't supported inside a DoWhile loop body. For more information, see [LoopIncompatibleNodeTypeFlowValidationDetails](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_LoopIncompatibleNodeTypeFlowValidationDetails.html) .
// - `LoopController` - The node that evaluates whether the loop should continue or exit based on a condition.
//
// These nodes work together to create a loop that runs at least once and continues until a specified condition is met or a maximum number of iterations is reached.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//   var collector interface{}
//   var input interface{}
//   var iterator interface{}
//   var loopFlowNodeConfigurationProperty_ LoopFlowNodeConfigurationProperty
//   var loopInput interface{}
//   var output interface{}
//
//   loopFlowNodeConfigurationProperty := &LoopFlowNodeConfigurationProperty{
//   	Definition: &FlowDefinitionProperty{
//   		Connections: []interface{}{
//   			&FlowConnectionProperty{
//   				Configuration: &FlowConnectionConfigurationProperty{
//   					Conditional: &FlowConditionalConnectionConfigurationProperty{
//   						Condition: jsii.String("condition"),
//   					},
//   					Data: &FlowDataConnectionConfigurationProperty{
//   						SourceOutput: jsii.String("sourceOutput"),
//   						TargetInput: jsii.String("targetInput"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Nodes: []interface{}{
//   			&FlowNodeProperty{
//   				Configuration: &FlowNodeConfigurationProperty{
//   					Agent: &AgentFlowNodeConfigurationProperty{
//   						AgentAliasArn: jsii.String("agentAliasArn"),
//   					},
//   					Collector: collector,
//   					Condition: &ConditionFlowNodeConfigurationProperty{
//   						Conditions: []interface{}{
//   							&FlowConditionProperty{
//   								Expression: jsii.String("expression"),
//   								Name: jsii.String("name"),
//   							},
//   						},
//   					},
//   					InlineCode: &InlineCodeFlowNodeConfigurationProperty{
//   						Code: jsii.String("code"),
//   						Language: jsii.String("language"),
//   					},
//   					Input: input,
//   					Iterator: iterator,
//   					KnowledgeBase: &KnowledgeBaseFlowNodeConfigurationProperty{
//   						GuardrailConfiguration: &GuardrailConfigurationProperty{
//   							GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   							GuardrailVersion: jsii.String("guardrailVersion"),
//   						},
//   						InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   							Text: &PromptModelInferenceConfigurationProperty{
//   								MaxTokens: jsii.Number(123),
//   								StopSequences: []*string{
//   									jsii.String("stopSequences"),
//   								},
//   								Temperature: jsii.Number(123),
//   								TopP: jsii.Number(123),
//   							},
//   						},
//   						KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   						ModelId: jsii.String("modelId"),
//   						NumberOfResults: jsii.Number(123),
//   						OrchestrationConfiguration: &KnowledgeBaseOrchestrationConfigurationProperty{
//   							AdditionalModelRequestFields: additionalModelRequestFields,
//   							InferenceConfig: &PromptInferenceConfigurationProperty{
//   								Text: &PromptModelInferenceConfigurationProperty{
//   									MaxTokens: jsii.Number(123),
//   									StopSequences: []*string{
//   										jsii.String("stopSequences"),
//   									},
//   									Temperature: jsii.Number(123),
//   									TopP: jsii.Number(123),
//   								},
//   							},
//   							PerformanceConfig: &PerformanceConfigurationProperty{
//   								Latency: jsii.String("latency"),
//   							},
//   							PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   								TextPromptTemplate: jsii.String("textPromptTemplate"),
//   							},
//   						},
//   						PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   							TextPromptTemplate: jsii.String("textPromptTemplate"),
//   						},
//   						RerankingConfiguration: &VectorSearchRerankingConfigurationProperty{
//   							BedrockRerankingConfiguration: &VectorSearchBedrockRerankingConfigurationProperty{
//   								MetadataConfiguration: &MetadataConfigurationForRerankingProperty{
//   									SelectionMode: jsii.String("selectionMode"),
//   									SelectiveModeConfiguration: &RerankingMetadataSelectiveModeConfigurationProperty{
//   										FieldsToExclude: []interface{}{
//   											&FieldForRerankingProperty{
//   												FieldName: jsii.String("fieldName"),
//   											},
//   										},
//   										FieldsToInclude: []interface{}{
//   											&FieldForRerankingProperty{
//   												FieldName: jsii.String("fieldName"),
//   											},
//   										},
//   									},
//   								},
//   								ModelConfiguration: &VectorSearchBedrockRerankingModelConfigurationProperty{
//   									AdditionalModelRequestFields: additionalModelRequestFields,
//   									ModelArn: jsii.String("modelArn"),
//   								},
//   								NumberOfRerankedResults: jsii.Number(123),
//   							},
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					LambdaFunction: &LambdaFunctionFlowNodeConfigurationProperty{
//   						LambdaArn: jsii.String("lambdaArn"),
//   					},
//   					Lex: &LexFlowNodeConfigurationProperty{
//   						BotAliasArn: jsii.String("botAliasArn"),
//   						LocaleId: jsii.String("localeId"),
//   					},
//   					Loop: loopFlowNodeConfigurationProperty_,
//   					LoopController: &LoopControllerFlowNodeConfigurationProperty{
//   						ContinueCondition: &FlowConditionProperty{
//   							Expression: jsii.String("expression"),
//   							Name: jsii.String("name"),
//   						},
//   						MaxIterations: jsii.Number(123),
//   					},
//   					LoopInput: loopInput,
//   					Output: output,
//   					Prompt: &PromptFlowNodeConfigurationProperty{
//   						GuardrailConfiguration: &GuardrailConfigurationProperty{
//   							GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   							GuardrailVersion: jsii.String("guardrailVersion"),
//   						},
//   						SourceConfiguration: &PromptFlowNodeSourceConfigurationProperty{
//   							Inline: &PromptFlowNodeInlineConfigurationProperty{
//   								InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   									Text: &PromptModelInferenceConfigurationProperty{
//   										MaxTokens: jsii.Number(123),
//   										StopSequences: []*string{
//   											jsii.String("stopSequences"),
//   										},
//   										Temperature: jsii.Number(123),
//   										TopP: jsii.Number(123),
//   									},
//   								},
//   								ModelId: jsii.String("modelId"),
//   								TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   									Text: &TextPromptTemplateConfigurationProperty{
//   										InputVariables: []interface{}{
//   											&PromptInputVariableProperty{
//   												Name: jsii.String("name"),
//   											},
//   										},
//   										Text: jsii.String("text"),
//   									},
//   								},
//   								TemplateType: jsii.String("templateType"),
//   							},
//   							Resource: &PromptFlowNodeResourceConfigurationProperty{
//   								PromptArn: jsii.String("promptArn"),
//   							},
//   						},
//   					},
//   					Retrieval: &RetrievalFlowNodeConfigurationProperty{
//   						ServiceConfiguration: &RetrievalFlowNodeServiceConfigurationProperty{
//   							S3: &RetrievalFlowNodeS3ConfigurationProperty{
//   								BucketName: jsii.String("bucketName"),
//   							},
//   						},
//   					},
//   					Storage: &StorageFlowNodeConfigurationProperty{
//   						ServiceConfiguration: &StorageFlowNodeServiceConfigurationProperty{
//   							S3: &StorageFlowNodeS3ConfigurationProperty{
//   								BucketName: jsii.String("bucketName"),
//   							},
//   						},
//   					},
//   				},
//   				Inputs: []interface{}{
//   					&FlowNodeInputProperty{
//   						Category: jsii.String("category"),
//   						Expression: jsii.String("expression"),
//   						Name: jsii.String("name"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Outputs: []interface{}{
//   					&FlowNodeOutputProperty{
//   						Name: jsii.String("name"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-loopflownodeconfiguration.html
//
type CfnFlowPropsMixin_LoopFlowNodeConfigurationProperty struct {
	// The definition of the DoWhile loop nodes and connections between nodes in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-loopflownodeconfiguration.html#cfn-bedrock-flow-loopflownodeconfiguration-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
}

