package awsstates


// Properties for CfnExecutionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnExecutionMixinProps := &CfnExecutionMixinProps{
//   	Input: jsii.String("input"),
//   	Name: jsii.String("name"),
//   	StateMachineArn: jsii.String("stateMachineArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html
//
type CfnExecutionMixinProps struct {
	// The string that contains the JSON input data for the execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html#cfn-states-execution-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The name of the execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html#cfn-states-execution-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the state machine that was executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-states-execution.html#cfn-states-execution-statemachinearn
	//
	StateMachineArn *string `field:"optional" json:"stateMachineArn" yaml:"stateMachineArn"`
}

