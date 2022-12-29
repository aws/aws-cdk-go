package awscodepipeline


// Properties of registering a custom Action.
//
// Example:
//   // Make a custom CodePipeline Action
//   // Make a custom CodePipeline Action
//   codepipeline.NewCustomActionRegistration(this, jsii.String("GenericGitSourceProviderResource"), &customActionRegistrationProps{
//   	category: codepipeline.actionCategory_SOURCE,
//   	artifactBounds: &actionArtifactBounds{
//   		minInputs: jsii.Number(0),
//   		maxInputs: jsii.Number(0),
//   		minOutputs: jsii.Number(1),
//   		maxOutputs: jsii.Number(1),
//   	},
//   	provider: jsii.String("GenericGitSource"),
//   	version: jsii.String("1"),
//   	entityUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	executionUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	actionProperties: []customActionProperty{
//   		&customActionProperty{
//   			name: jsii.String("Branch"),
//   			required: jsii.Boolean(true),
//   			key: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//   			queryable: jsii.Boolean(false),
//   			description: jsii.String("Git branch to pull"),
//   			type: jsii.String("String"),
//   		},
//   		&customActionProperty{
//   			name: jsii.String("GitUrl"),
//   			required: jsii.Boolean(true),
//   			key: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//   			queryable: jsii.Boolean(false),
//   			description: jsii.String("SSH git clone URL"),
//   			type: jsii.String("String"),
//   		},
//   	},
//   })
//
type CustomActionRegistrationProps struct {
	// The artifact bounds of the Action.
	ArtifactBounds *ActionArtifactBounds `field:"required" json:"artifactBounds" yaml:"artifactBounds"`
	// The category of the Action.
	Category ActionCategory `field:"required" json:"category" yaml:"category"`
	// The provider of the Action.
	//
	// For example, `'MyCustomActionProvider'`.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The properties used for customizing the instance of your Action.
	ActionProperties *[]*CustomActionProperty `field:"optional" json:"actionProperties" yaml:"actionProperties"`
	// The URL shown for the entire Action in the Pipeline UI.
	EntityUrl *string `field:"optional" json:"entityUrl" yaml:"entityUrl"`
	// The URL shown for a particular execution of an Action in the Pipeline UI.
	ExecutionUrl *string `field:"optional" json:"executionUrl" yaml:"executionUrl"`
	// The version of your Action.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

