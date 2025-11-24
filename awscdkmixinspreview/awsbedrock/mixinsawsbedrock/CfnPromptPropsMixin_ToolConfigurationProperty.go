package mixinsawsbedrock


// Configuration information for the tools that you pass to a model.
//
// For more information, see [Tool use (function calling)](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var any interface{}
//   var auto interface{}
//   var json interface{}
//
//   toolConfigurationProperty := &ToolConfigurationProperty{
//   	ToolChoice: &ToolChoiceProperty{
//   		Any: any,
//   		Auto: auto,
//   		Tool: &SpecificToolChoiceProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Tools: []interface{}{
//   		&ToolProperty{
//   			CachePoint: &CachePointBlockProperty{
//   				Type: jsii.String("type"),
//   			},
//   			ToolSpec: &ToolSpecificationProperty{
//   				Description: jsii.String("description"),
//   				InputSchema: &ToolInputSchemaProperty{
//   					Json: json,
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolconfiguration.html
//
type CfnPromptPropsMixin_ToolConfigurationProperty struct {
	// If supported by model, forces the model to request a tool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolconfiguration.html#cfn-bedrock-prompt-toolconfiguration-toolchoice
	//
	ToolChoice interface{} `field:"optional" json:"toolChoice" yaml:"toolChoice"`
	// An array of tools that you want to pass to a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-toolconfiguration.html#cfn-bedrock-prompt-toolconfiguration-tools
	//
	Tools interface{} `field:"optional" json:"tools" yaml:"tools"`
}

