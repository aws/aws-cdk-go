package awsstates


// Properties for defining a `CfnExecution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExecutionProps := &CfnExecutionProps{
//   	StateMachineArn: jsii.String("stateMachineArn"),
//
//   	// the properties below are optional
//   	Input: jsii.String("input"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html
//
type CfnExecutionProps struct {
	// The Amazon Resource Name (ARN) of the state machine that was executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html#cfn-states-execution-statemachinearn
	//
	StateMachineArn *string `field:"required" json:"stateMachineArn" yaml:"stateMachineArn"`
	// The string that contains the JSON input data for the execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html#cfn-states-execution-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The name of the execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html#cfn-states-execution-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

