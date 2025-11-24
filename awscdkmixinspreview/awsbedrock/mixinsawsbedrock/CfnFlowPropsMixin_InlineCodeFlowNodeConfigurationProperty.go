package mixinsawsbedrock


// Contains configurations for an inline code node in your flow.
//
// Inline code nodes let you write and execute code directly within your flow, enabling data transformations, custom logic, and integrations without needing an external Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inlineCodeFlowNodeConfigurationProperty := &InlineCodeFlowNodeConfigurationProperty{
//   	Code: jsii.String("code"),
//   	Language: jsii.String("language"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-inlinecodeflownodeconfiguration.html
//
type CfnFlowPropsMixin_InlineCodeFlowNodeConfigurationProperty struct {
	// The code that's executed in your inline code node.
	//
	// The code can access input data from previous nodes in the flow, perform operations on that data, and produce output that can be used by other nodes in your flow.
	//
	// The code must be valid in the programming `language` that you specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-inlinecodeflownodeconfiguration.html#cfn-bedrock-flow-inlinecodeflownodeconfiguration-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The programming language used by your inline code node.
	//
	// The code must be valid in the programming `language` that you specify. Currently, only Python 3 ( `Python_3` ) is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-inlinecodeflownodeconfiguration.html#cfn-bedrock-flow-inlinecodeflownodeconfiguration-language
	//
	Language *string `field:"optional" json:"language" yaml:"language"`
}

