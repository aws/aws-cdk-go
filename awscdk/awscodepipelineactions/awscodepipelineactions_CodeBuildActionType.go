package awscodepipelineactions


// The type of the CodeBuild action that determines its CodePipeline Category - Build, or Test.
//
// The default is Build.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var project pipelineProject
//
//   sourceOutput := codepipeline.NewArtifact()
//   testAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("IntegrationTest"),
//   	project: project,
//   	input: sourceOutput,
//   	type: codepipeline_actions.codeBuildActionType_TEST,
//   })
//
type CodeBuildActionType string

const (
	// The action will have the Build Category.
	//
	// This is the default.
	CodeBuildActionType_BUILD CodeBuildActionType = "BUILD"
	// The action will have the Test Category.
	CodeBuildActionType_TEST CodeBuildActionType = "TEST"
)

