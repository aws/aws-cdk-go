package mixinsawsbedrock


// Properties for CfnFlowPropsMixin.
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
//   cfnFlowMixinProps := &CfnFlowMixinProps{
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
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
//   					Loop: &LoopFlowNodeConfigurationProperty{
//   						Definition: flowDefinitionProperty_,
//   					},
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
//   	DefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   	DefinitionString: jsii.String("definitionString"),
//   	DefinitionSubstitutions: map[string]interface{}{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	Description: jsii.String("description"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TestAliasTags: map[string]*string{
//   		"testAliasTagsKey": jsii.String("testAliasTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html
//
type CfnFlowMixinProps struct {
	// The Amazon Resource Name (ARN) of the KMS key that the flow is encrypted with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-customerencryptionkeyarn
	//
	CustomerEncryptionKeyArn *string `field:"optional" json:"customerEncryptionKeyArn" yaml:"customerEncryptionKeyArn"`
	// The definition of the nodes and connections between the nodes in the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The Amazon S3 location of the flow definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-definitions3location
	//
	DefinitionS3Location interface{} `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
	// The definition of the flow as a JSON-formatted string.
	//
	// The string must match the format in [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-definitionstring
	//
	DefinitionString *string `field:"optional" json:"definitionString" yaml:"definitionString"`
	// A map that specifies the mappings for placeholder variables in the prompt flow definition.
	//
	// This enables the customer to inject values obtained at runtime. Variables can be template parameter names, resource logical IDs, resource attributes, or a variable in a key-value map. Only supported with the `DefinitionString` and `DefinitionS3Location` fields.
	//
	// Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-definitionsubstitutions
	//
	DefinitionSubstitutions interface{} `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// A description of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the service role with permissions to create a flow.
	//
	// For more information, see [Create a service row for flows](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-permissions.html) in the Amazon Bedrock User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// A map of tag keys and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-testaliastags
	//
	TestAliasTags interface{} `field:"optional" json:"testAliasTags" yaml:"testAliasTags"`
}

