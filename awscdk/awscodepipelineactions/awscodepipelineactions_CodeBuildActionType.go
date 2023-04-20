package awscodepipelineactions


// The type of the CodeBuild action that determines its CodePipeline Category - Build, or Test.
//
// The default is Build.
//
// Example:
//   var project pipelineProject
//
//   sourceOutput := codepipeline.NewArtifact()
//   testAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("IntegrationTest"),
//   	Project: Project,
//   	Input: sourceOutput,
//   	Type: codepipeline_actions.CodeBuildActionType_TEST,
//   })
//
// Experimental.
type CodeBuildActionType string

const (
	// The action will have the Build Category.
	//
	// This is the default.
	// Experimental.
	CodeBuildActionType_BUILD CodeBuildActionType = "BUILD"
	// The action will have the Test Category.
	// Experimental.
	CodeBuildActionType_TEST CodeBuildActionType = "TEST"
)

