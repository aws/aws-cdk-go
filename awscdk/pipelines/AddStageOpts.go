package pipelines


// Options to pass to `addStage`.
//
// Example:
//   var pipeline CodePipeline
//
//   topic := sns.NewTopic(this, jsii.String("SecurityChangesTopic"))
//   topic.AddSubscription(subscriptions.NewEmailSubscription(jsii.String("test@email.com")))
//
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.AddStage(stage, &AddStageOpts{
//   	Pre: []Step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &PermissionsBroadeningCheckProps{
//   			Stage: *Stage,
//   			NotificationTopic: topic,
//   		}),
//   	},
//   })
//
type AddStageOpts struct {
	// Additional steps to run after all of the stacks in the stage.
	// Default: - No additional steps.
	//
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Default: - No additional steps.
	//
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
	// Instructions for stack level steps.
	// Default: - No additional instructions.
	//
	StackSteps *[]*StackSteps `field:"optional" json:"stackSteps" yaml:"stackSteps"`
}

