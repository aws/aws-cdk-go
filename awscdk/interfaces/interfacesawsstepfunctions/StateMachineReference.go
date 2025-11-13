package interfacesawsstepfunctions


// A reference to a StateMachine resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateMachineReference := &StateMachineReference{
//   	StateMachineArn: jsii.String("stateMachineArn"),
//   }
//
type StateMachineReference struct {
	// The Arn of the StateMachine resource.
	StateMachineArn *string `field:"required" json:"stateMachineArn" yaml:"stateMachineArn"`
}

