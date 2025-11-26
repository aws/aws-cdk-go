package previewawsbedrockmixins


// Contains configurations for an input in an Amazon Bedrock Flows node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   flowNodeInputProperty := &FlowNodeInputProperty{
//   	Category: jsii.String("category"),
//   	Expression: jsii.String("expression"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeinput.html
//
type CfnFlowPropsMixin_FlowNodeInputProperty struct {
	// Specifies how input data flows between iterations in a DoWhile loop.
	//
	// - `LoopCondition` - Controls whether the loop continues by evaluating condition expressions against the input data. Use this category to define the condition that determines if the loop should continue.
	// - `ReturnValueToLoopStart` - Defines data to pass back to the start of the loop's next iteration. Use this category for variables that you want to update for each loop iteration.
	// - `ExitLoop` - Defines the value that's available once the loop ends. Use this category to expose loop results to nodes outside the loop.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeinput.html#cfn-bedrock-flow-flownodeinput-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// An expression that formats the input for the node.
	//
	// For an explanation of how to create expressions, see [Expressions in Prompt flows in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-expressions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeinput.html#cfn-bedrock-flow-flownodeinput-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Specifies a name for the input that you can reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeinput.html#cfn-bedrock-flow-flownodeinput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the data type of the input.
	//
	// If the input doesn't match this type at runtime, a validation error will be thrown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flownodeinput.html#cfn-bedrock-flow-flownodeinput-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

