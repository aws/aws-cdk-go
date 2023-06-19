package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties of the {@link CodeDeployServerDeployAction CodeDeploy server deploy CodePipeline Action}.
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
// Experimental.
type CodeDeployServerDeployActionProps struct {
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
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The CodeDeploy server Deployment Group to deploy to.
	// Experimental.
	DeploymentGroup awscodedeploy.IServerDeploymentGroup `field:"required" json:"deploymentGroup" yaml:"deploymentGroup"`
	// The source to use as input for deployment.
	// Experimental.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
}

