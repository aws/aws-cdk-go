package pipelines


// Options to pass to `addStage`.
//
// Example:
//   var pipeline codePipeline
//
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//
//   pipeline.addStage(preprod, &addStageOpts{
//   	post: []step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &shellStepProps{
//   			commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.addStage(prod, &addStageOpts{
//   	pre: []*step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd")),
//   	},
//   })
//
// Experimental.
type AddStageOpts struct {
	// Additional steps to run after all of the stacks in the stage.
	// Experimental.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Experimental.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
	// Instructions for stack level steps.
	// Experimental.
	StackSteps *[]*StackSteps `field:"optional" json:"stackSteps" yaml:"stackSteps"`
}

