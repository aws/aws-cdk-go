package awsbedrock


// Defines a condition in the condition node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowConditionProperty := &FlowConditionProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowcondition.html
//
type CfnFlow_FlowConditionProperty struct {
	// A name for the condition that you can reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowcondition.html#cfn-bedrock-flow-flowcondition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Defines the condition.
	//
	// You must refer to at least one of the inputs in the condition. For more information, expand the Condition node section in [Node types in prompt flows](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-how-it-works.html#flows-nodes) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowcondition.html#cfn-bedrock-flow-flowcondition-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

