package awsstepfunctions


// A reference to a StateMachineAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateMachineAliasReference := &StateMachineAliasReference{
//   	StateMachineAliasArn: jsii.String("stateMachineAliasArn"),
//   }
//
type StateMachineAliasReference struct {
	// The Arn of the StateMachineAlias resource.
	StateMachineAliasArn *string `field:"required" json:"stateMachineAliasArn" yaml:"stateMachineAliasArn"`
}

