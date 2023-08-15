package awscodepipeline


// Properties of registering a custom Action.
//
// Example:
//   // Make a custom CodePipeline Action
//   // Make a custom CodePipeline Action
//   codepipeline.NewCustomActionRegistration(this, jsii.String("GenericGitSourceProviderResource"), &CustomActionRegistrationProps{
//   	Category: codepipeline.ActionCategory_SOURCE,
//   	ArtifactBounds: &ActionArtifactBounds{
//   		MinInputs: jsii.Number(0),
//   		MaxInputs: jsii.Number(0),
//   		MinOutputs: jsii.Number(1),
//   		MaxOutputs: jsii.Number(1),
//   	},
//   	Provider: jsii.String("GenericGitSource"),
//   	Version: jsii.String("1"),
//   	EntityUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	ExecutionUrl: jsii.String("https://docs.aws.amazon.com/codepipeline/latest/userguide/actions-create-custom-action.html"),
//   	ActionProperties: []customActionProperty{
//   		&customActionProperty{
//   			Name: jsii.String("Branch"),
//   			Required: jsii.Boolean(true),
//   			Key: jsii.Boolean(false),
//   			Secret: jsii.Boolean(false),
//   			Queryable: jsii.Boolean(false),
//   			Description: jsii.String("Git branch to pull"),
//   			Type: jsii.String("String"),
//   		},
//   		&customActionProperty{
//   			Name: jsii.String("GitUrl"),
//   			Required: jsii.Boolean(true),
//   			Key: jsii.Boolean(false),
//   			Secret: jsii.Boolean(false),
//   			Queryable: jsii.Boolean(false),
//   			Description: jsii.String("SSH git clone URL"),
//   			Type: jsii.String("String"),
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
	// Default: [].
	//
	ActionProperties *[]*CustomActionProperty `field:"optional" json:"actionProperties" yaml:"actionProperties"`
	// The URL shown for the entire Action in the Pipeline UI.
	// Default: none.
	//
	EntityUrl *string `field:"optional" json:"entityUrl" yaml:"entityUrl"`
	// The URL shown for a particular execution of an Action in the Pipeline UI.
	// Default: none.
	//
	ExecutionUrl *string `field:"optional" json:"executionUrl" yaml:"executionUrl"`
	// The version of your Action.
	// Default: '1'.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

