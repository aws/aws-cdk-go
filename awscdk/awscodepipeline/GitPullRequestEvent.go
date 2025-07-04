package awscodepipeline


// Event for trigger with pull request filter.
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
//   				PullRequestFilter: []gitPullRequestFilter{
//   					&gitPullRequestFilter{
//   						BranchesExcludes: []*string{
//   							jsii.String("exclude1"),
//   							jsii.String("exclude2"),
//   						},
//   						BranchesIncludes: []*string{
//   							jsii.String("include1"),
//   							jsii.String("include2"),
//   						},
//   						Events: []gitPullRequestEvent{
//   							codepipeline.*gitPullRequestEvent_OPEN,
//   							codepipeline.*gitPullRequestEvent_CLOSED,
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
type GitPullRequestEvent string

const (
	// OPEN.
	GitPullRequestEvent_OPEN GitPullRequestEvent = "OPEN"
	// UPDATED.
	GitPullRequestEvent_UPDATED GitPullRequestEvent = "UPDATED"
	// CLOSED.
	GitPullRequestEvent_CLOSED GitPullRequestEvent = "CLOSED"
)

