package awscodepipeline


// Event for trigger with pull request filter.
//
// Example:
//   var sourceAction CodeStarConnectionsSourceAction
//   var buildAction CodeBuildAction
//
//
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   	Stages: []StageProps{
//   		&StageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []IAction{
//   				sourceAction,
//   			},
//   		},
//   		&StageProps{
//   			StageName: jsii.String("Build"),
//   			Actions: []IAction{
//   				buildAction,
//   			},
//   		},
//   	},
//   	Triggers: []TriggerProps{
//   		&TriggerProps{
//   			ProviderType: codepipeline.ProviderType_CODE_STAR_SOURCE_CONNECTION,
//   			GitConfiguration: &GitConfiguration{
//   				SourceAction: *SourceAction,
//   				PullRequestFilter: []GitPullRequestFilter{
//   					&GitPullRequestFilter{
//   						BranchesExcludes: []*string{
//   							jsii.String("exclude1"),
//   							jsii.String("exclude2"),
//   						},
//   						BranchesIncludes: []*string{
//   							jsii.String("include1"),
//   							jsii.String("include2"),
//   						},
//   						Events: []GitPullRequestEvent{
//   							codepipeline.GitPullRequestEvent_OPEN,
//   							codepipeline.GitPullRequestEvent_CLOSED,
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

