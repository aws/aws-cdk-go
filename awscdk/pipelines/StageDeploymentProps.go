package pipelines


// Properties for a `StageDeployment`.
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
//   stageDeploymentProps := &StageDeploymentProps{
//   	Post: []Step{
//   		step,
//   	},
//   	Pre: []Step{
//   		step,
//   	},
//   	StackSteps: []StackSteps{
//   		&StackSteps{
//   			Stack: stack,
//
//   			// the properties below are optional
//   			ChangeSet: []Step{
//   				step,
//   			},
//   			Post: []Step{
//   				step,
//   			},
//   			Pre: []Step{
//   				step,
//   			},
//   		},
//   	},
//   	StageName: jsii.String("stageName"),
//   }
//
type StageDeploymentProps struct {
	// Additional steps to run after all of the stacks in the stage.
	// Default: - No additional steps.
	//
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Default: - No additional steps.
	//
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
	// Instructions for additional steps that are run at the stack level.
	// Default: - No additional instructions.
	//
	StackSteps *[]*StackSteps `field:"optional" json:"stackSteps" yaml:"stackSteps"`
	// Stage name to use in the pipeline.
	// Default: - Use Stage's construct ID.
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

