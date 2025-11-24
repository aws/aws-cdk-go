package mixinsawsbedrock


// Contains configurations for an input in an Amazon Bedrock Flows node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   flowNodeInputProperty := &FlowNodeInputProperty{
//   	Expression: jsii.String("expression"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeinput.html
//
type CfnFlowVersionPropsMixin_FlowNodeInputProperty struct {
	// An expression that formats the input for the node.
	//
	// For an explanation of how to create expressions, see [Expressions in Prompt flows in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-expressions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeinput.html#cfn-bedrock-flowversion-flownodeinput-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Specifies a name for the input that you can reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeinput.html#cfn-bedrock-flowversion-flownodeinput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the data type of the input.
	//
	// If the input doesn't match this type at runtime, a validation error will be thrown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeinput.html#cfn-bedrock-flowversion-flownodeinput-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

