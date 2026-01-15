package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var annotations interface{}
//   var inputSchema interface{}
//   var outputSchema interface{}
//
//   orchestrationAIAgentConfigurationProperty := &OrchestrationAIAgentConfigurationProperty{
//   	ConnectInstanceArn: jsii.String("connectInstanceArn"),
//   	Locale: jsii.String("locale"),
//   	OrchestrationAiGuardrailId: jsii.String("orchestrationAiGuardrailId"),
//   	OrchestrationAiPromptId: jsii.String("orchestrationAiPromptId"),
//   	ToolConfigurations: []interface{}{
//   		&ToolConfigurationProperty{
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
//   			ToolName: jsii.String("toolName"),
//   			ToolType: jsii.String("toolType"),
//   			UserInteractionConfiguration: &UserInteractionConfigurationProperty{
//   				IsUserConfirmationRequired: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html
//
type CfnAIAgentPropsMixin_OrchestrationAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-connectinstancearn
	//
	ConnectInstanceArn *string `field:"optional" json:"connectInstanceArn" yaml:"connectInstanceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-orchestrationaiguardrailid
	//
	OrchestrationAiGuardrailId *string `field:"optional" json:"orchestrationAiGuardrailId" yaml:"orchestrationAiGuardrailId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-orchestrationaipromptid
	//
	OrchestrationAiPromptId *string `field:"optional" json:"orchestrationAiPromptId" yaml:"orchestrationAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-orchestrationaiagentconfiguration.html#cfn-wisdom-aiagent-orchestrationaiagentconfiguration-toolconfigurations
	//
	ToolConfigurations interface{} `field:"optional" json:"toolConfigurations" yaml:"toolConfigurations"`
}

