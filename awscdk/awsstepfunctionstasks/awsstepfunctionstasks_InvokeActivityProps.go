package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for FunctionTask.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   invokeActivityProps := &invokeActivityProps{
//   	heartbeat: duration,
//   }
//
// Deprecated: use `StepFunctionsInvokeActivity` and `StepFunctionsInvokeActivityProps`.
type InvokeActivityProps struct {
	// Maximum time between heart beats.
	//
	// If the time between heart beats takes longer than this, a 'Timeout' error is raised.
	// Deprecated: use `StepFunctionsInvokeActivity` and `StepFunctionsInvokeActivityProps`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
}

