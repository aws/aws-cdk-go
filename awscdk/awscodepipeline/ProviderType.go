package awscodepipeline


// Provider type for trigger.
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
//   				PushFilter: []GitPushFilter{
//   					&GitPushFilter{
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
type ProviderType string

const (
	// CodeStarSourceConnection.
	ProviderType_CODE_STAR_SOURCE_CONNECTION ProviderType = "CODE_STAR_SOURCE_CONNECTION"
)

