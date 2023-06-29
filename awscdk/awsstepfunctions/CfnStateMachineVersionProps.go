package awsstepfunctions


// Properties for defining a `CfnStateMachineVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStateMachineVersionProps := &CfnStateMachineVersionProps{
//   	StateMachineArn: jsii.String("stateMachineArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	StateMachineRevisionId: jsii.String("stateMachineRevisionId"),
//   }
//
type CfnStateMachineVersionProps struct {
	// The Amazon Resource Name (ARN) of the state machine.
	StateMachineArn *string `field:"required" json:"stateMachineArn" yaml:"stateMachineArn"`
	// An optional description of the state machine version.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Identifier for a state machine revision, which is an immutable, read-only snapshot of a state machine’s definition and configuration.
	//
	// Only publish the state machine version if the current state machine's revision ID matches the specified ID. Use this option to avoid publishing a version if the state machine has changed since you last updated it.
	//
	// To specify the initial state machine revision, set the value as `INITIAL` .
	StateMachineRevisionId *string `field:"optional" json:"stateMachineRevisionId" yaml:"stateMachineRevisionId"`
}

