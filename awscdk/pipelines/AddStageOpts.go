package pipelines


// Options to pass to `addStage`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline codePipeline
//
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//
//   pipeline.AddStage(preprod, &AddStageOpts{
//   	Post: []step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &ShellStepProps{
//   			Commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.AddStage(prod, &AddStageOpts{
//   	Pre: []*step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd")),
//   	},
//   })
//
type AddStageOpts struct {
	// Additional steps to run after all of the stacks in the stage.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
	// Instructions for stack level steps.
	StackSteps *[]*StackSteps `field:"optional" json:"stackSteps" yaml:"stackSteps"`
}

