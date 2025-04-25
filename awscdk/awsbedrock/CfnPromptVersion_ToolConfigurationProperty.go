package awsbedrock


// Configuration information for the tools that you pass to a model.
//
// For more information, see [Tool use (function calling)](https://docs.aws.amazon.com/bedrock/latest/userguide/tool-use.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var any interface{}
//   var auto interface{}
//   var json interface{}
//
//   toolConfigurationProperty := &ToolConfigurationProperty{
//   	Tools: []interface{}{
//   		&ToolProperty{
//   			CachePoint: &CachePointBlockProperty{
//   				Type: jsii.String("type"),
//   			},
//   			ToolSpec: &ToolSpecificationProperty{
//   				InputSchema: &ToolInputSchemaProperty{
//   					Json: json,
//   				},
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ToolChoice: &ToolChoiceProperty{
//   		Any: any,
//   		Auto: auto,
//   		Tool: &SpecificToolChoiceProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolconfiguration.html
//
type CfnPromptVersion_ToolConfigurationProperty struct {
	// An array of tools that you want to pass to a model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolconfiguration.html#cfn-bedrock-promptversion-toolconfiguration-tools
	//
	Tools interface{} `field:"required" json:"tools" yaml:"tools"`
	// If supported by model, forces the model to request a tool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolconfiguration.html#cfn-bedrock-promptversion-toolconfiguration-toolchoice
	//
	ToolChoice interface{} `field:"optional" json:"toolChoice" yaml:"toolChoice"`
}

