package awsbedrock


// Contains configurations about a node in the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//   var collector interface{}
//   var flowNodeProperty_ FlowNodeProperty
//   var input interface{}
//   var iterator interface{}
//   var loopInput interface{}
//   var output interface{}
//
//   flowNodeProperty := &FlowNodeProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Configuration: &FlowNodeConfigurationProperty{
//   		Agent: &AgentFlowNodeConfigurationProperty{
//   			AgentAliasArn: jsii.String("agentAliasArn"),
//   		},
//   		Collector: collector,
//   		Condition: &ConditionFlowNodeConfigurationProperty{
//   			Conditions: []interface{}{
//   				&FlowConditionProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Expression: jsii.String("expression"),
//   				},
//   			},
//   		},
//   		InlineCode: &InlineCodeFlowNodeConfigurationProperty{
//   			Code: jsii.String("code"),
//   			Language: jsii.String("language"),
//   		},
//   		Input: input,
//   		Iterator: iterator,
//   		KnowledgeBase: &KnowledgeBaseFlowNodeConfigurationProperty{
//   			KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   			// the properties below are optional
//   			GuardrailConfiguration: &GuardrailConfigurationProperty{
//   				GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   				GuardrailVersion: jsii.String("guardrailVersion"),
//   			},
//   			InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   				Text: &PromptModelInferenceConfigurationProperty{
//   					MaxTokens: jsii.Number(123),
//   					StopSequences: []*string{
//   						jsii.String("stopSequences"),
//   					},
//   					Temperature: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
//   			},
//   			ModelId: jsii.String("modelId"),
//   			NumberOfResults: jsii.Number(123),
//   			OrchestrationConfiguration: &KnowledgeBaseOrchestrationConfigurationProperty{
//   				AdditionalModelRequestFields: additionalModelRequestFields,
//   				InferenceConfig: &PromptInferenceConfigurationProperty{
//   					Text: &PromptModelInferenceConfigurationProperty{
//   						MaxTokens: jsii.Number(123),
//   						StopSequences: []*string{
//   							jsii.String("stopSequences"),
//   						},
//   						Temperature: jsii.Number(123),
//   						TopP: jsii.Number(123),
//   					},
//   				},
//   				PerformanceConfig: &PerformanceConfigurationProperty{
//   					Latency: jsii.String("latency"),
//   				},
//   				PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   					TextPromptTemplate: jsii.String("textPromptTemplate"),
//   				},
//   			},
//   			PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   				TextPromptTemplate: jsii.String("textPromptTemplate"),
//   			},
//   			RerankingConfiguration: &VectorSearchRerankingConfigurationProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				BedrockRerankingConfiguration: &VectorSearchBedrockRerankingConfigurationProperty{
//   					ModelConfiguration: &VectorSearchBedrockRerankingModelConfigurationProperty{
//   						ModelArn: jsii.String("modelArn"),
//
//   						// the properties below are optional
//   						AdditionalModelRequestFields: additionalModelRequestFields,
//   					},
//
//   					// the properties below are optional
//   					MetadataConfiguration: &MetadataConfigurationForRerankingProperty{
//   						SelectionMode: jsii.String("selectionMode"),
//
//   						// the properties below are optional
//   						SelectiveModeConfiguration: &RerankingMetadataSelectiveModeConfigurationProperty{
//   							FieldsToExclude: []interface{}{
//   								&FieldForRerankingProperty{
//   									FieldName: jsii.String("fieldName"),
//   								},
//   							},
//   							FieldsToInclude: []interface{}{
//   								&FieldForRerankingProperty{
//   									FieldName: jsii.String("fieldName"),
//   								},
//   							},
//   						},
//   					},
//   					NumberOfRerankedResults: jsii.Number(123),
//   				},
//   			},
//   		},
//   		LambdaFunction: &LambdaFunctionFlowNodeConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   		Lex: &LexFlowNodeConfigurationProperty{
//   			BotAliasArn: jsii.String("botAliasArn"),
//   			LocaleId: jsii.String("localeId"),
//   		},
//   		Loop: &LoopFlowNodeConfigurationProperty{
//   			Definition: &FlowDefinitionProperty{
//   				Connections: []interface{}{
//   					&FlowConnectionProperty{
//   						Name: jsii.String("name"),
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						Configuration: &FlowConnectionConfigurationProperty{
//   							Conditional: &FlowConditionalConnectionConfigurationProperty{
//   								Condition: jsii.String("condition"),
//   							},
//   							Data: &FlowDataConnectionConfigurationProperty{
//   								SourceOutput: jsii.String("sourceOutput"),
//   								TargetInput: jsii.String("targetInput"),
//   							},
//   						},
//   					},
//   				},
//   				Nodes: []interface{}{
//   					flowNodeProperty_,
//   				},
//   			},
//   		},
//   		LoopController: &LoopControllerFlowNodeConfigurationProperty{
//   			ContinueCondition: &FlowConditionProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Expression: jsii.String("expression"),
//   			},
//
//   			// the properties below are optional
//   			MaxIterations: jsii.Number(123),
//   		},
//   		LoopInput: loopInput,
//   		Output: output,
//   		Prompt: &PromptFlowNodeConfigurationProperty{
//   			SourceConfiguration: &PromptFlowNodeSourceConfigurationProperty{
//   				Inline: &PromptFlowNodeInlineConfigurationProperty{
//   					ModelId: jsii.String("modelId"),
//   					TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   						Text: &TextPromptTemplateConfigurationProperty{
//   							Text: jsii.String("text"),
//
//   							// the properties below are optional
//   							InputVariables: []interface{}{
//   								&PromptInputVariableProperty{
//   									Name: jsii.String("name"),
//   								},
//   							},
//   						},
//   					},
//   					TemplateType: jsii.String("templateType"),
//
//   					// the properties below are optional
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
//   				},
//   				Resource: &PromptFlowNodeResourceConfigurationProperty{
//   					PromptArn: jsii.String("promptArn"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			GuardrailConfiguration: &GuardrailConfigurationProperty{
//   				GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   				GuardrailVersion: jsii.String("guardrailVersion"),
//   			},
//   		},
//   		Retrieval: &RetrievalFlowNodeConfigurationProperty{
//   			ServiceConfiguration: &RetrievalFlowNodeServiceConfigurationProperty{
//   				S3: &RetrievalFlowNodeS3ConfigurationProperty{
//   					BucketName: jsii.String("bucketName"),
//   				},
//   			},
//   		},
//   		Storage: &StorageFlowNodeConfigurationProperty{
//   			ServiceConfiguration: &StorageFlowNodeServiceConfigurationProperty{
//   				S3: &StorageFlowNodeS3ConfigurationProperty{
//   					BucketName: jsii.String("bucketName"),
//   				},
//   			},
//   		},
//   	},
//   	Inputs: []interface{}{
//   		&FlowNodeInputProperty{
//   			Expression: jsii.String("expression"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Outputs: []interface{}{
//   		&FlowNodeOutputProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownode.html
//
type CfnFlowVersion_FlowNodeProperty struct {
	// A name for the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownode.html#cfn-bedrock-flowversion-flownode-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of node.
	//
	// This value must match the name of the key that you provide in the configuration you provide in the `FlowNodeConfiguration` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownode.html#cfn-bedrock-flowversion-flownode-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Contains configurations for the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownode.html#cfn-bedrock-flowversion-flownode-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// An array of objects, each of which contains information about an input into the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownode.html#cfn-bedrock-flowversion-flownode-inputs
	//
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
	// A list of objects, each of which contains information about an output from the node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownode.html#cfn-bedrock-flowversion-flownode-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
}

