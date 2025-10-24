package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Instructions for additional steps that are run at stack level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack Stack
//   var step Step
//
//   stackSteps := &StackSteps{
//   	Stack: stack,
//
//   	// the properties below are optional
//   	ChangeSet: []Step{
//   		step,
//   	},
//   	Post: []Step{
//   		step,
//   	},
//   	Pre: []Step{
//   		step,
//   	},
//   }
//
type StackSteps struct {
	// The stack you want the steps to run in.
	Stack awscdk.Stack `field:"required" json:"stack" yaml:"stack"`
	// Steps that execute after stack is prepared but before stack is deployed.
	// Default: - no additional steps.
	//
	ChangeSet *[]Step `field:"optional" json:"changeSet" yaml:"changeSet"`
	// Steps that execute after stack is deployed.
	// Default: - no additional steps.
	//
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Steps that execute before stack is prepared.
	// Default: - no additional steps.
	//
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
}

