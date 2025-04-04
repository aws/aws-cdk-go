package awscodepipeline


// Provider type for trigger.
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
type ProviderType string

const (
	// CodeStarSourceConnection.
	ProviderType_CODE_STAR_SOURCE_CONNECTION ProviderType = "CODE_STAR_SOURCE_CONNECTION"
)

