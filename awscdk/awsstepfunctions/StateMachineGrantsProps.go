package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstepfunctions"
)

// Properties for StateMachineGrants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stateMachineRef IStateMachineRef
//
//   stateMachineGrantsProps := &StateMachineGrantsProps{
//   	Resource: stateMachineRef,
//   }
//
type StateMachineGrantsProps struct {
	// The resource on which actions will be allowed.
	Resource interfacesawsstepfunctions.IStateMachineRef `field:"required" json:"resource" yaml:"resource"`
}

