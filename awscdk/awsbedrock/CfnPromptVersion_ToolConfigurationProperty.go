package awsbedrock


// Tool configuration.
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
	// List of Tools.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolconfiguration.html#cfn-bedrock-promptversion-toolconfiguration-tools
	//
	Tools interface{} `field:"required" json:"tools" yaml:"tools"`
	// Tool choice.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-toolconfiguration.html#cfn-bedrock-promptversion-toolconfiguration-toolchoice
	//
	ToolChoice interface{} `field:"optional" json:"toolChoice" yaml:"toolChoice"`
}

