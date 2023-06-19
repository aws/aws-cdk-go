package pipelines


// Properties for a `StageDeployment`.
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
//   stageDeploymentProps := &StageDeploymentProps{
//   	Post: []*step{
//   		step,
//   	},
//   	Pre: []*step{
//   		step,
//   	},
//   	StackSteps: []stackSteps{
//   		&stackSteps{
//   			Stack: stack,
//
//   			// the properties below are optional
//   			ChangeSet: []*step{
//   				step,
//   			},
//   			Post: []*step{
//   				step,
//   			},
//   			Pre: []*step{
//   				step,
//   			},
//   		},
//   	},
//   	StageName: jsii.String("stageName"),
//   }
//
// Experimental.
type StageDeploymentProps struct {
	// Additional steps to run after all of the stacks in the stage.
	// Experimental.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Experimental.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
	// Instructions for additional steps that are run at the stack level.
	// Experimental.
	StackSteps *[]*StackSteps `field:"optional" json:"stackSteps" yaml:"stackSteps"`
	// Stage name to use in the pipeline.
	// Experimental.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

