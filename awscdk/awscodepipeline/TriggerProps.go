package awscodepipeline


// Properties of trigger.
//
// Example:
//   var pipeline pipeline
//   var sourceAction codeStarConnectionsSourceAction
//
//
//   pipeline.AddTrigger(&TriggerProps{
//   	ProviderType: codepipeline.ProviderType_CODE_STAR_SOURCE_CONNECTION,
//   	GitConfiguration: &GitConfiguration{
//   		SourceAction: *SourceAction,
//   		PushFilter: []gitPushFilter{
//   			&gitPushFilter{
//   				TagsExcludes: []*string{
//   					jsii.String("exclude1"),
//   					jsii.String("exclude2"),
//   				},
//   				TagsIncludes: []*string{
//   					jsii.String("include*"),
//   				},
//   			},
//   		},
//   	},
//   })
//
type TriggerProps struct {
	// The source provider for the event, such as connections configured for a repository with Git tags, for the specified trigger configuration.
	ProviderType ProviderType `field:"required" json:"providerType" yaml:"providerType"`
	// Provides the filter criteria and the source stage for the repository event that starts the pipeline, such as Git tags.
	// Default: - no configuration.
	//
	GitConfiguration *GitConfiguration `field:"optional" json:"gitConfiguration" yaml:"gitConfiguration"`
}

