package awsbedrock


// Contains configurations for an output from a node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowNodeOutputProperty := &FlowNodeOutputProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeoutput.html
//
type CfnFlowVersion_FlowNodeOutputProperty struct {
	// A name for the output that you can reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeoutput.html#cfn-bedrock-flowversion-flownodeoutput-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the output.
	//
	// If the output doesn't match this type at runtime, a validation error will be thrown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-flownodeoutput.html#cfn-bedrock-flowversion-flownodeoutput-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

