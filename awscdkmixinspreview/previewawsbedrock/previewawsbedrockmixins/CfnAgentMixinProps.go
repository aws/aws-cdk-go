package previewawsbedrockmixins


// Properties for CfnAgentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//
//   cfnAgentMixinProps := &CfnAgentMixinProps{
//   	ActionGroups: []interface{}{
//   		&AgentActionGroupProperty{
//   			ActionGroupExecutor: &ActionGroupExecutorProperty{
//   				CustomControl: jsii.String("customControl"),
//   				Lambda: jsii.String("lambda"),
//   			},
//   			ActionGroupName: jsii.String("actionGroupName"),
//   			ActionGroupState: jsii.String("actionGroupState"),
//   			ApiSchema: &APISchemaProperty{
//   				Payload: jsii.String("payload"),
//   				S3: &S3IdentifierProperty{
//   					S3BucketName: jsii.String("s3BucketName"),
//   					S3ObjectKey: jsii.String("s3ObjectKey"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			FunctionSchema: &FunctionSchemaProperty{
//   				Functions: []interface{}{
//   					&FunctionProperty{
//   						Description: jsii.String("description"),
//   						Name: jsii.String("name"),
//   						Parameters: map[string]interface{}{
//   							"parametersKey": &ParameterDetailProperty{
//   								"description": jsii.String("description"),
//   								"required": jsii.Boolean(false),
//   								"type": jsii.String("type"),
//   							},
//   						},
//   						RequireConfirmation: jsii.String("requireConfirmation"),
//   					},
//   				},
//   			},
//   			ParentActionGroupSignature: jsii.String("parentActionGroupSignature"),
//   			SkipResourceInUseCheckOnDelete: jsii.Boolean(false),
//   		},
//   	},
//   	AgentCollaboration: jsii.String("agentCollaboration"),
//   	AgentCollaborators: []interface{}{
//   		&AgentCollaboratorProperty{
//   			AgentDescriptor: &AgentDescriptorProperty{
//   				AliasArn: jsii.String("aliasArn"),
//   			},
//   			CollaborationInstruction: jsii.String("collaborationInstruction"),
//   			CollaboratorName: jsii.String("collaboratorName"),
//   			RelayConversationHistory: jsii.String("relayConversationHistory"),
//   		},
//   	},
//   	AgentName: jsii.String("agentName"),
//   	AgentResourceRoleArn: jsii.String("agentResourceRoleArn"),
//   	AutoPrepare: jsii.Boolean(false),
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	CustomOrchestration: &CustomOrchestrationProperty{
//   		Executor: &OrchestrationExecutorProperty{
//   			Lambda: jsii.String("lambda"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FoundationModel: jsii.String("foundationModel"),
//   	GuardrailConfiguration: &GuardrailConfigurationProperty{
//   		GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   		GuardrailVersion: jsii.String("guardrailVersion"),
//   	},
//   	IdleSessionTtlInSeconds: jsii.Number(123),
//   	Instruction: jsii.String("instruction"),
//   	KnowledgeBases: []interface{}{
//   		&AgentKnowledgeBaseProperty{
//   			Description: jsii.String("description"),
//   			KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   			KnowledgeBaseState: jsii.String("knowledgeBaseState"),
//   		},
//   	},
//   	MemoryConfiguration: &MemoryConfigurationProperty{
//   		EnabledMemoryTypes: []*string{
//   			jsii.String("enabledMemoryTypes"),
//   		},
//   		SessionSummaryConfiguration: &SessionSummaryConfigurationProperty{
//   			MaxRecentSessions: jsii.Number(123),
//   		},
//   		StorageDays: jsii.Number(123),
//   	},
//   	OrchestrationType: jsii.String("orchestrationType"),
//   	PromptOverrideConfiguration: &PromptOverrideConfigurationProperty{
//   		OverrideLambda: jsii.String("overrideLambda"),
//   		PromptConfigurations: []interface{}{
//   			&PromptConfigurationProperty{
//   				AdditionalModelRequestFields: additionalModelRequestFields,
//   				BasePromptTemplate: jsii.String("basePromptTemplate"),
//   				FoundationModel: jsii.String("foundationModel"),
//   				InferenceConfiguration: &InferenceConfigurationProperty{
//   					MaximumLength: jsii.Number(123),
//   					StopSequences: []*string{
//   						jsii.String("stopSequences"),
//   					},
//   					Temperature: jsii.Number(123),
//   					TopK: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
//   				ParserMode: jsii.String("parserMode"),
//   				PromptCreationMode: jsii.String("promptCreationMode"),
//   				PromptState: jsii.String("promptState"),
//   				PromptType: jsii.String("promptType"),
//   			},
//   		},
//   	},
//   	SkipResourceInUseCheckOnDelete: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TestAliasTags: map[string]*string{
//   		"testAliasTagsKey": jsii.String("testAliasTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html
//
type CfnAgentMixinProps struct {
	// The action groups that belong to an agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-actiongroups
	//
	ActionGroups interface{} `field:"optional" json:"actionGroups" yaml:"actionGroups"`
	// The agent's collaboration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-agentcollaboration
	//
	AgentCollaboration *string `field:"optional" json:"agentCollaboration" yaml:"agentCollaboration"`
	// List of Agent Collaborators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-agentcollaborators
	//
	AgentCollaborators interface{} `field:"optional" json:"agentCollaborators" yaml:"agentCollaborators"`
	// The name of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-agentname
	//
	AgentName *string `field:"optional" json:"agentName" yaml:"agentName"`
	// The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-agentresourcerolearn
	//
	AgentResourceRoleArn *string `field:"optional" json:"agentResourceRoleArn" yaml:"agentResourceRoleArn"`
	// Specifies whether to automatically update the `DRAFT` version of the agent after making changes to the agent.
	//
	// The `DRAFT` version can be continually iterated upon during internal development. By default, this value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-autoprepare
	//
	// Default: - false.
	//
	AutoPrepare interface{} `field:"optional" json:"autoPrepare" yaml:"autoPrepare"`
	// The Amazon Resource Name (ARN) of the AWS  key that encrypts the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-customerencryptionkeyarn
	//
	CustomerEncryptionKeyArn *string `field:"optional" json:"customerEncryptionKeyArn" yaml:"customerEncryptionKeyArn"`
	// Contains custom orchestration configurations for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-customorchestration
	//
	CustomOrchestration interface{} `field:"optional" json:"customOrchestration" yaml:"customOrchestration"`
	// The description of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The foundation model used for orchestration by the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-foundationmodel
	//
	FoundationModel *string `field:"optional" json:"foundationModel" yaml:"foundationModel"`
	// Details about the guardrail associated with the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-guardrailconfiguration
	//
	GuardrailConfiguration interface{} `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// The number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent.
	//
	// A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Bedrock deletes any data provided before the timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-idlesessionttlinseconds
	//
	IdleSessionTtlInSeconds *float64 `field:"optional" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// Instructions that tell the agent what it should do and how it should interact with users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-instruction
	//
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// The knowledge bases associated with the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-knowledgebases
	//
	KnowledgeBases interface{} `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// Contains memory configuration for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-memoryconfiguration
	//
	MemoryConfiguration interface{} `field:"optional" json:"memoryConfiguration" yaml:"memoryConfiguration"`
	// Specifies the orchestration strategy for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-orchestrationtype
	//
	OrchestrationType *string `field:"optional" json:"orchestrationType" yaml:"orchestrationType"`
	// Contains configurations to override prompt templates in different parts of an agent sequence.
	//
	// For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-promptoverrideconfiguration
	//
	PromptOverrideConfiguration interface{} `field:"optional" json:"promptOverrideConfiguration" yaml:"promptOverrideConfiguration"`
	// Specifies whether to delete the resource even if it's in use.
	//
	// By default, this value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-skipresourceinusecheckondelete
	//
	// Default: - false.
	//
	SkipResourceInUseCheckOnDelete interface{} `field:"optional" json:"skipResourceInUseCheckOnDelete" yaml:"skipResourceInUseCheckOnDelete"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html#cfn-bedrock-agent-testaliastags
	//
	TestAliasTags interface{} `field:"optional" json:"testAliasTags" yaml:"testAliasTags"`
}

