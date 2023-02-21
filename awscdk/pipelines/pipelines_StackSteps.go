package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Instructions for additional steps that are run at stack level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//   var step step
//
//   stackSteps := &stackSteps{
//   	stack: stack,
//
//   	// the properties below are optional
//   	changeSet: []*step{
//   		step,
//   	},
//   	post: []*step{
//   		step,
//   	},
//   	pre: []*step{
//   		step,
//   	},
//   }
//
// Experimental.
type StackSteps struct {
	// The stack you want the steps to run in.
	// Experimental.
	Stack awscdk.Stack `field:"required" json:"stack" yaml:"stack"`
	// Steps that execute after stack is prepared but before stack is deployed.
	// Experimental.
	ChangeSet *[]Step `field:"optional" json:"changeSet" yaml:"changeSet"`
	// Steps that execute after stack is deployed.
	// Experimental.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Steps that execute before stack is prepared.
	// Experimental.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
}

