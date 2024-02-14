package awscodepipeline


// Git configuration for trigger.
//
// Example:
//   var sourceAction codeStarConnectionsSourceAction
//   var buildAction codeBuildAction
//
//
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   	Stages: []stageProps{
//   		&stageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []iAction{
//   				sourceAction,
//   			},
//   		},
//   		&stageProps{
//   			StageName: jsii.String("Build"),
//   			Actions: []*iAction{
//   				buildAction,
//   			},
//   		},
//   	},
//   	Triggers: []triggerProps{
//   		&triggerProps{
//   			ProviderType: codepipeline.ProviderType_CODE_STAR_SOURCE_CONNECTION,
//   			GitConfiguration: &GitConfiguration{
//   				SourceAction: *SourceAction,
//   				PushFilter: []gitPushFilter{
//   					&gitPushFilter{
//   						TagsExcludes: []*string{
//   							jsii.String("exclude1"),
//   							jsii.String("exclude2"),
//   						},
//   						TagsIncludes: []*string{
//   							jsii.String("include*"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
type GitConfiguration struct {
	// The pipeline source action where the trigger configuration, such as Git tags.
	//
	// The trigger configuration will start the pipeline upon the specified change only.
	// You can only specify one trigger configuration per source action.
	//
	// Since the provider for `sourceAction` must be `CodeStarSourceConnection`, you can use
	// `CodeStarConnectionsSourceAction` construct in `aws-codepipeline-actions` module.
	SourceAction IAction `field:"required" json:"sourceAction" yaml:"sourceAction"`
	// The field where the repository event that will start the pipeline, such as pushing Git tags, is specified with details.
	//
	// Git tags is the only supported event type.
	//
	// The length must be less than or equal to 3.
	// Default: - no filter.
	//
	PushFilter *[]*GitPushFilter `field:"optional" json:"pushFilter" yaml:"pushFilter"`
}

