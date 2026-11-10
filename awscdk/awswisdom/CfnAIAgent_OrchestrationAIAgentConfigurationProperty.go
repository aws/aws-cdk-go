package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var annotations interface{}
//   var inputSchema interface{}
//   var inputSchemas interface{}
//   var outputSchema interface{}
//   var outputSchemas interface{}
//
//   orchestrationAIAgentConfigurationProperty := &OrchestrationAIAgentConfigurationProperty{
//   	ConnectInstanceArn: jsii.String("connectInstanceArn"),
//   	InputSchemas: []interface{}{
//   		inputSchemas,
//   	},
//   	Locale: jsii.String("locale"),
//   	MultiAgentConfigurations: []interface{}{
//   		&MultiAgentConfigurationProperty{
//   			DelegateAgentConfiguration: &DelegateAgentConfigurationProperty{
//   				AgentTarget: &AgentTargetProperty{
//   					AiAgentId: jsii.String("aiAgentId"),
//   					ApplicationId: jsii.String("applicationId"),
//   				},
//
//   				// the properties below are optional
//   				Instruction: &MultiAgentInstructionProperty{
//   					Examples: []*string{
//   						jsii.String("examples"),
//   					},
//   					Instruction: jsii.String("instruction"),
//   				},
//   			},
//   			HandoffAgentConfiguration: &HandoffAgentConfigurationProperty{
//   				AgentTarget: &AgentTargetProperty{
//   					AiAgentId: jsii.String("aiAgentId"),
//   					ApplicationId: jsii.String("applicationId"),
//   				},
//
//   				// the properties below are optional
//   				AudioStreamingEnabled: jsii.Boolean(false),
//   				ImmediateHandoff: jsii.Boolean(false),
//   				Instruction: &MultiAgentInstructionProperty{
//   					Examples: []*string{
//   						jsii.String("examples"),
//   					},
//   					Instruction: jsii.String("instruction"),
//   				},
//   			},
//   		},
//   	},
//   	OrchestrationAiGuardrailId: jsii.String("orchestrationAiGuardrailId"),
//   	OrchestrationAiPromptId: jsii.String("orchestrationAiPromptId"),
//   	OutputSchemas: []interface{}{
//   		outputSchemas,
//   	},
//   	ToolConfigurations: []interface{}{
//   		&ToolConfigurationProperty{
//   			ToolName: jsii.String("toolName"),
//   			ToolType: jsii.String("toolType"),
//
//   			// the properties below are optional
//   			Annotations: annotations,
//   			Description: jsii.String("description"),
//   			InputSchema: inputSchema,
//   			Instruction: &ToolInstructionProperty{
//   				Examples: []*string{
//   					jsii.String("examples"),
//   				},
//   				Instruction: jsii.String("instruction"),
//   			},
//   			OutputFilters: []interface{}{
//   				&ToolOutputFilterProperty{
//   					JsonPath: jsii.String("jsonPath"),
//
//   					// the properties below are optional
//   					OutputConfiguration: &ToolOutputConfigurationProperty{
//   						OutputVariableNameOverride: jsii.String("outputVariableNameOverride"),
//   						SessionDataNamespace: jsii.String("sessionDataNamespace"),
//   					},
//   				},
//   			},
//   			OutputSchema: outputSchema,
//   			OverrideInputValues: []interface{}{
//   				&ToolOverrideInputValueProperty{
//   					JsonPath: jsii.String("jsonPath"),
//   					Value: &ToolOverrideInputValueConfigurationProperty{
//   						Constant: &ToolOverrideConstantInputValueProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			Title: jsii.String("title"),
//   			ToolId: jsii.String("toolId"),
//   			UserInteractionConfiguration: &UserInteractionConfigurationProperty{
//   				IsUserConfirmationRequired: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html
//
type CfnAIAgent_OrchestrationAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-connectinstancearn
	//
	ConnectInstanceArn *string `field:"optional" json:"connectInstanceArn" yaml:"connectInstanceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-inputschemas
	//
	InputSchemas interface{} `field:"optional" json:"inputSchemas" yaml:"inputSchemas"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-multiagentconfigurations
	//
	MultiAgentConfigurations interface{} `field:"optional" json:"multiAgentConfigurations" yaml:"multiAgentConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-orchestrationaiguardrailid
	//
	OrchestrationAiGuardrailId *string `field:"optional" json:"orchestrationAiGuardrailId" yaml:"orchestrationAiGuardrailId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-orchestrationaipromptid
	//
	OrchestrationAiPromptId *string `field:"optional" json:"orchestrationAiPromptId" yaml:"orchestrationAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-outputschemas
	//
	OutputSchemas interface{} `field:"optional" json:"outputSchemas" yaml:"outputSchemas"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-toolconfigurations
	//
	ToolConfigurations interface{} `field:"optional" json:"toolConfigurations" yaml:"toolConfigurations"`
}

