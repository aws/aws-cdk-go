package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties of {@link EcsDeployAction}.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//   var service fargateService
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   buildOutput := codepipeline.NewArtifact()
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []iAction{
//   		codepipeline_actions.NewEcsDeployAction(&EcsDeployActionProps{
//   			ActionName: jsii.String("DeployAction"),
//   			Service: *Service,
//   			// if your file is called imagedefinitions.json,
//   			// use the `input` property,
//   			// and leave out the `imageFile` property
//   			Input: buildOutput,
//   			// if your file name is _not_ imagedefinitions.json,
//   			// use the `imageFile` property,
//   			// and leave out the `input` property
//   			ImageFile: buildOutput.AtPath(jsii.String("imageDef.json")),
//   			DeploymentTimeout: awscdk.Duration_Minutes(jsii.Number(60)),
//   		}),
//   	},
//   })
//
// Experimental.
type EcsDeployActionProps struct {
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
	// The ECS Service to deploy.
	// Experimental.
	Service awsecs.IBaseService `field:"required" json:"service" yaml:"service"`
	// Timeout for the ECS deployment in minutes.
	//
	// Value must be between 1-60.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-ECS.html
	//
	// Experimental.
	DeploymentTimeout awscdk.Duration `field:"optional" json:"deploymentTimeout" yaml:"deploymentTimeout"`
	// The name of the JSON image definitions file to use for deployments.
	//
	// The JSON file is a list of objects,
	// each with 2 keys: `name` is the name of the container in the Task Definition,
	// and `imageUri` is the Docker image URI you want to update your service with.
	// Use this property if you want to use a different name for this file than the default 'imagedefinitions.json'.
	// If you use this property, you don't need to specify the `input` property.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-create.html#pipelines-create-image-definitions
	//
	// Experimental.
	ImageFile awscodepipeline.ArtifactPath `field:"optional" json:"imageFile" yaml:"imageFile"`
	// The input artifact that contains the JSON image definitions file to use for deployments.
	//
	// The JSON file is a list of objects,
	// each with 2 keys: `name` is the name of the container in the Task Definition,
	// and `imageUri` is the Docker image URI you want to update your service with.
	// If you use this property, it's assumed the file is called 'imagedefinitions.json'.
	// If your build uses a different file, leave this property empty,
	// and use the `imageFile` property instead.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/pipelines-create.html#pipelines-create-image-definitions
	//
	// Experimental.
	Input awscodepipeline.Artifact `field:"optional" json:"input" yaml:"input"`
}

