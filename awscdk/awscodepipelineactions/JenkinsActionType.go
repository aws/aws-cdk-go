package awscodepipelineactions


// The type of the Jenkins Action that determines its CodePipeline Category - Build, or Test.
//
// Note that a Jenkins provider, even if it has the same name,
// must be separately registered for each type.
//
// Example:
//   var jenkinsProvider jenkinsProvider
//
//   buildAction := codepipeline_actions.NewJenkinsAction(&JenkinsActionProps{
//   	ActionName: jsii.String("JenkinsBuild"),
//   	JenkinsProvider: jenkinsProvider,
//   	ProjectName: jsii.String("MyProject"),
//   	Type: codepipeline_actions.JenkinsActionType_BUILD,
//   })
//
type JenkinsActionType string

const (
	// The Action will have the Build Category.
	JenkinsActionType_BUILD JenkinsActionType = "BUILD"
	// The Action will have the Test Category.
	JenkinsActionType_TEST JenkinsActionType = "TEST"
)

