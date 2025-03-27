package awsbedrock


// Contains configurations about a node in the flow.
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
//   			ModelId: jsii.String("modelId"),
//   		},
//   		LambdaFunction: &LambdaFunctionFlowNodeConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   		Lex: &LexFlowNodeConfigurationProperty{
//   			BotAliasArn: jsii.String("botAliasArn"),
//   			LocaleId: jsii.String("localeId"),
//   		},
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

