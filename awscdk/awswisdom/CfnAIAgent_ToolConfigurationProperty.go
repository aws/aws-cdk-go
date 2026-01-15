package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var annotations interface{}
//   var inputSchema interface{}
//   var outputSchema interface{}
//
//   toolConfigurationProperty := &ToolConfigurationProperty{
//   	ToolName: jsii.String("toolName"),
//   	ToolType: jsii.String("toolType"),
//
//   	// the properties below are optional
//   	Annotations: annotations,
//   	Description: jsii.String("description"),
//   	InputSchema: inputSchema,
//   	Instruction: &ToolInstructionProperty{
//   		Examples: []*string{
//   			jsii.String("examples"),
//   		},
//   		Instruction: jsii.String("instruction"),
//   	},
//   	OutputFilters: []interface{}{
//   		&ToolOutputFilterProperty{
//   			JsonPath: jsii.String("jsonPath"),
//
//   			// the properties below are optional
//   			OutputConfiguration: &ToolOutputConfigurationProperty{
//   				OutputVariableNameOverride: jsii.String("outputVariableNameOverride"),
//   				SessionDataNamespace: jsii.String("sessionDataNamespace"),
//   			},
//   		},
//   	},
//   	OutputSchema: outputSchema,
//   	OverrideInputValues: []interface{}{
//   		&ToolOverrideInputValueProperty{
//   			JsonPath: jsii.String("jsonPath"),
//   			Value: &ToolOverrideInputValueConfigurationProperty{
//   				Constant: &ToolOverrideConstantInputValueProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Title: jsii.String("title"),
//   	ToolId: jsii.String("toolId"),
//   	UserInteractionConfiguration: &UserInteractionConfigurationProperty{
//   		IsUserConfirmationRequired: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html
//
type CfnAIAgent_ToolConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-toolname
	//
	ToolName *string `field:"required" json:"toolName" yaml:"toolName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-tooltype
	//
	ToolType *string `field:"required" json:"toolType" yaml:"toolType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-annotations
	//
	Annotations interface{} `field:"optional" json:"annotations" yaml:"annotations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-inputschema
	//
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-instruction
	//
	Instruction interface{} `field:"optional" json:"instruction" yaml:"instruction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-outputfilters
	//
	OutputFilters interface{} `field:"optional" json:"outputFilters" yaml:"outputFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-outputschema
	//
	OutputSchema interface{} `field:"optional" json:"outputSchema" yaml:"outputSchema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-overrideinputvalues
	//
	OverrideInputValues interface{} `field:"optional" json:"overrideInputValues" yaml:"overrideInputValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-toolid
	//
	ToolId *string `field:"optional" json:"toolId" yaml:"toolId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-toolconfiguration.html#cfn-wisdom-aiagent-toolconfiguration-userinteractionconfiguration
	//
	UserInteractionConfiguration interface{} `field:"optional" json:"userInteractionConfiguration" yaml:"userInteractionConfiguration"`
}

