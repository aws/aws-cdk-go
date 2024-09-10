package awsbedrock


// Properties for defining a `CfnFlow`.
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
//   cfnFlowProps := &CfnFlowProps{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	Definition: &FlowDefinitionProperty{
//   		Connections: []interface{}{
//   			&FlowConnectionProperty{
//   				Name: jsii.String("name"),
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Configuration: &FlowConnectionConfigurationProperty{
//   					Conditional: &FlowConditionalConnectionConfigurationProperty{
//   						Condition: jsii.String("condition"),
//   					},
//   					Data: &FlowDataConnectionConfigurationProperty{
//   						SourceOutput: jsii.String("sourceOutput"),
//   						TargetInput: jsii.String("targetInput"),
//   					},
//   				},
//   			},
//   		},
//   		Nodes: []interface{}{
//   			&FlowNodeProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Configuration: &FlowNodeConfigurationProperty{
//   					Agent: &AgentFlowNodeConfigurationProperty{
//   						AgentAliasArn: jsii.String("agentAliasArn"),
//   					},
//   					Collector: collector,
//   					Condition: &ConditionFlowNodeConfigurationProperty{
//   						Conditions: []interface{}{
//   							&FlowConditionProperty{
//   								Name: jsii.String("name"),
//
//   								// the properties below are optional
//   								Expression: jsii.String("expression"),
//   							},
//   						},
//   					},
//   					Input: input,
//   					Iterator: iterator,
//   					KnowledgeBase: &KnowledgeBaseFlowNodeConfigurationProperty{
//   						KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   						// the properties below are optional
//   						ModelId: jsii.String("modelId"),
//   					},
//   					LambdaFunction: &LambdaFunctionFlowNodeConfigurationProperty{
//   						LambdaArn: jsii.String("lambdaArn"),
//   					},
//   					Lex: &LexFlowNodeConfigurationProperty{
//   						BotAliasArn: jsii.String("botAliasArn"),
//   						LocaleId: jsii.String("localeId"),
//   					},
//   					Output: output,
//   					Prompt: &PromptFlowNodeConfigurationProperty{
//   						SourceConfiguration: &PromptFlowNodeSourceConfigurationProperty{
//   							Inline: &PromptFlowNodeInlineConfigurationProperty{
//   								ModelId: jsii.String("modelId"),
//   								TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   									Text: &TextPromptTemplateConfigurationProperty{
//   										Text: jsii.String("text"),
//
//   										// the properties below are optional
//   										InputVariables: []interface{}{
//   											&PromptInputVariableProperty{
//   												Name: jsii.String("name"),
//   											},
//   										},
//   									},
//   								},
//   								TemplateType: jsii.String("templateType"),
//
//   								// the properties below are optional
//   								InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   									Text: &PromptModelInferenceConfigurationProperty{
//   										MaxTokens: jsii.Number(123),
//   										StopSequences: []*string{
//   											jsii.String("stopSequences"),
//   										},
//   										Temperature: jsii.Number(123),
//   										TopK: jsii.Number(123),
//   										TopP: jsii.Number(123),
//   									},
//   								},
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
//   						Expression: jsii.String("expression"),
//   						Name: jsii.String("name"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Outputs: []interface{}{
//   					&FlowNodeOutputProperty{
//   						Name: jsii.String("name"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	DefinitionString: jsii.String("definitionString"),
//   	DefinitionSubstitutions: map[string]interface{}{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	Description: jsii.String("description"),
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
type CfnFlowProps struct {
	// The Amazon Resource Name (ARN) of the service role with permissions to create a flow.
	//
	// For more information, see [Create a service row for flows](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-permissions.html) in the Amazon Bedrock User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html#cfn-bedrock-flow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
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

