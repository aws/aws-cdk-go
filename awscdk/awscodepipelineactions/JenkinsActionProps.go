package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Construction properties of `JenkinsAction`.
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
type JenkinsActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Default: 1.
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Default: - a name will be generated, based on the stage and action names,
	// if any of the action's variables were referenced - otherwise,
	// no namespace will be set.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Jenkins Provider for this Action.
	JenkinsProvider IJenkinsProvider `field:"required" json:"jenkinsProvider" yaml:"jenkinsProvider"`
	// The name of the project (sometimes also called job, or task) on your Jenkins installation that will be invoked by this Action.
	//
	// Example:
	//   "MyJob"
	//
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The type of the Action - Build, or Test.
	Type JenkinsActionType `field:"required" json:"type" yaml:"type"`
	// The source to use as input for this build.
	Inputs *[]awscodepipeline.Artifact `field:"optional" json:"inputs" yaml:"inputs"`
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
}

