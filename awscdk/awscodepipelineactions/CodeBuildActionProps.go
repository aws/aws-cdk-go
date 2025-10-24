package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of the `CodeBuildAction CodeBuild build CodePipeline action`.
//
// Example:
//   // Create a Cloudfront Web Distribution
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//   var distribution Distribution
//
//
//   // Create the build project that will invalidate the cache
//   invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &PipelineProjectProps{
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
//   				},
//   			},
//   		},
//   	}),
//   	EnvironmentVariables: map[string]BuildEnvironmentVariable{
//   		"CLOUDFRONT_ID": &BuildEnvironmentVariable{
//   			"value": distribution.distributionId,
//   		},
//   	},
//   })
//
//   // Add Cloudfront invalidation permissions to the project
//   distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.Account, distribution.DistributionId)
//   invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Resources: []*string{
//   		distributionArn,
//   	},
//   	Actions: []*string{
//   		jsii.String("cloudfront:CreateInvalidation"),
//   	},
//   }))
//
//   // Create the pipeline (here only the S3 deploy and Invalidate cache build)
//   deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
//   deployInput := codepipeline.NewArtifact()
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	Stages: []StageProps{
//   		&StageProps{
//   			StageName: jsii.String("Deploy"),
//   			Actions: []IAction{
//   				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
//   					ActionName: jsii.String("S3Deploy"),
//   					Bucket: deployBucket,
//   					Input: deployInput,
//   					RunOrder: jsii.Number(1),
//   				}),
//   				codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   					ActionName: jsii.String("InvalidateCache"),
//   					Project: invalidateBuildProject,
//   					Input: deployInput,
//   					RunOrder: jsii.Number(2),
//   				}),
//   			},
//   		},
//   	},
//   })
//
type CodeBuildActionProps struct {
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
	// The source to use as input for this action.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The action's Project.
	Project awscodebuild.IProject `field:"required" json:"project" yaml:"project"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	// Default: true.
	//
	CheckSecretsInPlainTextEnvVariables *bool `field:"optional" json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Combine the build artifacts for a batch builds.
	//
	// Enabling this will combine the build artifacts into the same location for batch builds.
	// If `executeBatchBuild` is not set to `true`, this property is ignored.
	// Default: false.
	//
	CombineBatchBuildArtifacts *bool `field:"optional" json:"combineBatchBuildArtifacts" yaml:"combineBatchBuildArtifacts"`
	// The environment variables to pass to the CodeBuild project when this action executes.
	//
	// If a variable with the same name was set both on the project level, and here,
	// this value will take precedence.
	// Default: - No additional environment variables are specified.
	//
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Trigger a batch build.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	// Default: false.
	//
	ExecuteBatchBuild *bool `field:"optional" json:"executeBatchBuild" yaml:"executeBatchBuild"`
	// The list of additional input Artifacts for this action.
	//
	// The directories the additional inputs will be available at are available
	// during the project's build in the CODEBUILD_SRC_DIR_<artifact-name> environment variables.
	// The project's build always starts in the directory with the primary input artifact checked out,
	// the one pointed to by the `input` property.
	// For more information,
	// see https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html .
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The list of output Artifacts for this action.
	//
	// **Note**: if you specify more than one output Artifact here,
	// you cannot use the primary 'artifacts' section of the buildspec;
	// you have to use the 'secondary-artifacts' section instead.
	// See https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	// for details.
	// Default: the action will not have any outputs.
	//
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
	// The type of the action that determines its CodePipeline Category - Build, or Test.
	// Default: CodeBuildActionType.BUILD
	//
	Type CodeBuildActionType `field:"optional" json:"type" yaml:"type"`
}

