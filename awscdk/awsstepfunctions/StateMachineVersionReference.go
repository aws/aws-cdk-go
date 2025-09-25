package awsstepfunctions


// A reference to a StateMachineVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateMachineVersionReference := &StateMachineVersionReference{
//   	StateMachineVersionArn: jsii.String("stateMachineVersionArn"),
//   }
//
type StateMachineVersionReference struct {
	// The Arn of the StateMachineVersion resource.
	StateMachineVersionArn *string `field:"required" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
}

