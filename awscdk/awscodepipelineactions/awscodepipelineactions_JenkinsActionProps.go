package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Construction properties of {@link JenkinsAction}.
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
// Experimental.
type JenkinsActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Jenkins Provider for this Action.
	// Experimental.
	JenkinsProvider IJenkinsProvider `field:"required" json:"jenkinsProvider" yaml:"jenkinsProvider"`
	// The name of the project (sometimes also called job, or task) on your Jenkins installation that will be invoked by this Action.
	//
	// Example:
	//   "MyJob"
	//
	// Experimental.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The type of the Action - Build, or Test.
	// Experimental.
	Type JenkinsActionType `field:"required" json:"type" yaml:"type"`
	// The source to use as input for this build.
	// Experimental.
	Inputs *[]awscodepipeline.Artifact `field:"optional" json:"inputs" yaml:"inputs"`
	// Experimental.
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
}

