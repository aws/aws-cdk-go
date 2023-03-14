package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties of {@link EcrSourceAction}.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var ecrRepository repository
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewEcrSourceAction(&ecrSourceActionProps{
//   	actionName: jsii.String("ECR"),
//   	repository: ecrRepository,
//   	imageTag: jsii.String("some-tag"),
//   	 // optional, default: 'latest'
//   	output: sourceOutput,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
// Experimental.
type EcrSourceActionProps struct {
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
	// Experimental.
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The repository that will be watched for changes.
	// Experimental.
	Repository awsecr.IRepository `field:"required" json:"repository" yaml:"repository"`
	// The image tag that will be checked for changes.
	//
	// Provide an empty string to trigger on changes to any tag.
	// Experimental.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

