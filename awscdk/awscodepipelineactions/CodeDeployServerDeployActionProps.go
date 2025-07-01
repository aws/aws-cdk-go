package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of the `CodeDeployServerDeployAction CodeDeploy server deploy CodePipeline Action`.
//
// Example:
//   var deploymentGroup serverDeploymentGroup
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
//   	PipelineName: jsii.String("MyPipeline"),
//   })
//
//   // add the source and build Stages to the Pipeline...
//   buildOutput := codepipeline.NewArtifact()
//   deployAction := codepipeline_actions.NewCodeDeployServerDeployAction(&CodeDeployServerDeployActionProps{
//   	ActionName: jsii.String("CodeDeploy"),
//   	Input: buildOutput,
//   	DeploymentGroup: DeploymentGroup,
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type CodeDeployServerDeployActionProps struct {
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
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	// Default: a new Role will be generated.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The CodeDeploy server Deployment Group to deploy to.
	DeploymentGroup awscodedeploy.IServerDeploymentGroup `field:"required" json:"deploymentGroup" yaml:"deploymentGroup"`
	// The source to use as input for deployment.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
}

